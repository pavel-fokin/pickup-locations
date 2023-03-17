package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Add more cases for the unit tests.
func Test_ParseRoutesParams_Valid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		url string
	}{
		{"/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219"},
		{"/routes?src=13.388860,52.517037&dst=13.397634,52.529407"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("", test.url, nil)
		_, _, err := parseRoutesParams(req)
		assert.NoError(err)
	}

}

func Test_ParseRoutesParams_NotValid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		url string
	}{
		{"/routes"},
		{"/routes?src=13.388860,52.517037"},
		{"/routes?dst=13.388860,52.517037"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("", test.url, nil)
		_, _, err := parseRoutesParams(req)
		assert.Error(err)
	}

}
