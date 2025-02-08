package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  time.Time
}

func indexHeandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreateAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHeandler)

	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})

	return mux
}

type WeatherData struct {
	Response struct {
		Header struct {
			ResultCode    string `json:"resultCode"`
			ResultMsg     string `json:"resultMsg"`
		} `json:"header"`

		Body struct {
			Items struct {
				Item []struct {
					BaseDate  string `json:"baseDate"`
					BaseTime  string `json:"baseTime"`
					Category  string `json:"category"`
					FcstDate  string `json:"fcstDate"`
					FcstTime  string `json:"fcstTime"`
					FcstValue string `json:"fcstValue"`
					Nx        string `json:"nx"`
					Ny        string `json:"ny"`
				} `json:"item"`
			} `json:"items"`
		} `json:"body"`
	} `json:"response"`
}

const (
	API_KEY    = "발급받은_인증키"
	NX         = "73"   // 강원도 춘천시 x좌표
	NY         = "134"  // 강원도 춘천시 y좌표
	API_URL    = "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getUltraSrtFcst"
)


func fetchWeather(now time.Time) map[string]string {
	baseDate := now.Format("20060102")
	baseTime := fmt.Sprintf("%02d00", now.Hour()-1)

	resp, err := http.Get(fmt.Sprintf("%s?serviceKey=%s&pageNo=1&numOfRows=1000&dataType=JSON&base_date=%s&base_time=%s&nx=%s&ny=%s",
		API_URL, API_KEY, baseDate, baseTime, NX, NY))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data WeatherData
	json.Unmarshal(body, &data)

	weather := make(map[string]string)
	for _, item := range data.Response.Body.Items.Item {
		if item.FcstTime == fmt.Sprintf("%02d00", now.Hour()) {
			switch item.Category {
			case "T1H":
				weather["temp"] = item.FcstValue + "°C"
			case "REH":
				weather["humidity"] = item.FcstValue + "%"
			case "SKY":
				weather["sky"] = map[string]string{
					"1": "☀️ 맑음",
					"3": "⛅ 구름많음",
				}[item.FcstValue]
			}
		}
	}
	return weather
}
