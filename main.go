package main

import (
	"net/http"
	"github.com/mseongj/Weather-reminder/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
