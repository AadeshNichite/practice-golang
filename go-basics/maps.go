package main

import "fmt"

func main() {
	statePop := map[int]string{
		1: "Delhi",
		2: "Mumbai",
		3: "Pune",
		4: "Delhi",
		5: "Mumbai",
		6: "Pune",
	}

	fmt.Println(statePop)
	statePop[1] = "kolhapur"
	fmt.Println(statePop)
	fmt.Println(len(statePop))
	pop, ok := statePop[1]
	if ok {
		fmt.Println(pop)
	} else {
		fmt.Println("No contain value for " + statePop[5])
	}
}
