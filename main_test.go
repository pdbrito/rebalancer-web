package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

type RequestContext struct {
	serverOnline bool
	response     struct {
		code int
		body []byte
	}
}

func FeatureContext(s *godog.Suite) {

	context := &RequestContext{}

	s.Step(`^the api server is running`,
		context.theAPIServerIsRunning)
	s.Step(`^I make a "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`,
		context.iMakeARequestTo)
	s.Step(`^the response code should be (\d+)$`,
		context.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`,
		context.theResponseShouldMatchJSON)

}

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"./http/"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func (rc *RequestContext) theAPIServerIsRunning() error {
	if rc.serverOnline == false {
		go main()
		rc.serverOnline = true
	}
	return nil
}

func (rc *RequestContext) iMakeARequestTo(method, endpoint string) error {
	r, err := http.NewRequest(method, "http://localhost:8080"+endpoint, nil)
	if err != nil {
		return err
	}
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	rc.response.code = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rc.response.body = body

	return resp.Body.Close()
}

func (rc *RequestContext) theResponseCodeShouldBe(code int) error {
	if status := rc.response.code; status != http.StatusOK {
		return fmt.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	return nil
}

func (rc *RequestContext) theResponseShouldMatchJSON(want *gherkin.DocString) (err error) {
	var expected []byte
	var data interface{}
	if err := json.Unmarshal([]byte(want.Content), &data); err != nil {
		return err
	}
	if expected, err = json.Marshal(data); err != nil {
		return err
	}
	if !bytes.Equal(rc.response.body, expected) {
		return fmt.Errorf("invalid json: want %s got %s",
			expected, rc.response.body)
	}
	return nil
}
