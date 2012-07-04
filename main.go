package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"tweetdumper/streamer"
)

var (
	length      int
	help        bool
	geo         bool
	writestream     = make(chan []byte)
)

func init() {
	flag.IntVar(&length, "l", 10, "number of tweets to dump: default = 10: for uninterupted stream l=0")
	flag.BoolVar(&geo, "g", false, "only dump tweets with coordinate values")
}

func main() {
	flag.Parse()
	if flag.Arg(0) != "" && flag.Arg(1) != "" {
		execute()
	} else {
		fmt.Println("Invalid Use. Use -h for help")
	}
}

func execute() {
	go streamer.Stream(flag.Arg(0), flag.Arg(1), length, geo, writestream)
	writeLines(os.Stdout)
}

func writeLines(w io.Writer) (err error) {
	for b := range writestream {
		_, err = w.Write(b)
		_, err = w.Write([]byte(fmt.Sprint("\n")))
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
