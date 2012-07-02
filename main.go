package main

import (
    "TweetTracer/twitterstream"
    "fmt"
    "encoding/json"
    "os"
    "io"
    "bufio"
)
var (
        path string = "/home/student/data"
        extension string = ".json"
    )

func main() {
    stream()
    //jsonDecode()
    //oldJsonDecode()
}

func stream() {
    stream := make(chan *twitterstream.Tweet)
    for {
        username, password := input()
        client := twitterstream.NewClient(username, password)
        err := client.Sample(stream)
        if err != nil {
            fmt.Println(err)
        } else {
            break
        }
    }

    for {
        tw := <-stream
        if tw.Coordinates.Coordinates  != nil {
           jsonDump(tw)
        }
    }
}

func input() (username string, password string) {
    //fmt.Println("Enter username for twitter account")
    fmt.Scan(&username)
    //fmt.Println("Enter password for twiiter account")
    fmt.Scan(&password)
    return
}

func jsonDump(tw *twitterstream.Tweet) {
    btw, err := json.Marshal(tw)
    if err == nil {
        writeLines(os.Stdout, btw)
        //fmt.Println("DUMP : ", dumpCounter)
    }
}

func writeLines(w io.Writer, b []byte) (err error) {
    _, err = w.Write(b)

    _, err = w.Write([]byte(fmt.Sprint("\n")))
    if err != nil {
        fmt.Println(err)
    }
    return
}

func oldJsonDecode() {
    f,_ := os.Open("/home/student/d1.json")
    data := make([]byte, 1000000000) 
    i,_ := f.Read(data)
    data = data[:i]
    var tweet []twitterstream.Tweet
    json.Unmarshal(data, &tweet)

    fmt.Println(tweet[7543].Text, len(tweet))
}

func jsonDecode() {
    f, _ := os.Open("/home/student/yay.json")
    r := bufio.NewReader(f)
    var tweet twitterstream.Tweet
    tweets := make([]twitterstream.Tweet, 1, 100000)

    line, isPrefix, err := r.ReadLine()
    for err == nil && !isPrefix {
        json.Unmarshal(line, &tweet)
        tweets[len(tweets)-1] = tweet
        tweets = tweets[:len(tweets)+1]
        line, isPrefix, err = r.ReadLine()
    }
    fmt.Println(tweets[17].Text)
}