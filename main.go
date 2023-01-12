package main

import (
	"flag"
	"fmt"
	"time"
)

var mounths = map[string]struct{}{
	"jan": {},
	"feb": {},
	"mar": {},
	"apr": {},
	"may": {},
	"jun": {},
	"jul": {},
	"aug": {},
	"sep": {},
	"oct": {},
	"nov": {},
	"dec": {},
}

func main() {
	mounth := flag.String("m", "jan", "mounth: jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec")
	day := flag.Int("d", 1, "day")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	if _, ok := mounths[*mounth]; !ok {
		fmt.Printf("invalid mounth: %s\n", *mounth)
		return
	}
	if *day < 1 || *day > 31 {
		fmt.Printf("invalid day: %d\n", *day)
		return
	}

	l := &logger{verbose: *verbose}
	s := &solver{
		l: l,
		cs: []coord{
			empty.find(*mounth),
			empty.find(fmt.Sprintf("%d", *day)),
		},
	}

	f := initField()
	start := time.Now()
	solution, _ := s.solve(f, sets)
	l.Info(solution)
	l.Info(time.Since(start))
}
