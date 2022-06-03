package app_todo

import (
	"context"
	"encoding/json"
	"hex-arch-go/core/entities"
	todo "hex-arch-go/core/infrastructure"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// In this case we do a Unit Test for my compression of Unit (Handler + Gateway)
// If we test the Handler, also test the Gateway

var (
	todo_handler = NewToDoHTTPServiceTest()
	ctx          = context.Background()
)

func TestCreateHandler(t *testing.T) {
	testCases := []struct {
		Name         string
		Body         map[string]string
		CodeExpected int
		JsonExpected entities.Response
	}{
		{
			Name: "it should be a success response",
			Body: map[string]string{
				"to": "To",
				"do": "Do",
			},
			CodeExpected: http.StatusAccepted,
			JsonExpected: entities.Response{
				Message: "TODO successfully added.",
				Success: true,
			},
		},
		{
			Name: "it should be a non success response",
			Body: map[string]string{
				"to": "To",
				"do": "Dox",
			},
			CodeExpected: http.StatusOK,
			JsonExpected: entities.Response{
				Message: "The do is invalid.",
				Success: false,
			},
		},
		// More cases with validations and anothers http errors,
		// it depends of your logic implemented on Handler and Gateway
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			f := make(url.Values)
			for name, value := range tc.Body {
				f.Set(name, value)
			}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result := todo_handler.CreateHandler(c)

			assert.NoError(t, result)
			assert.Equal(t, tc.CodeExpected, rec.Code)
			expected, _ := json.Marshal(tc.JsonExpected)
			assert.Equal(t, string(expected)+"\n", rec.Body.String())
		})
	}
}

func TestListHandler(t *testing.T) {
	testCases := []struct {
		Name         string
		CodeExpected int
		JsonExpected []entities.ToDo
	}{
		{
			Name:         "it should be a list of TODO",
			CodeExpected: http.StatusOK,
			JsonExpected: []entities.ToDo{
				{
					ID: "00a95ff5-496c-4c72-87b8-fbc787a53aae",
					To: "To",
					Do: "Do!",
				},
				{
					ID: "e2663c89-e637-445d-aff2-70836d26aee5",
					To: "Too",
					Do: "Doo!",
				},
				{
					ID: "d1826c77-ef4d-47fb-b1a9-24fb0df8551e",
					To: "Tooo",
					Do: "Dooo!",
				},
			},
		},
		// More cases with validations and anothers http errors,
		// it depends of your logic implemented on Handler and Gateway
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			result := todo_handler.ListHandler(c)

			assert.NoError(t, result)
			assert.Equal(t, tc.CodeExpected, rec.Code)
			expected, _ := json.Marshal(tc.JsonExpected)
			assert.Equal(t, string(expected)+"\n", rec.Body.String())
		})
	}
}

func TestNewToDoHTTPServiceTest(t *testing.T) {
	handler := NewToDoHTTPService(ctx, nil)
	var expect *ToDoHTTPService = &ToDoHTTPService{
		todo.NewToDoGateway(ctx, nil),
	}
	assert.Equal(t, handler, expect)
}

// We need a new constructor for inject a Mock
func NewToDoHTTPServiceTest() *ToDoHTTPService {
	// Create storage
	mock := todo.NewToDoMockStorage()

	// Create gateway
	var gateway todo.ToDoGateway = &todo.ToDoLogic{
		St: mock,
	}

	// Create the Fake HTTP Service
	return &ToDoHTTPService{gateway}
}
