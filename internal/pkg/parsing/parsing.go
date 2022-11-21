package parsing

import (
	"log"
	"net/http"
)

type Website struct {
	Id              int
	Url             string
	TitleItem       string
	NewsItem        string
	DescriptionItem string
	LinkItem        string
	PubDateItem     string
}

type News struct {
	Id          int
	Title       string
	Link        string
	Description string
	PubDate     string
}

type parsing struct {
	client http.Client
}

func NewParsing(client http.Client) *parsing {
	return &parsing{client: client}
}

func (p *parsing) Parse(website Website) (news []News, err error) {
	response, err := p.client.Get(website.Url)
	if err != nil {
		log.Printf("Can't request to %s, error - %s", website.Url, err)
	}

}
