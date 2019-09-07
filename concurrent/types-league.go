package main

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
