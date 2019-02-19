package rebalancerweb_test

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func iMakeARequestTo(method, endpoint string) (err error) {
	return godog.ErrPending
}

func theResponseCodeShouldBe(code int) error {
	return godog.ErrPending
}

func theResponseShouldMatchJson(json *gherkin.DocString) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {

	s.Step(`^I make a "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`,
		iMakeARequestTo)
	s.Step(`^the response code should be (\d+)$`,
		theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`,
		theResponseShouldMatchJson)

}
