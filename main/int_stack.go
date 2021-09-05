package main

import "errors"

type intStack []int

func NewintStack() intStack {
	return intStack{}
}

func(s *intStack) Add(elem int) {
	*s = append(*s, elem)
	return
}

func(s *intStack) Pop() (int, error) {
	var elem int
	if len(*s) > 0 {
		elem = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return elem, nil 
	}
	return elem, errors.New("empty stack")
}
