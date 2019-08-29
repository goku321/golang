package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`: []string{"shaken", "not stirred", "martini"},
		`moneypenny_miss`: []string{"James Bond", "Literature", "Computer Science"},
	}

	for k, v := range x {
		fmt.Printf("%v: ", k)
		for _, i := range v {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}

	x["sah_deepak"] = []string{`the`, `time`, `has`, `come`}

	for k, v := range x {
		fmt.Println(k, v)
	}

	if v, ok := x["moneypenny_miss"]; ok {
		fmt.Println("value: ", v)
		delete(x, "moneypenny_miss")
	}
	
	fmt.Println("After Deleting a key")
	for k, v := range x {
		fmt.Println(k, v)
	}
}