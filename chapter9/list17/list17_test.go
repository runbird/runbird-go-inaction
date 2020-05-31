package main

//测试服务端点
func main() {
	handlers.Routes()
	log.Println("listener: Started:Listening on :4000")
	http.ListenAndServe("4000", nil)

	time.Sleep(1000 * time.Second)
}
