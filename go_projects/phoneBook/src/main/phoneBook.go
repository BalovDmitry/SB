package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data []Entry

const (
	MIN              = 0
	MAX              = 26
	MIN_PHONE_NUMBER = 100
	MAX_PHONE_NUMBER = 199
	NAME_LENGTH      = 4
	SURNAME_LENGTH   = 5
)

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(NAME_LENGTH)
		surname := getString(SURNAME_LENGTH)
		n := strconv.Itoa(random(MIN_PHONE_NUMBER, MAX_PHONE_NUMBER))
		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	// How many records to insert
	n := 100
	populate(n)
	fmt.Printf("Data has %d entries.\n", len(data))

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found: ", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")

	}
}
