package main

func main() {

	var channel chan string = make(chan string)
	println("---- GOROUTINES ----")
	//sleep()
	println("---- GOROUTINES ----")
	SleepWithGoRoutines(channel)
	//time.Sleep(2000 * time.Millisecond) // waiting for go routing to end - simple way
	println("---- GOROUTINES ----")
	res := <-channel
	println(res)
}
