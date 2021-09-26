package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
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

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("Error: ", err.Error())
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	handleError(err)

	defer conn.Close()

	// var resp Response
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input, err := strconv.ParseInt(scan.Text(), 10, 64)
		handleError(err)
		request := RequestFib{Number: int(input)}

		encoder := json.NewEncoder(conn)
		handleError(encoder.Encode(request))

		var resp ResponseFib
		decoder := json.NewDecoder(conn)
		handleError(decoder.Decode(&resp))

		fmt.Printf("%s %d\n", resp.Time, resp.Fib)
	}
}
