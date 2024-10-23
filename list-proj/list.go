package main

import "fmt"

func main() {
	prices := []float64{1.2, 4.5}
	fmt.Println(prices[0:1])

	// appending
	prices = append(prices, 5.99)
	fmt.Println(prices)

	// make a queue like array
	prices = append(prices, 66.99)
	prices = prices[1:]
	fmt.Println(prices)

	newPrices := []float64{22.4, 245.55, 876.33}
	prices = append(prices, newPrices...)
}

// func main() {
// 	prices := [4]float64{1.2, 3.2, 4, 44}
// 	fmt.Println(prices)

// 	//Slices
// 	featuredPrices := prices[1:3] /*i,j-1*/
// 	featuredPrices = prices[1:] /*i,len-1*/
// 	featuredPrices = prices[:3] /*0,j-1*/
// 	fmt.Println(featuredPrices)
// }
