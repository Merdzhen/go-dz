package main

import "fmt"

func main() {
	const usdEur = 0.86
	const usdRub = 81.27
	var eurUsd = usdRub / usdEur
	fmt.Println(eurUsd)
	sum, currencyFrom, currencyTo := scanData()
	convert(sum, currencyFrom, currencyTo)
}

func convert(sum int, currencyFrom string, currencyTo string) int {
	scanData()
	return 0
}

func scanData() (int, string, string) {
	var sum int
	var currencyFrom, currencyTo string
	fmt.Scan(&sum, &currencyFrom, &currencyTo)
	return sum, currencyFrom, currencyTo
}
