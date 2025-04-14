package main

func main() {

	// var channel chan string = make(chan string)
	var channelBuf chan string = make(chan string, 2)
	println("---- GOROUTINES ----")
	//sleep()
	println("---- GOROUTINES ----")
	// SleepWithGoRoutines(channel)
	//time.Sleep(2000 * time.Millisecond) // waiting for go routing to end - simple way
	println("---- GOROUTINES ----")
	go Buff(channelBuf)
	// res := <-channel
	// println(res)
	for message := range channelBuf {
		println("Received message:", message)
	}
	// for {
	// 	select {
	// 	case message, ok := <-channelBuf:
	// 		if !ok {
	// 			// channle closed
	// 			return
	// 		}
	// 		println("Received message:", message)
	// 	}
	// }
}
