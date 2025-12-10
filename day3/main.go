package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sum int64 = 0
var sample = []string{"987654321111111","811111111111119","234234234234278","818181911112111"}
func joltageCalculator(joltage string) {
	targetLength := 12
	result := make([]byte, 0, targetLength)
	currentIndex := 0

	for digitsToTake := targetLength; digitsToTake > 0; digitsToTake-- {
		endSearchIndex := len(joltage) - digitsToTake
		
		maxDigit := ' '
		maxDigitIndex := -1

		for i := currentIndex; i <= endSearchIndex; i++ {
			if rune(joltage[i]) > maxDigit {
				maxDigit = rune(joltage[i])
				maxDigitIndex = i
			}
		}
		result = append(result, byte(maxDigit))
		currentIndex = maxDigitIndex + 1
	}

	lj := string(result)
	n, err := strconv.ParseInt(lj, 10, 64) // Using ParseInt with 64-bit to handle large numbers
	if err != nil {
		panic(err)
	}
	sum += n
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
	//	joltageCalculator(sample[i])
	//}
	fmt.Println(sum)
}
