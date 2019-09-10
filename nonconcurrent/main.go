package main

import (
	"fmt"
	prmt "github.com/gitchander/permutation"
	"time"
)

func main() {
	// Timings
	// 5 - 120 - 384.524µs
	// 8 - 40320 - 216.444217ms
	// 10 - 3628800 - 5.153742481s, 4.532734226s, 5.564259408s
	// 11 - 39916800 - 2m3.055477916s
	// 12 - 479001600 -

	num := 5
	opponentTeam := getOpponentTeam(num)
	activeTeam := getActiveTeam(num)

	var leagueMatches LeagueMatches
	p := prmt.New(opponentTeam)
	count := 0
	start := time.Now()

	for p.Next() {
		count += 1
		//fmt.Println(count)

		leagueMatch := NewLeagueMatch(count)
		for i, opp := range opponentTeam {
			match := NewMatchup(i+1, activeTeam[i], opp)
			leagueMatch.Matchups = append(leagueMatch.Matchups, match)
		}
		leagueMatch.Eval()
		//leagueMatches = append(leagueMatches, leagueMatch)
		leagueMatches.Set(leagueMatch)

		//fmt.Printf("leagueMatches: %+v\n", leagueMatches.ToString())
	}

	elapsed := time.Since(start)
	fmt.Printf("team: %v\n", getActiveTeam(num))
	fmt.Printf("opponent: %v\n", getOpponentTeam(num))
	fmt.Printf("leagueMatches: %v\n", leagueMatches.ToString())
	fmt.Printf("%d Combinations took %s\n", count, elapsed)
	fmt.Println("Done")

}
