package streamer

import (
	"encoding/json"
	"fmt"
	"os"
	"tweetdumper/twitterstream"
)

var counter = 0

func Stream(username string, password string, length *int, geo *bool, ch chan []byte) {
	stream := make(chan *twitterstream.Tweet)
	client := twitterstream.NewClient(username, password)
	err := client.Sample(stream)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for counter < *length {
		tw := <-stream
		if *geo {
			if tw.Coordinates.Coordinates != nil {
				jsonDump(tw, ch)
			}
		} else {
			jsonDump(tw, ch)
		}
	}
	close(ch)
}

func jsonDump(tw *twitterstream.Tweet, ch chan []byte) {
	btw, err := json.Marshal(tw)
	if err == nil {
		ch <- btw
		counter += 1
	}
}
