package main

import (
	"fmt"
	prmt "github.com/gitchander/permutation"
	"github.com/jinzhu/copier"
	"github.com/mr51m0n/gorc"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	var gorc0 gorc.Gorc

	// Timings
	// 5/5 - 120 - 275.43µs
	// 8/8 - 40320 - 148.265136ms (aws2) / 87.021991ms (mac)
	// 10/10 (aws2) - 3628800 -  12.978298938s / 10.671094261s
	// 10/10 (mac) - 3628800 - 3.751199486s, 3.405009508s, 3.851032714s
	// 11 - 39916800 - 2m28.170180976s (aws2) / 13m54.083020212s (mac)
	// 12 - 479001600 -

	num := 10
	routines := 10
	opponentTeam := getOpponentTeam(num)
	activeTeam := getActiveTeam(num)

	leagueMatches := NewLeagueMatches()
	p := prmt.New(opponentTeam)
	count := 0
	start := time.Now()

	for p.Next() {
		count += 1

		//wg.Add(1)
		gorc0.Inc()
		t := League{}
		copier.Copy(&t, &opponentTeam)

		//go func(team League, wg *sync.WaitGroup, c int) {
		go func(team League, c int) {
			//defer wg.Done()
			defer gorc0.Dec()

			//leagueMatch := NewConcurrentSlice()
			//for i, opp := range opponentTeam {
			//leagueMatch.Append(NewMatchup(
			//    activeTeam[i], // team
			//    opp,      // opponent
			//))
			//}
			//for t := range team.Iter() {
			//    //fmt.Printf("Team %+v\n", t)
			//    leagueMatch.Append(NewMatchup(
			//        activeTeam.GetItem(t.Index).(Team), // team
			//        t.Value.(Team),                     // opponent
			//    ))
			//}
			//_opp := make([]Team, len(opponentTeam))
			leagueMatch := NewLeagueMatch(c)

			//fmt.Printf("%v\n", team)
			//if hasDuplicates(team) {
			//    panic("duplicate")
			//}
			for i, opp := range team {
				match := NewMatchup(i+1, activeTeam[i], opp)
				leagueMatch.Matchups = append(leagueMatch.Matchups, match)
			}
			leagueMatch.Eval()
			//leagueMatches.Append(leagueMatch)

			// Eval and Compare with best.
			leagueMatches.Set(leagueMatch)

		}(t, count)
		//}(opponentTeam, &wg, count)

	}
	gorc0.WaitLow(routines)

	//wg.Wait()

	elapsed := time.Since(start)
	//fmt.Printf("leagueMatches Diff: %d\n", leagueMatches.Diff())
	fmt.Printf("team: %v\n", getActiveTeam(num))
	fmt.Printf("opponent: %v\n", getOpponentTeam(num))
	fmt.Printf("leagueMatches: %v\n", leagueMatches.ToString())
	//fmt.Printf("leagueMatches: %+v\n", leagueMatches.ToString())
	fmt.Printf("%d Combinations took %s\n", count, elapsed)
	fmt.Println("Done")

}

func hasDuplicates(elements []Team) bool {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}

	for _, v := range elements {
		if encountered[v.Name] == true {
			return true
		} else {
			// Record this element as an encountered element.
			encountered[v.Name] = true
		}
	}

	// Return the new slice.
	return false
}
