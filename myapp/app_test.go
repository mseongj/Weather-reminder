package myapp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=moon", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello moon!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"seongjae", "last_name":"moon", "email":"sample@gmail.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("seongjae", user.FirstName)
	assert.Equal("moon", user.LastName)

}
func TestFetchWeather(t *testing.T) {
	assert := assert.New(t)

	// Mock the HTTP response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/1360000/VilageFcstInfoService_2.0/getUltraSrtFcst", r.URL.Path)
		assert.Equal("발급받은_인증키", r.URL.Query().Get("serviceKey"))
		assert.Equal("73", r.URL.Query().Get("nx"))
		assert.Equal("134", r.URL.Query().Get("ny"))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"response": {
				"header": {
					"resultCode": "00",
					"resultMsg": "NORMAL_SERVICE"
				},
				"body": {
					"items": {
						"item": [
							{
								"baseDate": "20231010",
								"baseTime": "0900",
								"category": "T1H",
								"fcstDate": "20231010",
								"fcstTime": "1000",
								"fcstValue": "20",
								"nx": "73",
								"ny": "134"
							},
							{
								"baseDate": "20231010",
								"baseTime": "0900",
								"category": "REH",
								"fcstDate": "20231010",
								"fcstTime": "1000",
								"fcstValue": "60",
								"nx": "73",
								"ny": "134"
							},
							{
								"baseDate": "20231010",
								"baseTime": "0900",
								"category": "SKY",
								"fcstDate": "20231010",
								"fcstTime": "1000",
								"fcstValue": "1",
								"nx": "73",
								"ny": "134"
							},
							{
								"baseDate": "20231010",
								"baseTime": "0900",
								"category": "PTY",
								"fcstDate": "20231010",
								"fcstTime": "1000",
								"fcstValue": "0",
								"nx": "73",
								"ny": "134"
							}
						]
					}
				}
			}
		}`)
	}))
	defer server.Close()

}
// go install github.com/smartystreets/goconvey@latest
// and past in terminal
// goconvey
