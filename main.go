package main

import (
	"flag"
	"fmt"
	"os"
	"tweetdumper/streamer"
)

var (
	length int
	help   bool
	geo    bool
)

func init() {
	flag.IntVar(&length, "l", 0, "number of tweets to dump: default = 0: for uninterupted stream l=0")
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
	streamer.Stream(flag.Arg(0), flag.Arg(1), length, geo, os.Stdout)
}
