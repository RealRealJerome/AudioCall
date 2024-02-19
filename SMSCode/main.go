package main

import (
	"fmt"
	"github.com/RealRealJerome/AudioCall/sms"
	"log"
	"net/http"
)

func SmsCodeRpc(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	tel := queryParams.Get("phoneNum")
	randCode := sms.Code()
	res := sms.SendMsg(tel, randCode)
	// 上限
	if res == "frequency_limit" {
		_, err := fmt.Fprint(w, "服务已达上限，请稍后请求验证码")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	// 请求失败
	if res == "failed" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "请求验证码失败：服务器内部错误")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, tel+":"+randCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "请求验证码失败：服务器内部错误")
		return
	}
}
func main() {
	// 监听接口
	http.HandleFunc("/sms", SmsCodeRpc) // 设置访问的路由
	fmt.Println("started successfully")
	if err := http.ListenAndServe(":6666", nil); err != nil {
		log.Fatal(err)
	}
}
