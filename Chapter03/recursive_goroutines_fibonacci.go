package Chapter03

func FibConcurrent(n int) int {
	c := make(chan int)
	go fibWorker(n, c)
	return <-c
}

func fibWorker(n int, c chan int) {
	if n <= 1 {
		c <- n
		return
	}
	c1 := make(chan int)
	c2 := make(chan int)
	go fibWorker(n-1, c1)
	go fibWorker(n-2, c2)
	c <- <-c1 + <-c2
}
