package main

import (
	"fmt"
	"strconv"
	"strings"
)
func ingredientChecker(ranges []string,ingredients []string) {
	fresh := 0
	for i := range ingredients{
		n,err := strconv.Atoi(ingredients[i])
		if err != nil {
			panic(err)
		}
		for j := range ranges{
			rangesArr := strings.Split(ranges[j],"-")
			m,err := strconv.Atoi(rangesArr[0])
			if err != nil {
				panic(err)
			}
			k,err := strconv.Atoi(rangesArr[1])
			if err != nil {
				panic(err)
			}
			if n >= m && n <= k{
				fresh++
				break
			}
		}
		
	}
	fmt.Println(fresh)
}
