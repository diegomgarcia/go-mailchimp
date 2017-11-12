package mailchimp_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mailchimp "github.com/diegomgarcia/go-mailchimp"
	"github.com/stretchr/testify/assert"
)

func TestCheckSubscriptionNotFoundError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, notFoundErrorResponse)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := mailchimp.NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.CheckSubscription("list_id", "john@reese.com")
	assert.Nil(t, memberResponse)
	assert.Equal(t, "Error 404 Resource Not Found (The requested resource could not be found.)", err.Error())

	errResponse, ok := err.(*mailchimp.ErrorResponse)
	assert.True(t, ok)
	assert.Equal(t, "Resource Not Found", errResponse.Title)
	assert.Equal(t, 404, errResponse.Status)
	assert.Equal(t, "The requested resource could not be found.", errResponse.Detail)
}
