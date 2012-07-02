package main

import (
    "TweetTracer/twitterstream"
    "fmt"
    "encoding/json"
    "os"
    "io"
    "bufio"
    "flag"
)
var (
    length *int = flag.Int("l", 10, "number of tweets to dump")
    help *bool = flag.Bool("h", false, "helpfile")
    counter int = 0
)

func main() {
    flag.Parse()
    if *help{
        helpfile()
    } else if flag.Arg(0) != "" && flag.Arg(1) != ""{
        stream()
    } else {
        fmt.Println("Invalid Use. Use -h for help")
    }

    //jsonDecode()
    //oldJsonDecode()
}

func helpfile(){
    fmt.Println("\n -h   : helpfile")
    fmt.Println(" -l=i : A total of i tweets will be dumped (default = 10)")
    fmt.Println("You Need to enter a valid twitter account's username and password")
    fmt.Println("example: main.go -l=100 Username PaSsWoRd\n")
}
func stream() {
    stream := make(chan *twitterstream.Tweet)
    client := twitterstream.NewClient(flag.Arg(0), flag.Arg(1))
    err := client.Sample(stream)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for counter < *length{
        tw := <-stream
        if tw.Coordinates.Coordinates  != nil {
           jsonDump(tw)
        }
    }
}

func jsonDump(tw *twitterstream.Tweet) {
    btw, err := json.Marshal(tw)
    if err == nil {
        writeLines(os.Stdout, btw)
    }
}

func writeLines(w io.Writer, b []byte) (err error) {
    _, err = w.Write(b)
    _, err = w.Write([]byte(fmt.Sprint("\n")))
    counter += 1
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
    f, _ := os.Open("/home/student/sdata.json")
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
    tweets = tweets[:len(tweets)-1]
    fmt.Println(tweets[10].Text)
}