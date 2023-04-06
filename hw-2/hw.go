// package hw2

// import "strconv"

// func unpackString(input string) string {
// 	for i := 0; i < len(input); i++ {
// 		if val, err := strconv.Atoi(input[i]); err != nil {

// 		}
// 	}

// }
package hw2

import (
	"fmt"
	"time"
)

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk() {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
}
func walkTheDogs(dogs []Dog) {
	for _, d := range dogs {
		go func(d Dog) {
			d.Walk()
		}(d)
	}
	fmt.Println("everybody's home")
}
func tts() bool {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second * 2}}
	walkTheDogs(dogs)
	return true
}
