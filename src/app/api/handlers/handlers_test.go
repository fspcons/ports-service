package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fspcons/ports-service/src/app/api/handlers"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/usecases/port"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TODO With more time, I'd use the httptest from the common go library to run a http server and mock the useCase
//  calls to test the rest handlers for both success and failure, checking the proper http status codes

func TestNewCreateHandler(t *testing.T) {
	defaultUC := &port.UseCaseMock{
		CreateFunc: func(_ context.Context, p *domain.Port) error {
			return nil
		},
	}

	validInput := domain.Port{ID: "someID"}
	testCases := []struct {
		name           string
		requestBody    any
		responseStatus int
		testResponse   func(in *bytes.Buffer) error
		expectedError  error
		portUseCase    port.UseCase
	}{
		{
			name:           "should return a bad request status for an invalid body",
			requestBody:    "invalid body",
			responseStatus: http.StatusBadRequest,
			expectedError:  echo.NewHTTPError(http.StatusBadRequest, handlers.MsgInvalidFormat),
			testResponse: func(in *bytes.Buffer) error {
				var got string
				if err := json.Unmarshal(in.Bytes(), &got); err != nil {
					t.Errorf("failed to unmarshal response %s", err.Error())
				} else {
					if got != handlers.MsgInvalidFormat {
						t.Errorf("unexpected post result. expected: %s - got: %s", handlers.MsgInvalidFormat, got)
					}
				}

				return nil
			},
			portUseCase: defaultUC,
		},
		{
			name:           "should return a unprocessable entity status for an invalid body",
			requestBody:    validInput,
			expectedError:  nil,
			responseStatus: http.StatusUnprocessableEntity,
			testResponse: func(in *bytes.Buffer) error {
				return nil
			},
			portUseCase: defaultUC,
		},
		{
			name:           "should return a created status for a valid body",
			requestBody:    validInput,
			responseStatus: http.StatusCreated,
			expectedError:  nil,
			testResponse: func(in *bytes.Buffer) error {
				got := domain.Port{}
				if err := json.Unmarshal(in.Bytes(), &got); err != nil {
					t.Errorf("failed to unmarshal response %s", err.Error())
				} else {
					if got.ID != validInput.ID {
						t.Errorf("unexpected post result. expected: %s - got: %s", validInput.ID, got.ID)
					}
				}

				return nil
			},
			portUseCase: &port.UseCaseMock{
				CreateFunc: func(_ context.Context, p *domain.Port) error {
					return nil
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := echo.New()
			body := new(bytes.Buffer)
			if err := json.NewEncoder(body).Encode(testCase.requestBody); err != nil {
				t.Error(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/port", body)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			writer := httptest.NewRecorder()
			c := server.NewContext(req, writer)

			handler := handlers.PostPort(testCase.portUseCase)
			err := handler(c)
			if err != nil {
				if testCase.expectedError != nil {
					if err.Error() != testCase.expectedError.Error() {
						t.Errorf("unexpected handler error: %s", err.Error())
					}
				} else {
					t.Errorf("handler error should be nil but got: %s", err.Error())
				}
			}
		})
	}
}
