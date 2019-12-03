package main

import (
	"fmt"
	"os"

	"github.com/MindTooth/domeneshop-sdk-go/domeneshop"
)

func main() {
	tokenp, err := os.LookupEnv("DOMENESHOP_API_TOKEN")
	if !err {
		fmt.Println("Token is not set...")
		panic(err)
	}

	secretp, err := os.LookupEnv("DOMENESHOP_API_SECRET")
	if !err {
		fmt.Println("Secret is not set...")
		panic(err)
	}

	client := domeneshop.BasicAuth(tokenp, secretp)

	domains, _ := client.GetDomains()

	fmt.Println(domains)

	fmt.Printf("Hello, %s!  How are you, %s", tokenp, secretp)
}