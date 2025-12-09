package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sum = 0
var sample = []string{"987654321111111","811111111111119","234234234234278","818181911112111"}
func joltageCalculator(joltage string)  {
	leng := len(joltage)
	first := 0
	last := 0
	i := 0
	largest := []string{}
	l := 0
	for l < leng-1{
		s := string(joltage[l])
		m, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if first < m {
			first = m
			i = l 
		}
		l++
	}
	largest = append(largest, strconv.Itoa(first))
	for j:= i+1; j<leng;j++ {
		st := string(joltage[j])
		k,err := strconv.Atoi(st)
		if err != nil {
			panic(err)
		}
		last = max(last,k)
	}
	largest = append(largest, strconv.Itoa(last))
	lj := strings.Join(largest,"")
	n,err := strconv.Atoi(lj)
	if err != nil {
		panic(err)
	}
	sum += n
	fmt.Println(sum)
}
func main() {
	
	file,err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		joltageCalculator(line)
	}
	//for i := range sample {
		//joltageCalculator(sample[i])
	//}

}
