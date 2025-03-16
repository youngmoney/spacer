package main

import (
	"errors"
	"strconv"
	"strings"
)

func asNumbers(s []string) ([]int, error) {
	var i []int
	for _, p := range s {
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return []int{}, errors.New("invalid position number: " + p)
		}
		i = append(i, n)
	}
	return i, nil

}

func ParsePositions(s string) ([]int, error) {
	p := strings.Split(s, ",")
	return asNumbers(p)
}

func PositionString(i []int) string {
	var s []string
	for _, n := range i {
		s = append(s, strconv.Itoa(n))
	}
	return strings.Join(s, ",")
}
