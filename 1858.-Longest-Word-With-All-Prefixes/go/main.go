package main

import (
	"sort"
)

type X struct {
	value    byte
	children []*X
}

func (x *X) add(word string) bool {
	if word == "" {
		return true
	}
	if len(word) == 1 {
		for _, c := range x.children {
			if c.value == word[0] {
				return true
			}
		}
		x.children = append(x.children, &X{value: word[0]})
		return true
	}
	for _, c := range x.children {
		if c.value == word[0] {
			return c.add(word[1:])
		}
	}
	return false
}

func AlgoOne(words []string) string {
	var result string
	sort.Strings(words)
	xx := X{}
	for _, word := range words {
		if xx.add(word) {
			if len(word) > len(result) {
				result = word
			}
		}
	}
	return result
}
