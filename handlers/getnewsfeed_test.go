package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewsFeed(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/news", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetNewsFeed)
	handler.ServeHTTP(rr, req)
	// Out handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in out Request ans ResponseRecorder.
	resp := rr.Result()
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected 200 status code")
}
