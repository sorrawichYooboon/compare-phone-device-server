package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	mocks "flowech-device-server/domain/usecases/device/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Controller_Device_Db(t *testing.T) {
	tests := []struct {
		name                 string
		wantStatusCode       int
		wantReqAndResUsecase interface{}
		wantData             interface{}
	}{
		{
			name:           "should call get device db names",
			wantStatusCode: 200,
			wantReqAndResUsecase: mocks.MockDeviceUsecase{
				GetDeviceDbNamesIsCalled: true,
			},
		},
	}

	for _, test := range tests {
		response := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(response)

		mockDeviceUsecase := &mocks.MockDeviceUsecase{}
		controller := HealthCheckController(mockDeviceUsecase)
		controller.GetDeviceDbNames(ctx)

		if !assert.Equal(t, test.wantStatusCode, response.Code) {
			t.Errorf("Test %s failed: want status code %d, got %d", test.name, test.wantStatusCode, response.Code)
		}

		t.Run(test.name, func(t *testing.T) {
			if test.wantReqAndResUsecase != nil {
				wantReqAndResUsecaseBytes, _ := json.Marshal(test.wantReqAndResUsecase)
				mockDeviceUsecaseBytes, _ := json.Marshal(mockDeviceUsecase)
				assert.Equal(t, string(wantReqAndResUsecaseBytes), string(mockDeviceUsecaseBytes))
			}

			if test.wantData != nil {
				wantDataBytes, _ := json.Marshal(test.wantData)
				assert.Equal(t, string(wantDataBytes), response.Body.String())
			}
		})
	}
}
