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
				t.Errorf("handler returned wrong body: got '%s' want '%s'",
					w.Body.String(), want)
			}
		})
	})
}

func TestPricelistHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Errorf("http.NewRequest() err = %s, want nil", err)
	}

	pricelist := rebalancer.Pricelist{
		"ETH": decimal.NewFromFloat(200),
		"BTC": decimal.NewFromFloat(5000),
	}

	handlerFunc := PricelistHandler(pricelist)

	handlerFunc(w, r)

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
}

func TestNewServer(t *testing.T) {
	pricelist := rebalancer.Pricelist{
		"ETH": decimal.NewFromFloat(200),
		"BTC": decimal.NewFromFloat(5000),
	}

	_ = NewServer(pricelist)
}
