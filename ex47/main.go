package main

import "fmt"

func main() {
	r1 := []string{"James", "Bond", "Shaken, not stirred"}
	r2 := []string{"Miss", "Money Penny", "Hello James"}
	records := [][]string{r1, r2}
	for _, record := range records {
		for _, data := range record {
			fmt.Println(data)
		}
	}
}
