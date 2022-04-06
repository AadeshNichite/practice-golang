package main

import "fmt"

func main() {
	statePop := map[string]int64{
		"Delhi":     50000,
		"Mumbai":    23000,
		"Pune":      856682,
		"Bangalore": 323222,
		"Panaji":    976232,
		"Kolhapur":  392772,
	}

	fmt.Println("Initial State Population Map")
	fmt.Println(statePop)

	//fetch value by key
	fmt.Printf("Population of Panaji %d", statePop["Panaji"])
	fmt.Println()

	//Length of map
	fmt.Printf("Length of Map %d", len(statePop))
	fmt.Println()

	//Fetch and assing value of map & use the ok keyword to verify the value is there or not
	val, ok := statePop["Bangalore"]
	if ok {
		fmt.Printf("State Popluation of Bangalore is %d", val)
		fmt.Println()
	} else {
		fmt.Println("Data not available for this state")
	}

	//Delete the value from map
	delete(statePop, "Bangalore")
	fmt.Println("After Delete state ", statePop)

	//Iterate Over map
	for state, population := range statePop {
		fmt.Println("State: ", state, "Population: ", population)
	}
}
