package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s \n", input)
	}
}


func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}