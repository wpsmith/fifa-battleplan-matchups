package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/thecodeteam/goodbye"
	"github.com/wpsmith/fifa-battleplan-matchups/data"
	"log"
	"os"
	"sync"
	"time"

	prmt "github.com/gitchander/permutation"
	"github.com/jinzhu/copier"
	//"github.com/mr51m0n/gorc"
)

var (
	numFlag, routinesFlag int
	//done chan bool
)

// Timings
// 5/5 - 120 - 275.43µs
// 8/8 - 40,320 - 148.265136ms (aws2) / 87.021991ms (mac)
// 10/10 (aws2) - 3,628,800 - 2m10.954049369s //  12.978298938s / 10.671094261s
// 10/10 (mac) - 3,628,800 - 3.751199486s, 3.405009508s, 3.851032714s
// 11 - 39,916,800 - 2m28.170180976s (aws2) / 13m54.083020212s (mac)
// 12 - 479,001,600 - (aws2) /  (mac)
func main() {
	// Create a context to use with the Goodbye library's functions.
	ctx := context.Background()

	// Always defer `goodbye.Exit` as early as possible since it is
	// safe to execute no matter what.
	defer goodbye.Exit(ctx, -1)

	// Invoke `goodbye.Notify` to begin trapping signals this process
	// might receive. The Notify function can specify which signals are
	// trapped, but if none are specified then a default list is used.
	// The default set is platform dependent. See the files
	// "goodbye_GOOS.go" for more information.
	goodbye.Notify(ctx)

	//done = make(chan bool)
	//runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	num := numFlag
	//routines := routinesFlag

	// Register two functions that will be executed when this process
	// exits.
	goodbye.Register(func(ctx context.Context, sig os.Signal) {
		num := flag.CommandLine.Lookup("n").Value.String()
		f, err := os.OpenFile(fmt.Sprintf("c%s.log", num), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)

		fmt.Printf("GOODBYE: %[1]d: %[1]s\n", sig)
		log.Printf("GOODBYE: %[1]d: %[1]s\n", sig)
	})

	f, err := os.OpenFile(fmt.Sprintf("c%d.log", num), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	var wg sync.WaitGroup
	//var gorc0 gorc.Gorc

	opponentTeam := data.GetOpponentTeam(num)
	activeTeam := data.GetActiveTeam(num)

	leagueMatches := data.NewLeagueMatches()
	p := prmt.New(opponentTeam)
	count := 0
	start := time.Now()

	for p.Next() {
		count += 1

		wg.Add(1)
		//gorc0.Inc()

		t := data.League{}
		copier.Copy(&t, &opponentTeam)

		go func(team data.League, wg *sync.WaitGroup, c int) {
			defer wg.Done()
			//go func(team League, c int) {
			//defer gorc0.Dec()

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
			leagueMatch := data.NewLeagueMatch(c)

			//fmt.Printf("%v\n", team)
			//if hasDuplicates(team) {
			//    panic("duplicate")
			//}
			for i, opp := range team {
				match := data.NewMatchup(i+1, activeTeam[i], opp)
				leagueMatch.Matchups = append(leagueMatch.Matchups, match)
			}
			leagueMatch.Eval()
			//leagueMatches.Append(leagueMatch)

			// Eval and Compare with best.
			leagueMatches.Set(leagueMatch)

			//}(t, count)
		}(opponentTeam, &wg, count)

	}
	wg.Wait()
	//gorc0.WaitLow(routines)

	elapsed := time.Since(start)
	//fmt.Printf("leagueMatches Diff: %d\n", leagueMatches.Diff())
	log.Printf("team: %v\n", data.GetActiveTeam(num))
	log.Printf("opponent: %v\n", data.GetOpponentTeam(num))
	log.Printf("leagueMatches: %v\n", leagueMatches.ToString())
	//fmt.Printf("leagueMatches: %+v\n", leagueMatches.ToString())
	log.Printf("%d Combinations took %s\n", count, elapsed)
	log.Println("Done")

	//go forever()
	//<-done
}

func forever() {
	//for {
	//    log.Printf("%v+\n", time.Now())
	//    time.Sleep(time.Minute)
	//}

	//done <- true
}

func hasDuplicates(elements []data.Team) bool {
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

func init() {
	flag.IntVar(&numFlag, "n", 10, "Number of teams")
	flag.IntVar(&routinesFlag, "r", numFlag*1000, "Number of routines to allow (defaults to n)")
}
