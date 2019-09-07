package main

type ExpectedGoals struct {
    Diff   int
    Normal   int
    Advanced int
}

type Team struct {
    Name string
    OVR  int16
}

type League []Team
func (l League) Len() int {
    return len(l)
}

func (l League) Swap(i, j int) {
   l[i], l[j] = l[j], l[i]
}

type Matchup struct {
    Team Team
    Opponent Team
}

//func (m Matchup) ExpectedGoals() int {
//
//}

func (m *Matchup) Diff() int16 {
    return m.Team.OVR - m.Opponent.OVR
}

func NewMatchup(home, opponent Team) Matchup {
    return Matchup{home, opponent}
}

type LeagueMatch []Matchup

type LeagueMatches []LeagueMatch

//var league = leagueMatches

type Chances []Chance

type Chance struct {
    Delta   int `json:"delta"`
    Great   int `json:"great"`
    Good    int `json:"good"`
    Basic   int `json:"basic"`
    Counter int `json:"counter"`
}
