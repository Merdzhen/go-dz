package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.86,
	"RUB": 81.27,
}

func main() {
	sum, originalCurrency, exchangeCurrency := scanData()
	convertedSum := convert(sum, originalCurrency, exchangeCurrency, rates)
	convertedSum = math.Round(convertedSum * 100) / 100
	fmt.Printf("%.2f %s\n", convertedSum, strings.ToUpper(exchangeCurrency))
}

func convert(sum float64, originalCurrency string, exchangeCurrency string, rates map[string]float64) float64 {
	originalCurrency = strings.ToUpper(originalCurrency)
	exchangeCurrency = strings.ToUpper(exchangeCurrency)
	
	fromRate, ok1 := rates[originalCurrency]
	toRate, ok2 := rates[exchangeCurrency]
	if !ok1 || !ok2 {
		fmt.Println("Unknown currency")
		return 0
	}

	// Переводим в USD, затем в нужную валюту
	usdValue := sum / fromRate
	converted := usdValue * toRate

	return converted
}

func scanData() (float64, string, string) {
	var originalCurrency, sum, exchangeCurrency string
	isValidOriginalCurr := false
	for !isValidOriginalCurr {
		fmt.Println("Enter original currency (USD/EUR/RUB): ")
		fmt.Scan(&originalCurrency)
		isValidOriginalCurr = checkCurrency(originalCurrency)
	}

	isValidSum := false
	for !isValidSum {
		fmt.Println("Enter sum to exchange: ")
		fmt.Scan(&sum)
		isValidSum = checkSum(sum)
	}
	floatSum, _ := strconv.ParseFloat(sum, 64)

	isValidExchangeCurr := false
	for !isValidExchangeCurr {
		fmt.Printf("Enter exchange currency (%s): \n", getExchangeCurr(originalCurrency))
		fmt.Scan(&exchangeCurrency)
		isValidExchangeCurr = checkCurrency(exchangeCurrency) && !strings.EqualFold(exchangeCurrency, originalCurrency)
	}

	return floatSum, strings.ToUpper(originalCurrency), strings.ToUpper(exchangeCurrency)
}

func checkCurrency(currency string) bool {
	currLowerCase := strings.ToLower(currency)
	return currLowerCase == "usd" || currLowerCase == "eur" || currLowerCase == "rub"
}

func getExchangeCurr(currency string) string {
	curr := strings.ToLower(currency)
	switch curr {
	case "usd":
		return ("EUR/RUB")
	case "eur":
		return ("USD/RUB")
	default:
		return ("USD/EUR")
	}
}

func checkSum(sum string) bool {
	sum = strings.TrimSpace(sum)
	_, err := strconv.ParseFloat(sum, 64)
	return err == nil
}
