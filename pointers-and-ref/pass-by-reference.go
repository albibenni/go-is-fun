package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytics *Analytics, message Message) {
	if message.Success {
		analytics.MessagesTotal += 1
		//both syntax work - the first is the shorthand
		(*analytics).MessagesSucceeded += 1
	} else {
		analytics.MessagesFailed += 1
		//both syntax work
		(*analytics).MessagesTotal += 1
	}
}
