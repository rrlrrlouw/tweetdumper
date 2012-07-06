package streamer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"errors"
	"tweetdumper/twitterstream"
)

func Stream(username string, password string, length int, geo bool, w io.Writer) (err error) {
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

type TweetReader struct {
	bufr   *bufio.Reader
}

func NewTweetReader(r io.Reader) *TweetReader {
	return &TweetReader{
		bufr:   bufio.NewReader(r),
	}
}

func (tr *TweetReader) jsonRead() (tweet twitterstream.Tweet, err error) {
	line, isPrefix, err := tr.bufr.ReadLine()
	if err != nil {
		return tweet, err
	}
	if isPrefix {
		return tweet, errors.New("tweet is a Prefix")
	}
	json.Unmarshal(line, &tweet)
	return tweet, err
}
