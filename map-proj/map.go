package main

import "fmt"

type floatMap map[string]float64

func output(m floatMap) {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["AWS"])

	websites["LinkedIn"] = "https://linkedin.com"

	delete(websites, "Google")
	fmt.Println(websites)

	// creates an array with mentione length
	userName := make([]string, 2)

	// error prone if use userName := []string{}
	userName[0] = "Pqr"

	userName = append(userName, "Abc")
	userName = append(userName, "Xyz")

	// second arg is for length
	rating := make(floatMap, 3)

	rating["Go"] = 4.5
	rating["React"] = 4.8

	for range userName {
		// print(i)
	}

	for i, val := range userName {
		print(i, val)
	}

	for i := range userName {
		print(i)
	}
}
