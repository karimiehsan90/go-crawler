package parser

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type Parser struct{}

func (p *Parser) Parse(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Warn("Unable to get url: ", url, err.Error())
	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			log.Warn("Unable to get page: ", response.StatusCode)
		} else {
			_, _ = html.Parse(response.Body)
			fmt.Println(response.Body)
		}
	}
}
