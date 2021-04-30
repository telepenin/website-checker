package listener

import (
	"github.com/telepenin/website-checker/shared"
	"log"
)

type StdoutListener struct {
}

func (l *StdoutListener) Process(resp *shared.Response) error {
	log.Printf("=====================")
	log.Printf("url: %v", resp.Website.Url)
	log.Printf("status code: %v", resp.Code)
	log.Printf("duration: %v", resp.Duration)
	if resp.Website.Regex != "" {
		var found bool
		if resp.Content != nil {
			found = true
		}
		log.Printf("is content found by regex: %v\n", found)
	}

	return nil

}
