package main

import "fmt"

func main() {
	c := make(chan string, 1)

	c <- "kek"

	for {
		go start(c)
	}
}

func start(c <-chan string) {
	fmt.Println(<-c)
}
