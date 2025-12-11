package main
import (
	"bufio"
	"os"
)
var sample = []string{"..@@.@@@@.","@@@.@.@.@@","@@@@@.@.@@","@.@@@@..@.","@@.@@@@.@@",".@@@@@@@.@",".@.@.@.@@@","@.@@@.@@@@",".@@@@@@@@.","@.@.@@@.@."}
var inputs = []string{}
func main() {
	file,err := os.Open("input.txt")
	if err != nil {
		panic(err)
		
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		inputs = append(inputs,line)
	}
	ForkLiftPart2(inputs)	

}
