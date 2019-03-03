package rebalancerweb_test

import (
	"encoding/json"
	"github.com/pdbrito/rebalancer"
	. "github.com/pdbrito/rebalancer-web/http"
	"github.com/shopspring/decimal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	pricelist := rebalancer.Pricelist{
		"ETH": decimal.NewFromFloat(200),
		"BTC": decimal.NewFromFloat(5000),
	}

	s := NewServer(pricelist)

	t.Run("handles request to /healtcheck", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)

		s.Handler.ServeHTTP(w, r)

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
					t.Errorf("handler returned wrong body: got '%s' want '%s'",
						w.Body.String(), want)
				}
			})
		})
	})

	t.Run("handles request to /pricelist", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/pricelist", nil)

		pricelist := rebalancer.Pricelist{
			"ETH": decimal.NewFromFloat(200),
			"BTC": decimal.NewFromFloat(5000),
		}

		s.Handler.ServeHTTP(w, r)

		t.Run("successful response", func(t *testing.T) {
			t.Run("returns 200 OK", func(t *testing.T) {
				if status := w.Code; status != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, http.StatusOK)
				}
			})
			t.Run("body contains expected json", func(t *testing.T) {
				want, err := json.Marshal(pricelist)
				if err != nil {
					t.Errorf("error marshalling pricelist: %s", err)
				}
				if w.Body.String() != string(want) {
					t.Errorf("handler returned wrong body: got '%s' want '%s'",
						w.Body.String(), want)
				}
			})
		})
	})
}
