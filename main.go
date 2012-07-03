package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"tweetdumper/streamer"
)

var (
	length      *int  = flag.Int("l", 10, "number of tweets to dump")
	help        *bool = flag.Bool("h", false, "helpfile")
	geo         *bool = flag.Bool("g", false, "only dump tweets with coordinate values")
	counter     int   = 0
	writestream       = make(chan []byte)
)

func main() {
	flag.Parse()
	if *help {
		helpfile()
	} else if flag.Arg(0) != "" && flag.Arg(1) != "" {
		execute()
	} else {
		fmt.Println("Invalid Use. Use -h for help")
	}
}

func helpfile() {
	fmt.Println("\n -h   : helpfile")
	fmt.Println(" -g   : only dump tweets with coordinate values")
	fmt.Println(" -l=i : A total of i tweets will be dumped (default = 10)")
	fmt.Println("You need to enter a valid twitter account's username and password")
	fmt.Println("example: main.go -g -l=100 Username PaSsWoRd\n")
}

func execute() {
	go streamer.Stream(flag.Arg(0), flag.Arg(1), length, geo, writestream)
	writeLines(os.Stdout)
}

func writeLines(w io.Writer) (err error) {
	for _ = range writestream {
		b := <-writestream
		_, err = w.Write(b)
		_, err = w.Write([]byte(fmt.Sprint("\n")))
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
