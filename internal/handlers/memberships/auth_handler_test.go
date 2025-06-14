package memberships

import (
	"bytes"
	"catalog-music/internal/models/memberships"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	controlMock := gomock.NewController(t)
	defer controlMock.Finish()

	mockService := NewMockmembershipService(controlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockService.EXPECT().SignUp(&memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				}).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			api := gin.New()

			h := &Handler{
				Engine:            api,
				membershipService: mockService,
			}

			h.AuthRoute()

			w := httptest.NewRecorder()
			endpoint := "/auth/signup"
			model := memberships.SignUpRequest{
				Email:    "test@test.com",
				Username: "test",
				Password: "test",
			}

			jsonReq, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(jsonReq)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			api.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
