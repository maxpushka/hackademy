package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"
)

type RequestFib struct {
	Number int
}

type ResponseFib struct {
	Number int
	Fib    *big.Int
	Time   time.Duration
}

var fibCache map[int]*big.Int

func fib(num int) *big.Int {
	if num < 0 {
		panic("num < 0")
	}

	if val, ok := fibCache[num]; ok {
		return val
	}

	size := len(fibCache)
	a := fibCache[size-2]
	b := fibCache[size-1]

	for i := size; i <= num; i++ {
		a.Add(a, b)
		fibCache[i] = a
		a, b = b, a
	}
	return b
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("Error:", err.Error())
}

func main() {
	fibCache = make(map[int]*big.Int)
	fibCache[0] = big.NewInt(0)
	fibCache[1] = big.NewInt(1)

	fmt.Println("Booting up server...")
	const port string = "1234"

	ln, err := net.Listen("tcp", ":"+port)
	handleError(err)
	fmt.Printf("Listening on localhost:%v\n", port)
	conn, err := ln.Accept()
	handleError(err)

	for {
		var request RequestFib
		decoder := json.NewDecoder(conn)
		handleError(decoder.Decode(&request))

		start := time.Now()
		resp := ResponseFib{
			Number: request.Number,
			Fib:    fib(request.Number),
			Time:   time.Since(start),
		}

		enc := json.NewEncoder(conn)
		handleError(enc.Encode(resp))
	}
}
