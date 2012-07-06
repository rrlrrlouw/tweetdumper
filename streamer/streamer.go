package streamer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"tweetdumper/twitterstream"
)

func Stream(username string, password string, length int, geo bool, w io.WriteCloser) (err error) {
	stream := make(chan *twitterstream.Tweet)
	stream, err = connect(username, password, stream)
	if err != nil {
		return
	}
	err = read(length, geo, stream, w)
	return
}

func connect(username string, password string, stream chan *twitterstream.Tweet) (chan *twitterstream.Tweet, error) {
	client := twitterstream.NewClient(username, password)
	err := client.Sample(stream)
	if err != nil {
		return nil, err
	}
	return stream, err
}

func read(length int, geo bool, stream chan *twitterstream.Tweet, w io.WriteCloser) (err error) {
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
	w.Close()
	return nil
}

func jsonDump(tw *twitterstream.Tweet, counter int, w io.WriteCloser) (int, error) {
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

func jsonRead(r io.Reader) (tw []twitterstream.Tweet) {
	bufr := bufio.NewReader(r)
	var tweet twitterstream.Tweet
	tweets := make([]twitterstream.Tweet, 1, 100000)
	line, isPrefix, err := bufr.ReadLine()
	for err == nil && !isPrefix {
		json.Unmarshal(line, &tweet)
		tweets[len(tweets)-1] = tweet
		tweets = tweets[:len(tweets)+1]
		line, isPrefix, err = bufr.ReadLine()
	}
	tweets = tweets[:len(tweets)-1]
	return tweets
}
