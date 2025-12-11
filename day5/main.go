package main
import (
	"os"
	"bufio"
)
var rangeSample = []string{"3-5","10-14","16-20","12-18"}
var ingredientSample = []string{"1","5","11","17","32"}
var r = []string{}
var i = []string{}
func main() {
	ranges,err := os.Open("range.txt")
	if err != nil {
		panic(err)
	}
	defer ranges.Close()
	rangesScanner := bufio.NewScanner(ranges)
	for rangesScanner.Scan() {
		rang := rangesScanner.Text()
		r = append(r,rang)
	}
	ingrendients,err := os.Open("ingredient.txt")
	if err != nil {
		panic(err)
	}
	defer ingrendients.Close()
	ingrendientScanner := bufio.NewScanner(ingrendients)
	for ingrendientScanner.Scan() {
		ingrendient := ingrendientScanner.Text()
		i = append(i,ingrendient)
	}
	println(r)
	ingredientChecker(rangeSample,ingredientSample)
	ingredientChecker(r,i)
	allFresh(rangeSample)
	allFresh(r)
}
