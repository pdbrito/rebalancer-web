package rebalancerweb_test

import (
	. "github.com/pdbrito/rebalancer-web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health-check", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `{"alive": true}`
	if rr.Body.String() != want {
		t.Errorf("handler returned wrong body: got %s want %s",
			rr.Body.String(), want)
	}
}
