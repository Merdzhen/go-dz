package main

import "fmt"

func main() {
	const usdEur = 0.86
	const usdRub = 81.27
	var eurUsd = usdRub / usdEur
	fmt.Println(eurUsd)
}
