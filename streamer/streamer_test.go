package streamer

import (
	"io"
	"testing"
	"tweetdumper/twitterstream"
)

func TestRead(t *testing.T) {
	stream := make(chan *twitterstream.Tweet, 4)
	num := 4
	geo := false

	reader, writer := io.Pipe()
	go read(num, geo, stream, writer)

	testtweet1 := twitterstream.Tweet{
		Retweet_count: 99,
		User:          twitterstream.User{Lang: "TestLang"},
		Place:         twitterstream.Place{Country: "TestCountry"},
		Coordinates:   twitterstream.Coordinates{Coordinates: []float64{23.22, 28.44}},
	}
	stream <- &testtweet1

	testtweet2 := twitterstream.Tweet{
		Text:        "This is a test",
		User:        twitterstream.User{Verified: false},
		Place:       twitterstream.Place{Country_code: "TestCountry_code"},
		Coordinates: twitterstream.Coordinates{Type: "TestType"},
	}
	stream <- &testtweet2

	testtweet3 := twitterstream.Tweet{
		Truncated: true,
		User:      twitterstream.User{Followers_count: 88},
		Place:     twitterstream.Place{Full_name: "TestFull_name"},
		In_reply_to_screen_name: "TestIn_reply_to_screen_name",
	}
	stream <- &testtweet3

	testtweet4 := twitterstream.Tweet{
		Favorited: true,
		User:      twitterstream.User{Location: "TestLocation"},
		Place:     twitterstream.Place{Id: "TestId"},
		Source:    "TestSource",
	}
	stream <- &testtweet4

	tr := NewTweetReader(4, reader)
	tr.jsonRead()
	tr.jsonRead()
	tr.jsonRead()
	tr.jsonRead()
	tweets := tr.tweets

	if tweets[0].Retweet_count != 99 {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[0].User.Lang != "TestLang" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[0].Place.Country != "TestCountry" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[0].Coordinates.Coordinates[0] != 23.22 {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[0].Coordinates.Coordinates[1] != 28.44 {
		t.Errorf("Not returning same data as was sent")
	}

	if tweets[1].Text != "This is a test" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[1].User.Verified != false {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[1].Place.Country_code != "TestCountry_code" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[1].Coordinates.Type != "TestType" {
		t.Errorf("Not returning same data as was sent")
	}

	if tweets[2].Truncated != true {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[2].User.Followers_count != 88 {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[2].Place.Full_name != "TestFull_name" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[2].In_reply_to_screen_name != "TestIn_reply_to_screen_name" {
		t.Errorf("Not returning same data as was sent")
	}

	if tweets[3].Favorited != true {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[3].User.Location != "TestLocation" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[3].Place.Id != "TestId" {
		t.Errorf("Not returning same data as was sent")
	}
	if tweets[3].Source != "TestSource" {
		t.Errorf("Not returning same data as was sent")
	}
}
