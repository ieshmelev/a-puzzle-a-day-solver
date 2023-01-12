package main

import (
	"errors"
	"strings"
)

var (
	errCrossBottom   = errors.New("corss bottom error")
	errCrossRight    = errors.New("corss rig	ht error")
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

func (f field) isEmpty(cs ...coord) bool {
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
