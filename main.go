package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var (
	errCrossBottom   = errors.New("corss bottom error")
	errCrossRight    = errors.New("corss right error")
	errAlreadyFilled = errors.New("already filled error")
)

/*
### ### ### ### ### ### ### ### ###
### jan feb mar apr may jun ### ###
### jul aug sep oct nov dec ### ###
###   1   2   3   4   5   6   7 ###
###   8   9  10  11  12  13  14 ###
###  15  16  17  18  19  20  21 ###
###  22  23  24  25  26  27  28 ###
###  29  30  31 ### ### ### ### ###
### ### ### ### ### ### ### ### ###
*/

var empty = field{
	{"jan", "feb", "mar", "apr", "may", "jun"},
	{"jul", "aug", "sep", "oct", "nov", "dec"},
	{"1", "2", "3", "4", "5", "6", "7"},
	{"8", "9", "10", "11", "12", "13", "14"},
	{"15", "16", "17", "18", "19", "20", "21"},
	{"22", "23", "24", "25", "26", "27", "28"},
	{"29", "30", "31"},
}

var sets = []set{
	/*
		 #		#		##		##		 ##		##		###		###
		##		##		##		##		###		###		 ##		##
		##		##		#		 #
	*/
	{
		{"a", []coord{{0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}},
		{"a", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 2}}},
		{"a", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}}},
		{"a", []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {1, 2}}},
		{"a", []coord{{0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1}}},
		{"a", []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 1}}},
		{"a", []coord{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {2, 1}}},
		{"a", []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}}},
	},
	/*
		 #		#		##		##		   #	#		####	####
		 #		#		#		 #		####	####	   #	#
		 #		#		#		 #
		##		##		#		 #
	*/
	{
		{"b", []coord{{0, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}},
		{"b", []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}}},
		{"b", []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 0}}},
		{"b", []coord{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}},
		{"b", []coord{{0, 1}, {1, 1}, {2, 1}, {3, 0}, {3, 1}}},
		{"b", []coord{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {3, 1}}},
		{"b", []coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1}}},
		{"b", []coord{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {3, 0}}},
	},
	/*
		  #		#		###		###		  #		#		###		###
		  #		#		#		  #		  #		#  		  #		#
		###		###		#		  #		###		###		  #		#
	*/
	{
		{"c", []coord{{0, 2}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{"c", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}},
		{"c", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}},
		{"c", []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}},
		{"c", []coord{{0, 2}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{"c", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}},
		{"c", []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}},
		{"c", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}},
	},
	/*
		 #		#		 #		#		  ##	##		###		 ###
		 #		#		##		##		###		 ###	  ##	##
		##		##		#		 #
		#		 #		#		 #
	*/
	{
		{"d", []coord{{0, 2}, {0, 3}, {1, 0}, {1, 1}, {1, 2}}},
		{"d", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}}},
		{"d", []coord{{0, 1}, {0, 2}, {0, 3}, {1, 0}, {1, 1}}},
		{"d", []coord{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {1, 3}}},
		{"d", []coord{{0, 1}, {1, 1}, {2, 0}, {2, 1}, {3, 0}}},
		{"d", []coord{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {3, 1}}},
		{"d", []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 1}}},
		{"d", []coord{{0, 1}, {1, 0}, {1, 1}, {2, 0}, {3, 0}}},
	},
	/*
		 #		#		#		 #		  #		 #		####	####
		 #		#		##		##		####	####	  #		 #
		##		##		#		 #
		 #		#		#		 #
	*/
	{
		{"e", []coord{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}},
		{"e", []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 2}}},
		{"e", []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 1}}},
		{"e", []coord{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}},
		{"e", []coord{{0, 1}, {1, 1}, {2, 0}, {2, 1}, {3, 1}}},
		{"e", []coord{{0, 1}, {1, 0}, {1, 1}, {2, 1}, {3, 1}}},
		{"e", []coord{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 0}}},
		{"e", []coord{{0, 0}, {1, 0}, {1, 1}, {2, 0}, {3, 0}}},
	},
	/*
		 ##		##		  #		#
		 #		 #		###		###
		##		 ##		#		  #
	*/
	{
		{"f", []coord{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}}},
		{"f", []coord{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}},
		{"f", []coord{{0, 1}, {0, 2}, {1, 1}, {2, 0}, {2, 1}}},
		{"f", []coord{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 2}}},
	},
	/*
		##		##		# #		###
		 #		#		###		# #
		##		##
	*/
	{
		{"g", []coord{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}},
		{"g", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}}},
		{"g", []coord{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}}},
		{"g", []coord{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {2, 1}}},
	},
	/*
		##		###
		##		###
		##
	*/
	{
		{"h", []coord{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}},
		{"h", []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1}}},
	},
}

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

type field [][]string

func initField() field {
	return field{
		{"_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_"},
	}
}

func (f field) find(v string) coord {
	for y, row := range f {
		for x, cell := range row {
			if cell == v {
				return coord{x, y}
			}
		}
	}
	return coord{}
}

func (f field) copy() field {
	nf := make(field, len(f))
	for y, row := range f {
		nrow := make([]string, len(row))
		copy(nrow, row)
		nf[y] = nrow
	}
	return nf
}

func (f field) String() string {
	rows := make([]string, len(f))
	for y, row := range f {
		rows[y] = strings.Join(row, " ")
	}
	return strings.Join(rows, "\n")
}

func (f field) check(cs ...coord) bool {
	for _, c := range cs {
		if f[c.y][c.x] != "_" {
			return false
		}
	}
	return true
}

func (f field) put(p piece) (field, bool) {
	for y, row := range f {
		for x := range row {
			nf, err := f.putTo(p, coord{x, y})
			if err == nil {
				return nf, true
			} else if err == errCrossRight {
				break
			} else if err == errCrossBottom {
				return f, false
			}
		}
	}
	return f, false
}

func (f field) putTo(p piece, s coord) (field, error) {
	nf := f.copy()
	for _, v := range p.cells {
		x, y := s.x+v.x, s.y+v.y
		if len(nf) <= y {
			return f, errCrossBottom
		}
		if len(nf[y]) <= x {
			return f, errCrossRight
		}
		if nf[y][x] != "_" {
			return f, errAlreadyFilled
		}
		nf[y][x] = p.sym
	}
	return nf, nil
}

type coord struct {
	x, y int
}

type piece struct {
	sym   string
	cells []coord
}

type set []piece

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

	mc, dc := empty.find(*mounth), empty.find(fmt.Sprintf("%d", *day))
	f := initField()
	start := time.Now()
	solution, _ := solve(*verbose, f, sets, mc, dc)
	fmt.Println(solution)
	fmt.Println(time.Since(start))
}

func solve(verbose bool, f field, pool []set, cs ...coord) (field, bool) {
	if len(pool) == 0 {
		return f, true
	}
	for k, s := range pool {
		for _, p := range s {
			nf, ok := f.put(p)
			if !ok {
				continue
			}
			if !nf.check(cs...) {
				continue
			}
			if verbose {
				fmt.Println(nf)
			}
			npool := make([]set, len(pool))
			copy(npool, pool)
			solution, ok := solve(
				verbose,
				nf,
				append(npool[:k], npool[k+1:]...),
				cs...,
			)
			if ok {
				return solution, true
			}
		}
	}
	return nil, false
}
