package main

// Домашнее задание
// Hello now()
// Завести Go репозиторий на GitHub, написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
// Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(0)
	}
	fmt.Println(time)
}
