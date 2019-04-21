package main

import (
	"fmt"
	"sort"
)

type sortByLen []string

func (s sortByLen) Len() int  {
	return len(s)
}

func (s sortByLen) Less(i, j int ) bool {
	return len(s[i]) < len(s[j])
}

func (s sortByLen) Swap(i,j int ) {
		s[i], s[j] = s[j], s[i]
}

func main() {
	toSort := []string{"kiwi","qs","a", "peach", "banana"}
	sort.Strings(toSort);
	fmt.Println("default sort:", toSort)
	sort.Sort(sortByLen(toSort))
	fmt.Println("customized sort",toSort)
}