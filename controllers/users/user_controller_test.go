package users

//
//import (
//	"bytes"
//	"encoding/json"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"net/url"
//	"testing"
//
//	"jwt_auth/domain/users"
//
//	"github.com/DATA-DOG/go-sqlmock"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//)
//
//func generateContextAndRecorder(content interface{}) (*gin.Context, *httptest.ResponseRecorder) {
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//	jsonBytes, err := json.Marshal(content)
//	if err != nil {
//		panic(err)
//	}
//	c.Request = &http.Request{
//		Method: http.MethodPost,
//		URL:    &url.URL{},
//		Header: http.Header{"Accept": []string{"application/json"}},
//		Body:   io.NopCloser(bytes.NewBuffer(jsonBytes)),
//	}
//	return c, w
//}
//
//func TestRegister(t *testing.T) {
//	testCases := []struct {
//		testName    string
//		jsonContent *users.User
//		want        int
//	}{
//		{
//			testName: "test ok, json format is correct",
//			jsonContent: &users.User{
//				ID:        1,
//				FirstName: "test",
//				LastName:  "test",
//				Password:  "test",
//				Email:     "test",
//			},
//			want: http.StatusOK,
//		},
//		{
//			testName:    "test not ok, json format is incorrect",
//			jsonContent: nil,
//			want:        http.StatusBadRequest,
//		},
//	}
//	db, _, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	for _, testCase := range testCases {
//		t.Run(testCase.testName, func(t *testing.T) {
//			_, w := generateContextAndRecorder(testCase.jsonContent)
//			Register(db)
//			assert.Equal(t, testCase.want, w.Code)
//		})
//	}
//}
