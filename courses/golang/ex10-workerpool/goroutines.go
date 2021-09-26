package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func Run(poolSize int) {
	jobs := make(chan float64, poolSize)
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)

	id := 1
	for scanner.Scan() {
		input := string(scanner.Bytes())
		f, err := strconv.ParseFloat(input, 64)
		handleError(err)
		jobs <- f

		if id <= poolSize {
			wg.Add(1)
			go workerTask(id, jobs, &wg)
			id++
		}
	}
	close(jobs)
	wg.Wait()
}

func workerTask(id int, jobs <-chan float64, wg *sync.WaitGroup) {
	fmt.Printf("worker:%d spawning\n", id)
	for job := range jobs {
		fmt.Printf("worker:%d sleep:%.1f\n", id, job)
		time.Sleep(time.Duration(int(job*1000)) * time.Millisecond)
	}
	fmt.Printf("worker:%d stopping\n", id)
	wg.Done()
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("ERROR: ", err)
}
