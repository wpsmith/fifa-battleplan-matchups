package main

import (
	"fmt"
	prmt "github.com/gitchander/permutation"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Timings
	// 5 - 120 - 810.127µs
	// 8 - 40320 - 237.132058ms
	// 10 - 3628800 - 3.751199486s, 3.405009508s, 3.851032714s
	// 11 - 39916800 - 1m45.037341738s
	// 12 - 479001600 -

	num := 8
	opponentTeam := getOpponentTeam(num)
	activeTeam := getActiveTeam(num)

	leagueMatches := NewLeagueMatches()
	//leagueMatches := NewConcurrentSlice()
	p := prmt.New(opponentTeam)
	count := 0
	start := time.Now()

	for p.Next() {
		count += 1
		//fmt.Println(count)

		wg.Add(1)
		go func(team League, wg *sync.WaitGroup, c int) {
			defer wg.Done()

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
			leagueMatch := NewLeagueMatch(c)
			for i, opp := range opponentTeam {
				match := NewMatchup(activeTeam[i], opp)
				leagueMatch.Matchups = append(leagueMatch.Matchups, match)
			}
			leagueMatch.Eval()
			leagueMatches.Append(leagueMatch)

			// Eval and Compare with best.
			leagueMatches.SetBestDelta(leagueMatch)
			leagueMatches.SetMostGreatChances(leagueMatch)
			leagueMatches.SetMostGoodChances(leagueMatch)
			leagueMatches.SetLeastCounterChances(leagueMatch)

		}(opponentTeam, &wg, count)

	}

	wg.Wait()

	elapsed := time.Since(start)
	//fmt.Printf("leagueMatches Diff: %d\n", leagueMatches.Diff())
	fmt.Printf("leagueMatches: %v\n", leagueMatches.ToString())
	//fmt.Printf("leagueMatches: %+v\n", leagueMatches.ToString())
	fmt.Printf("%d Combinations took %s\n", count, elapsed)
	fmt.Println("Done")

}
