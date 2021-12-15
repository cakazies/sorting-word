package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestSortingMap ...
func TestSortingMap(t *testing.T) {

	tests := []struct {
		name               string
		initRouter         func() (*gin.Context, *httptest.ResponseRecorder)
		expectedStatusCode int
	}{
		{
			name: "error from form data",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				c.Request, _ = http.NewRequest(http.MethodPost, "api/sorting", nil)
				c.Set("paragraf", "4403")
				return c, h
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "paragraf nil",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				apiUrl := "https://api.com/"
				data := url.Values{}
				data.Set("paragraf", "")

				u, _ := url.ParseRequestURI(apiUrl)
				u.Path = "api/sorting"
				urlStr := u.String()
				c, _ := gin.CreateTestContext(h)
				c.Request, _ = http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
				c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				return c, h
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "paragraf nil",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				apiUrl := "https://api.com/"
				data := url.Values{}
				data.Set("paragraf", `at once. Before assigning the duplicate You can assign the object returned by the object to another object, you can change any of the properties of the duplicate
										object without affecting the original. You can assign the object returned by the settings of all the properties of a duplicated ParagraphFormat You can use the Duplicate`)

				u, _ := url.ParseRequestURI(apiUrl)
				u.Path = "api/sorting"
				urlStr := u.String()
				c, _ := gin.CreateTestContext(h)
				c.Request, _ = http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
				c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				return c, h
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c, _ := tt.initRouter()

			SortingMap(c)

			if !reflect.DeepEqual(c.Writer.Status(), tt.expectedStatusCode) {
				t.Errorf("ArticleController.GetArticle got = %v, want %v", c.Writer.Status(), tt.expectedStatusCode)
			}
		})

	}
}
