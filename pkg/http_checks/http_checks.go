package http_checks

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sikalabs/mon/pkg/config"
)

func RunHttpChecks(
	config config.Config,
) (error, string) {
	var failed []string
	var out string
	for _, check := range config.HTTPChecks {
		resp, err := http.Get(check.URL)
		if err != nil {
			log.Printf("ERROR http check %s: %s\n", check.URL, err)
			out += fmt.Sprintf("ERROR http check %s: %s\n", check.URL, err)
			failed = append(failed, check.URL)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			log.Printf("ERROR http check %s: status %d\n", check.URL, resp.StatusCode)
			out += fmt.Sprintf("ERROR http check %s: status %d\n", check.URL, resp.StatusCode)
			failed = append(failed, check.URL)
			continue
		}
		log.Printf("OK http check %s: status %d\n", check.URL, resp.StatusCode)
		out += fmt.Sprintf("OK http check %s: status %d\n", check.URL, resp.StatusCode)
	}
	if len(failed) > 0 {
		return fmt.Errorf("http checks failed: %v", failed), out
	}
	return nil, out
}
