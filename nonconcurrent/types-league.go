package main

type League []Team

func (l League) Len() int {
	return len(l)
}

func (l League) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func NewLeague(teams ...Team) League {
	matchups := make(League, 0)
	for _, t := range teams {
		matchups = append(matchups, t)
	}

	//cs.Appends(teams...)
	return matchups
}
