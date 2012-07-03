package streamer

import (
	"encoding/json"
	"tweetdumper/twitterstream"
)

var (
	counter = 0
	stream  = make(chan *twitterstream.Tweet)
)

func Stream(username string, password string, length int, geo bool, ch chan []byte) {
	connect(username, password)
	read(length, geo, ch)
}

func connect(username string, password string) error {
	client := twitterstream.NewClient(username, password)
	return client.Sample(stream)
}

func read(length int, geo bool, ch chan []byte) error {
	defer close(ch)
	for counter < length {
		tw := <-stream
		if geo {
			if tw.Coordinates.Coordinates != nil {
				err := jsonDump(tw, ch)
				if err != nil {
					return err
				}
			}
		} else {
			err := jsonDump(tw, ch)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func jsonDump(tw *twitterstream.Tweet, ch chan []byte) error {
	btw, err := json.Marshal(tw)
	if err == nil {
		ch <- btw
		counter += 1
	}
	return err
}
