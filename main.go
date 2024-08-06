package main

import "fmt"

func main() {
	fmt.Printf("%q\n", ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}

// Function ReverseMenuIndex() rearanges slice items in the correct order without using append
func ReverseMenuIndex(menu []string) []string {
	cap := len(menu)
	menuArr := make([]string, cap)
	for i := cap - 1; i >= 0; i-- {
		menuArr[cap-i-1] = menu[i]
	}
	return menuArr
}
