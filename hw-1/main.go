// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/beevik/ntp"
// )

// func main() {
// 	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
// 	if err != nil {
// 		os.Stderr.WriteString(err.Error())
// 		os.Exit(0)
// 	}
// 	fmt.Println(time)

// 	a := "hello"
// 	b := []byte(a)
// 	c := []byte(a)
// 	fmt.Printf("%p %p", b, c)
// }
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создание случайной строки
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, 10000000)
	rand.Read(s)

	// Измерение времени выполнения преобразования string в []byte
	start := time.Now()
	b := []byte(s)
	fmt.Printf("%p\n", b)
	elapsed := time.Since(start)
	fmt.Println("Time taken to convert string to []byte:", elapsed.Nanoseconds())
	fmt.Printf("Time taken to convert string to []byte: %d", elapsed.Nanoseconds())
}
