package main

import "fmt"

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
}
