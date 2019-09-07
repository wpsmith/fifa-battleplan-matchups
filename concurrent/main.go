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

    num := 12
    opponentTeam := getOpponentTeam(num)
    activeTeam := getActiveTeam(num)

    leagueMatches := NewConcurrentSlice()
    p := prmt.New(opponentTeam)
    count := 0
    start := time.Now()

    for p.Next() {
        count += 1
        //fmt.Println(count)

        wg.Add(1)
        go func(team League, wg *sync.WaitGroup) {
            defer wg.Done()

            leagueMatch := NewConcurrentSlice()
            //for t := range team.Iter() {
            //    //fmt.Printf("Team %+v\n", t)
            //    leagueMatch.Append(NewMatchup(
            //        activeTeam.GetItem(t.Index).(Team), // team
            //        t.Value.(Team),                     // opponent
            //    ))
            //}
            for i, opp := range opponentTeam {
                //fmt.Printf("Team %+v\n", t)
                leagueMatch.Append(NewMatchup(
                    activeTeam[i], // team
                    opp,      // opponent
                ))
            }
            leagueMatches.Append(leagueMatch)
            //fmt.Printf("League Matchup %+v\n", leagueMatch)
        }(opponentTeam, &wg)

    }

    wg.Wait()

    elapsed := time.Since(start)
    //fmt.Printf("leagueMatches: %+v\n", leagueMatches)
    fmt.Printf("%d Combinations took %s\n", count, elapsed)
    fmt.Println("Done")

}
