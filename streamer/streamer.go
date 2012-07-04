package streamer

import (
	"encoding/json"
	"tweetdumper/twitterstream"
	"math"
)

func Stream(username string, password string, length int, geo bool, ch chan []byte) {
	stream  := make(chan *twitterstream.Tweet)
	stream, _ = connect(username, password, stream)
	read(length, geo, stream, ch)
}

func connect(username string, password string, stream chan *twitterstream.Tweet) (chan *twitterstream.Tweet , error) {
	client := twitterstream.NewClient(username, password)
	err := client.Sample(stream)
	if err != nil {
		return nil, err
	}
	return stream, err
}

func read(length int, geo bool, stream chan *twitterstream.Tweet, ch chan []byte) (err error) {
	defer close(ch)
	counter := 0
	if length == 0 {
		length = math.MaxInt32
	}
	for counter < length {
		tw := <-stream
		if geo {
			if tw.Coordinates.Coordinates != nil {
				counter, err = jsonDump(tw, counter, ch)
				if err != nil {
					return err
				}
			}
		} else {
			counter, err = jsonDump(tw, counter, ch)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func jsonDump(tw *twitterstream.Tweet, counter int, ch chan []byte) (int, error) {
	btw, err := json.Marshal(tw)
	if err == nil {
		ch <- btw
		counter += 1
	}
	return counter, err
}
