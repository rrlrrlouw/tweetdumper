package main

import "twitterstream"

func main() {
	stream := make(chan *twitterstream.Tweet)
	client := twitterstream.NewClient("Username", "Password")
	err := client.Sample(stream)
	if err != nil {
		println(err)
	}
	for {
		tw := <-stream
		println(tw.User.Screen_name, ": ", tw.Text)

	}
}
