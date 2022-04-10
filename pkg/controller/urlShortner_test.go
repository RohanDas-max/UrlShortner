package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestUrlShortener(t *testing.T) {
	dsn := "host=localhost user=postgres password=abc@123 dbname=urlshortner_test port=4000 sslmode=disable"
	database.SetUp(dsn)

	tests := []struct {
		testName string
		BodyData []byte
		want     int
	}{
		{
			testName: "200 returned with right input",
			BodyData: []byte(`{"url":"http://google.com"}`),
			want:     200,
		},
		{
			testName: "400 returned with empty input",
			BodyData: []byte{},
			want:     400,
		},
		{
			testName: "400 returned because url already exists",
			BodyData: []byte(`{"url":"http://google.com"}`),
			want:     400,
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		t.Run(tc.testName, func(t *testing.T) {
			c, _ := gin.CreateTestContext(rec)
			requestData := tc.BodyData
			c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestData))
			c.Request.Header.Set("Content-Type", "application/json")
			UrlShortener(c)
			defer c.Request.Body.Close()
			assert.Equal(t, tc.want, rec.Result().StatusCode, rec.Body.String())
		})
		t.Cleanup(func() {
			database.DropTable()
		})
	}
}
