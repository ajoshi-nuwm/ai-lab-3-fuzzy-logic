package main

import (
	"github.com/ajoshi-nuwm/ai-lab-3-fuzzy-logic/wiper"
)

func main() {
	newWiper := wiper.NewWiper(10, 10, 0.3)
	for i := 0; i < 100; i++ {
		newWiper.Step()
	}
}
