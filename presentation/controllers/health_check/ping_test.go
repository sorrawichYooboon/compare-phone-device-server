package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Controller_Ping(t *testing.T) {
	tests := []struct {
		name           string
		wantStatusCode int
		wantData       interface{}
	}{
		{
			name:           "should response pong",
			wantStatusCode: 200,
			wantData:       PingResponse{Message: "pong"},
		},
	}

	for _, test := range tests {
		response := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(response)
		controller := NewHealthCheckController(nil)
		controller.Ping(ctx)

		t.Run(test.name, func(t *testing.T) {
			if test.wantData != nil {
				wantDataBytes, _ := json.Marshal(test.wantData)
				assert.Equal(t, string(wantDataBytes), response.Body.String())
			}
		})
	}
}
