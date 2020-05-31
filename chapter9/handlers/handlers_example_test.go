package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

//$ go test -v -run="ExampleSendJSON"
//testing: warning: no tests to run
//PASS
//ok      runbird-go-inaction/chapter9/handlers   3.423s

//示例代码以Example开头，包含已经存在的方法 SendJSON
func ExampleSendJSON() {
	request, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, request)

	var u struct {
		Name  string
		Email string
	}
	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("Error", err)
	}

	fmt.Println(u)
}
