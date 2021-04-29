package checker

import (
	"encoding/json"
	"fmt"
	config "github.com/telepenin/website-checker/config/src"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Processors []func(resp *Response) error

type Checker struct {
	Config     config.Checker
	Processors Processors
}

func (r *Response) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (c *Checker) Run() {
	ticker := time.NewTicker(time.Second * time.Duration(c.Config.Interval))

	for {
		select {
		case <-ticker.C:
			for _, site := range c.Config.Websites {
				go c.check(site)
			}
		}
	}
}

func (c *Checker) check(site config.Website) {
	client := http.Client{
		Timeout: time.Second * time.Duration(c.Config.Timeout),
	}

	resp, err := fetchWebsite(&client, site)
	if err != nil {
		log.Fatalf("unable to fetch the website %s: %v", site.Name, err)
	}
	for _, f := range c.Processors {
		err := f(resp)
		if err != nil {
			log.Fatalf("unable to process response: %v", err)
		}
	}
}

func getContent(regex string, body []byte) [][]byte {
	// FIXME: init regexp.MustCompile(regex) before the use
	return regexp.MustCompile(regex).FindAll(body, -1)
}

func fetchWebsite(client *http.Client, site config.Website) (*Response, error) {
	// for more complex metrics you may use https://golang.org/pkg/net/http/httptrace/
	// or if you need to exclude dns resolving, writing the request, reading the response, etc.
	start := time.Now()
	resp, err := client.Get(site.Url)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		return nil, fmt.Errorf("cannot fetch website %s: %w", site.Name, err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var content []byte
	if site.Regex != "" {
		content = getContent(site.Regex, bytes)
	}

	return &config.Response{
		Website:  site,
		Code:     resp.StatusCode,
		Duration: elapsed,
		Content:  content,
	}, nil
}