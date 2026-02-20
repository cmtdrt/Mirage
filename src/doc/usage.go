package doc

import "fmt"

func DisplayUsages(err error) {
	fmt.Println("Usage: mirage serve <config.json>")
	fmt.Println("       mirage serve --example")
	fmt.Println("       mirage serve <config.json> --port=8081")
	fmt.Println("       mirage serve --example --port=8081")
	fmt.Println("")
	fmt.Println("       mirage guide-en   (generate mirage-guide-en.md)")
	fmt.Println("       mirage guide-fr   (generate mirage-guide-fr.md)")
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
	}
}
