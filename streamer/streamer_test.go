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

	tr := NewTweetReader(reader)

	

	tw, err := tr.jsonRead()
	if err != nil {
		if tw.Retweet_count != 99 {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.User.Lang != "TestLang" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Place.Country != "TestCountry" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Coordinates.Coordinates[0] != 23.22 {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Coordinates.Coordinates[1] != 28.44 {
			t.Errorf("Not returning same data as was sent")
		}
	}


	tw, _ = tr.jsonRead()
	if err != nil {
		if tw.Text != "This is a test" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.User.Verified != false {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Place.Country_code != "TestCountry_code" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Coordinates.Type != "TestType" {
			t.Errorf("Not returning same data as was sent")
		}
	}

	tw, err = tr.jsonRead()
	if err != nil {
		if tw.Truncated != true {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.User.Followers_count != 88 {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Place.Full_name != "TestFull_name" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.In_reply_to_screen_name != "TestIn_reply_to_screen_name" {
			t.Errorf("Not returning same data as was sent")
		}
	}

	tw, err = tr.jsonRead()
	if err != nil {
		if tw.Favorited != true {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.User.Location != "TestLocation" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Place.Id != "TestId" {
			t.Errorf("Not returning same data as was sent")
		}
		if tw.Source != "TestSource" {
			t.Errorf("Not returning same data as was sent")
		}
	}
	writer.Close()

	tw, err = tr.jsonRead()
	if err == nil{
		t.Errorf("Error expected")
	}
}
