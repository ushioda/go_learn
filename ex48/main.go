package main

import "fmt"

func main() {
	nl := map[string][]string{
		"bond_james": {"shaken, not stirred", "martinis"},
		"no_dr":      {"being evil", "sunsets"},
	}

	nl["man_spider"] = []string{"pizza", "beer"}

	delete(nl, "man_spider")

	for key, p := range nl {
		for i, item := range p {
			fmt.Printf("%v likes %v, %d \n", key, item, i)
		}
	}
}
