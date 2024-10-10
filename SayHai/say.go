package SayHai

import "fmt"

func SayHai(name string, time ...string) {
	fmt.Println("hai", name, "sekarang jam", time)
}

func SayHaiv2(name string, time []string) {
	for _, time := range time {
		fmt.Println("hai", name, "sekarang jam", time)
	}
}


