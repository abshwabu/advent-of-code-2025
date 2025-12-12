package main
import (
	"os"
	"bufio"
)
var sample = []string{"123 328  51 64","45 64  387 23","6 98  215 314","*   +   *   +"}
func main() {
	arr := []string{}
	file,err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		arr = append(arr, line)
	}
	//calculator(sample)
	calculator(arr)
}
