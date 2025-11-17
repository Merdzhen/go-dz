package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	command, arr := scanData()
	var result = calculate(command, arr)
	fmt.Printf("%s: %.2f\n", command, result)
}

func calculate(command string, nums []int) float64 {
	commands := map[string]func([]int) float64{
		"SUM": func(numbers []int) float64{
			var sum int 
			for _, num := range numbers {
				sum += num
			}
			return float64(sum)
		},
		"AVG": func(numbers []int) float64{
			var sum int 
			n := len(numbers)
			for _, num := range numbers {
				sum += num
			}
			return float64(sum) / float64(n)
		},
		"MED": func(numbers []int) float64{
			sort.Ints(numbers)
			n := len(numbers)
			var med float64
			if n%2 == 1 {
				med = float64(numbers[n/2])
			} else {
				med = float64(numbers[n/2-1] + numbers[n/2]) / 2
			}
			return med
		},
	}
	
	commandFunc, ok := commands[command]
	if !ok {
		return 0.0
	}

	result := commandFunc(nums)
	return result
}

func scanData() (string, []int) {
	var command string
	isValidCommand := false
	for !isValidCommand {
		fmt.Println("Enter command (AVG/SUM/MED: )")
		fmt.Scan(&command)
		isValidCommand = checkCommand(command)
	}

	
	var userArr string
	isValidUserArr := false
	for !isValidUserArr {
		fmt.Println("Enter numbers separated by commas: ")
		reader := bufio.NewReader(os.Stdin)
		userArr, _ = reader.ReadString('\n')
		isValidUserArr = checkUserArr(userArr)
	}

	arr := formatUserArr(userArr)

	return strings.ToUpper(command), arr
}

func checkCommand(command string) bool {
	commandLowerCase := strings.ToLower(command)
	return commandLowerCase == "avg" || commandLowerCase == "sum" || commandLowerCase == "med"
}

func checkUserArr(userArr string) bool {
	if strings.TrimSpace(userArr) == "" {
		return false
	}

	parts := strings.Split(userArr, ",")
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return false
		}
		_, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
	}
	return true
}

func formatUserArr(userArr string) []int {
	parts := strings.Split(userArr, ",")
	var result []int
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return []int{}
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return []int{}
		}
		result = append(result, num)
	}
	return result
}
