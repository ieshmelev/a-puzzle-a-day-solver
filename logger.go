package main

import "fmt"

type logger struct {
	verbose bool
}

func (l *logger) Info(a ...interface{}) {
	fmt.Println(a...)
}

func (l *logger) Debug(a ...interface{}) {
	if !l.verbose {
		return
	}
	fmt.Println(a...)
}
