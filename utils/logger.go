package utils

import "fmt"

func ContainsKey(key string, players map[string]int) bool {
	for k, _ := range players {
		if k == key {
			return true
		}
	}
	return false
}

func PrintMap(players map[string]int) {
	fmt.Println("---Player Map---")
	for k, v := range players {
		if k != "" && k != "0" {
			fmt.Printf("%s : %d\n", k, v)
		}
	}
	fmt.Println("----------------")
}
