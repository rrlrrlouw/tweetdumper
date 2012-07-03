package streamer

import (
	"encoding/json"
	"testing"
	"tweetdumper/twitterstream"
)

func TestRead(t *testing.T) {
	ch := make(chan []byte)
	tw := new(twitterstream.Tweet)
	num := 4
	geo := false

	go read(num, geo, ch)
	testtweet1 := new(twitterstream.Tweet)
	testtweet1.Retweet_count = 99
	testtweet1.User.Lang = "TestLang"
	testtweet1.Place.Country = "TestCountry"
	testtweet1.Coordinates.Coordinates = []float64{23.22, 28.44}
	stream <- testtweet1

	data := <-ch
	json.Unmarshal(data, &tw)
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

	testtweet2 := new(twitterstream.Tweet)
	testtweet2.Text = "This is a test"
	testtweet2.User.Verified = false
	testtweet2.Place.Country_code = "TestCountry_code"
	testtweet2.Coordinates.Type = "TestType"
	stream <- testtweet2

	data = <-ch
	json.Unmarshal(data, &tw)
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

	testtweet3 := new(twitterstream.Tweet)
	testtweet3.Truncated = true
	testtweet3.User.Followers_count = 88
	testtweet3.Place.Full_name = "TestFull_name"
	testtweet3.In_reply_to_screen_name = "TestIn_reply_to_screen_name"
	stream <- testtweet3

	data = <-ch
	json.Unmarshal(data, &tw)
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

	testtweet4 := new(twitterstream.Tweet)
	testtweet4.Favorited = true
	testtweet4.User.Location = "TestLocation"
	testtweet4.Place.Id = "TestId"
	testtweet4.Source = "TestSource"
	stream <- testtweet4

	data = <-ch
	json.Unmarshal(data, &tw)
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
