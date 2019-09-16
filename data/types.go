package data

import (
	"fmt"
	"sync"
)

type ExpectedGoals struct {
	Delta     int
	Normal    int
	Advanced  int
	Predicted int
}
type Goals []ExpectedGoals

func (g Goals) GetMin() ExpectedGoals {
	return g[len(g)-1]
}
func (g Goals) GetMax() ExpectedGoals {
	return g[0]
}
func (g Goals) Get(diff int) ExpectedGoals {
	switch {
	case diff > 16:
		return g.GetMax()
	case diff < -16:
		return g.GetMin()
	default:
		return g[16-diff]
	}
}

type Team struct {
	Name string
	OVR  int
}

type Chance struct {
	Delta   int
	Optimal int
	Great   int
	Good    int
	Basic   int
	Counter int
}

type Chances []Chance

func (c Chances) GetMin() Chance {
	return c[len(c)-1]
}
func (c Chances) GetMax() Chance {
	return c[0]
}
func (c Chances) Get(diff int) Chance {
	switch {
	case diff > 16:
		return c.GetMax()
	case diff < -16:
		return c.GetMin()
	default:
		return c[16-diff]
	}
}

type Matchup struct {
	Index         int
	Team          Team
	Opponent      Team
	Delta         int
	Chances       Chance
	ExpectedGoals ExpectedGoals
}

func (m *Matchup) Diff() int {
	return m.Team.OVR - m.Opponent.OVR
}

func NewMatchup(i int, home, opponent Team) *Matchup {
	match := &Matchup{
		Index:    i,
		Team:     home,
		Opponent: opponent,
	}

	// Record Delta.
	match.Delta = match.Diff()

	// Record Chances.
	match.Chances = chances.Get(match.Delta)

	// Record Expected Goals.
	match.ExpectedGoals = expectedGoals.Get(match.Delta)

	return match
}

type LeagueMatch struct {
	Index    int
	Matchups []*Matchup
	Chance   Chance
	Delta    int
}

//func (lm *LeagueMatch) GetChances() Chance {
//    chances := Chance{0, 0, 0, 0, 0, 0}
//    for _, m := range lm.Matchups {
//        chances.Great += m.Chances.Great
//        chances.Good += m.Chances.Good
//        chances.Basic += m.Chances.Basic
//        chances.Counter += m.Chances.Counter
//    }
//
//    l := len(lm.Matchups)
//    chances.Great /= l
//    chances.Good /= l
//    chances.Basic /= l
//    chances.Counter /= l
//    return chances
//}
//func (lm *LeagueMatch) Diff() int {
//    totalDiff := 0
//    for _, m := range lm.Matchups {
//        totalDiff += m.Delta
//    }
//    return totalDiff / len(lm.Matchups)
//}
func (lm *LeagueMatch) Eval() {
	chances := Chance{0, 0, 0, 0, 0, 0}
	totalDiff := 0
	for _, m := range lm.Matchups {
		totalDiff += m.Delta
		chances.Great += m.Chances.Great
		chances.Good += m.Chances.Good
		chances.Basic += m.Chances.Basic
		chances.Counter += m.Chances.Counter
	}
	l := len(lm.Matchups)
	lm.Delta = totalDiff / l
	chances.Great /= l
	chances.Good /= l
	chances.Basic /= l
	chances.Counter /= l
	lm.Chance = chances
}

func (lm LeagueMatch) ToString() string {
	var output string = fmt.Sprintf("LEAGUE #%d\n", lm.Index)
	for _, leagueMatch := range lm.Matchups {
		output += fmt.Sprintf(
			"%d) %s (%d) v. %s (%d) : %d\n",
			leagueMatch.Index,
			leagueMatch.Team.Name,
			leagueMatch.Team.OVR,
			leagueMatch.Opponent.Name,
			leagueMatch.Opponent.OVR,
			leagueMatch.Delta,
		)
	}

	return output
}

func NewLeagueMatch(index int) *LeagueMatch {
	return &LeagueMatch{
		Index: index,
	}
}

type LeagueMatches struct {
	sync.RWMutex
	*ConcurrentSlice
	TopByMostGreatChances    *LeagueMatch
	TopByMostGoodChances     *LeagueMatch
	TopByLeastCounterChances *LeagueMatch
	TopByDelta               *LeagueMatch
}

func (lm *LeagueMatches) Set(match *LeagueMatch) {
	lm.Lock()
	defer lm.Unlock()

	lm.setBestDelta(match)
	lm.setMostGreatChances(match)
	lm.setMostGoodChances(match)
	lm.setLeastCounterChances(match)
}
func (lm *LeagueMatches) setMostGreatChances(match *LeagueMatch) {
	//lm.Lock()
	//defer lm.Unlock()

	topLeague := lm.TopByMostGreatChances
	if topLeague == nil || topLeague.Chance.Great < match.Chance.Great {
		lm.TopByMostGreatChances = match
	} else if topLeague.Chance.Great == match.Chance.Great {
		// @todo Determine tie breaker if they are equal
		switch {
		case topLeague.Chance.Good < match.Chance.Good:
			lm.TopByMostGreatChances = match
		}
	}
}
func (lm *LeagueMatches) setMostGoodChances(match *LeagueMatch) {
	//lm.Lock()
	//defer lm.Unlock()

	topLeague := lm.TopByMostGoodChances
	if topLeague == nil || topLeague.Chance.Good < match.Chance.Good {
		lm.TopByMostGoodChances = match
	} else if topLeague.Chance.Good == match.Chance.Good {
		// @todo Determine tie breaker if they are equal
		switch {
		case topLeague.Chance.Great < match.Chance.Great:
			lm.TopByMostGoodChances = match
		}
	}
}
func (lm *LeagueMatches) setLeastCounterChances(match *LeagueMatch) {
	//lm.Lock()
	//defer lm.Unlock()

	topLeague := lm.TopByLeastCounterChances
	if topLeague == nil || topLeague.Chance.Counter > match.Chance.Counter {
		lm.TopByLeastCounterChances = match
	} else if topLeague.Chance.Counter == match.Chance.Counter {
		// @todo Determine tie breaker if they are equal
		switch {
		case topLeague.Chance.Basic > match.Chance.Basic:
			lm.TopByLeastCounterChances = match
			break
		}
	}
}
func (lm *LeagueMatches) setBestDelta(match *LeagueMatch) {
	//lm.Lock()
	//defer lm.Unlock()

	topLeague := lm.TopByDelta
	if topLeague == nil || topLeague.Delta < match.Delta {
		lm.TopByDelta = match
	} else if topLeague.Chance.Counter == match.Chance.Counter {
		// @todo Determine tie breaker if they are equal
		if topLeague.Chance.Great < match.Chance.Great {
			lm.TopByDelta = match
		}
	}
}

func (lm LeagueMatches) MatchToString(match *LeagueMatch) string {
	return fmt.Sprintf(
		"LEAGUE %d (Diff: %d; %s):\n%s \n",
		match.Index,
		match.Delta,
		fmt.Sprintf(
			"Great: %d; Good: %d; Basic: %d; Counter: %d",
			match.Chance.Great,
			match.Chance.Good,
			match.Chance.Basic,
			match.Chance.Counter,
		),
		match.ToString(),
	)
}
func (lm LeagueMatches) ToString() string {
	var output string = ""
	//for leagueMatch := range lm.Iter() {
	//match := leagueMatch.Value.(*LeagueMatch)
	//match.Eval()
	//output += lm.MatchToString(match)
	//}

	output += fmt.Sprintf(
		"TopDelta: %v\nTopByMostGreatChances: %v\nTopByMostGoodChances: %v\nTopByLeastCounterChances: %v\n",
		//lm.TopByDelta.ToString(),
		//lm.TopByMostGreatChances.ToString(),
		//lm.TopByMostGoodChances.ToString(),
		//lm.TopByLeastCounterChances.ToString(),
		lm.MatchToString(lm.TopByDelta),
		lm.MatchToString(lm.TopByMostGreatChances),
		lm.MatchToString(lm.TopByMostGoodChances),
		lm.MatchToString(lm.TopByLeastCounterChances),
	)

	return output
}

//func (lm LeagueMatches) Diff() float64 {
//    var diff int = 0
//    for leagueMatch := range lm.Iter() {
//        match := leagueMatch.Value.(LeagueMatch)
//        diff += match.Diff()
//    }
//
//    return (float64)(diff / lm.Len())
//}

// NewLeagueMatches creates a new concurrent slice
func NewLeagueMatches() *LeagueMatches {
	cs := NewConcurrentSlice()
	return &LeagueMatches{
		ConcurrentSlice: cs,
	}
}

type League []Team

func (l League) Len() int {
	return len(l)
}

func (l League) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func NewLeague(teams ...Team) *ConcurrentSlice {
	cs := NewConcurrentSizedSlice(len(teams))
	for _, t := range teams {
		cs.Append(t)
	}
	//cs.Appends(teams...)
	return cs
}
