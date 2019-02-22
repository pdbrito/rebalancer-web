package rebalancerweb_test

import (
	. "github.com/pdbrito/rebalancer-web/http"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Errorf("http.NewRequest() err = %s, want nil", err)
	}

	HealthCheckHandler(w, r)

	t.Run("successful response", func(t *testing.T) {
		t.Run("returns 200 OK", func(t *testing.T) {
			if status := w.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
		t.Run("body contains expected json", func(t *testing.T) {
			want := `{"alive":true}`
			if w.Body.String() != want {
				t.Errorf("handler returned wrong body: got %s want %s",
					w.Body.String(), want)
			}
		})
	})
}
