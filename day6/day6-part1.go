package main
import (
	"strconv"
	"strings"
	"fmt"
)
func calculator(nums []string) {
	total := 0
	operatorLine := nums[len(nums)-1]
	operators := strings.Fields(operatorLine)
	fmt.Println()
	for i := 0; i<len(operators);i++ {
		result := 0
		if operators[i]=="+" {
		for j:=0; j < len(nums)-1;j++{

				arr := strings.Fields(nums[j])
				if i < len(arr) {
					n,err := strconv.Atoi(arr[i])
					if err != nil {
						panic(err)
					}
					fmt.Println(n)
					result += n

				}else {
					fmt.Printf("Warning: Missing data at row %d, col %d. Assuming 0 for sum.\n", j, i)
				}
							}
		}else if operators[i]=="*" {
			result = 1
		for j:=0; j < len(nums)-1;j++{
			
				arr := strings.Fields(nums[j])
				if i < len(arr) {
					n,err := strconv.Atoi(arr[i])
				if err != nil {
					panic(err)
				}
				fmt.Println(n)
				result *= n

				} else {
					fmt.Printf("Warning: Missing data at row %d, col %d. Assuming 1 for product.\n", j, i)
				}
							}
		}
	total += result
	}
	print(total)
}
