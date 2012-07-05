package streamer

import (
	"encoding/json"
	"tweetdumper/twitterstream"
	"math"
	"io"
	"fmt"
)

func Stream(username string, password string, length int, geo bool, w io.Writer) (err error) {
	stream  := make(chan *twitterstream.Tweet)
	stream, err = connect(username, password, stream)
	if err != nil {
		return
	}
	err = read(length, geo, stream, w)
	return
}

func connect(username string, password string, stream chan *twitterstream.Tweet) (chan *twitterstream.Tweet , error) {
	client := twitterstream.NewClient(username, password)
	err := client.Sample(stream)
	if err != nil {
		return nil, err
	}
	return stream, err
}

func read(length int, geo bool, stream chan *twitterstream.Tweet, w io.Writer) (err error) {
	counter := 0
	if length == 0 {
		length = math.MaxInt32
	}
	for counter < length {
		tw := <-stream
		
		if geo {
			if tw.Coordinates.Coordinates != nil {
				counter, err = jsonDump(tw, counter, w)
				if err != nil {
					return err
				}
			}
		} else {
			counter, err = jsonDump(tw, counter, w)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func jsonDump(tw *twitterstream.Tweet, counter int, w io.Writer) (int, error) {
	btw, err := json.Marshal(tw)
	if err == nil {
		_, err = w.Write(btw)
		_, err = w.Write([]byte(fmt.Sprint("\n")))
		if err != nil {
			return counter, err
		}
		counter += 1
	}
	return counter, err
}
