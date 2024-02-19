package sms

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)
import "fmt"

// 向手机发送验证码
type SMSInfo struct {
	AccessKey       string `json:"accessKey"`
	AccessSecretKey string `json:"accessSecretKey"`
}

var smsInfo map[string]SMSInfo

func init() {
	yamlFile, e := os.ReadFile("config/config.json")
	if e != nil {
		log.Fatal(e)
	}
	e = json.Unmarshal(yamlFile, &smsInfo)
	if e != nil {
		log.Fatal(e)
	}
}
func SendMsg(tel string, code string) string {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", smsInfo["sms"].AccessKey, smsInfo["sms"].AccessSecretKey)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel             //手机号变量值
	request.SignName = "JeromeBlog"        //签名
	request.TemplateCode = "SMS_292455087" //模板编码
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err := client.SendSms(request)
	fmt.Println(response.Code)
	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		return "frequency_limit"
	}
	if err != nil {
		fmt.Print(err.Error())
		return "failed"
	}
	return "success"
}

// 随机验证码

func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}
