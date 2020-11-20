package main

import (
	"flag"
	"os"

	"github.com/karimiehsan90/go-crawler/parser"
	"github.com/karimiehsan90/go-crawler/queue"
	log "github.com/sirupsen/logrus"
)

var pageParser = &parser.Parser{}

func main() {
	var url string
	flag.StringVar(&url, "url", "", "The url to start with it")
	flag.Parse()
	if url == "" {
		flag.Usage()
		os.Exit(1)
	}
	urlsQueue := &queue.Queue{}
	urlsQueue.PushBack(url)
	start(urlsQueue)
}

func start(urlsQueue *queue.Queue) {
	for {
		url, err := urlsQueue.PopFront()
		if err != nil {
			log.Warn("Unable to get any urls: ", err.Error())
		} else {
			pageParser.Parse(url)
		}
		break
	}
}
