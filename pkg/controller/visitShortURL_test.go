package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestGetShortURL(t *testing.T) {
	dsn := "host=localhost user=postgres password=abc@123 dbname=urlshortner_test port=4000 sslmode=disable"
	database.SetUp(dsn)

	tests := []struct {
		testName string
		query    string
		want     int
	}{
		{
			testName: "returned 301 ",
			query:    "56ti",
			want:     301,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			url := url.Values{}
			url.Add("url", tc.query)
			c.Request, _ = http.NewRequest(http.MethodGet, "/", strings.NewReader(url.Encode()))
			c.Request.Header.Set("Content-Type", "application/json")
			VisitShortURL(c)
			defer c.Request.Body.Close()
			assert.Equal(t, tc.want, rec.Result().StatusCode, rec.Body.String())
		})
	}
}
