package main

import "fmt"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(1*time.Second)
		//fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int)

	results := make(chan int,30000)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	// 3是缓存池大小
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//5 是所有线程数
	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	jobscount:=20

	for j := 1; j <= jobscount; j++ {
		jobs <- j
		//<-results
	}
	close(jobs)
	// Finally we collect all the results of the work.
	for a := 1; a <= 10; a++ {
		<-results
	}
}