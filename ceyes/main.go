package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

type ceye struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		RemoteAddr string `json:"remote_addr"`
		CreatedAt  string `json:"created_at"`
	} `json:"data"`
}
type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	USER        = "zhangjieer2022@163.com" //发送邮件的邮箱
	PASSWORD    = "KZJGOWKDBHVMSCVA"       //发送邮件邮箱的密码
)

// 新建邮件
func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}
func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg
			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		<-done
	}
	return nil
}

//获取ceye的id
func gCeye(uri string) map[string]string {
	var Dic map[string]string
	Dic = make(map[string]string) //dic用来接收获取到的所有ID及values
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res ceye
	_ = json.Unmarshal(body, &res)
	for _, v := range res.Data {
		Dic[v.ID] = v.Name
	}
	return Dic
}

//读取指定目录下的json文件并转换为map
func readJson(filepath string) (conf map[string]interface{}) {
	file, err_f := os.Open(filepath)
	if err_f != nil {
		os.Create("./info.json")
		return
	}
	defer file.Close()
	conf = make(map[string]interface{})
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := json.NewDecoder(file).Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return
}

//字典合并
func MergeMap(mObj ...map[string]string) map[string]string {
	newObj := map[string]string{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

//转字典
func f(m map[string]interface{}) map[string]string {
	ret := make(map[string]string, len(m))
	for k, v := range m {
		ret[k] = fmt.Sprint(v)
	}
	return ret
}

func main() {
	token := "6cf8d5bb7be0f5553842b2bc356dc1e5"
	dicEye := gCeye("http://api.ceye.io/v1/records?token=" + token + "&type=dns&filter=1")
	dicFile := readJson("./info.json")
	for ek, ev := range dicEye {
		if _, ok := dicFile[ek]; ok {
			//别问,问就是抄的。
		} else if ev != "1b2c91.ceye.io" {
			var tDic map[string]string
			tDic = make(map[string]string) //dic用来接收新的内容信息
			tDic[ek] = ev
			//将老文件及刚接收到内容进行合并
			sum := MergeMap(f(dicFile), tDic)
			file, _ := os.OpenFile("./info.json", os.O_WRONLY|os.O_CREATE, 0666)
			defer file.Close()
			//创建encoder 数据输出到file中
			encoder := json.NewEncoder(file)
			//把dataMap的数据encode到file中
			err := encoder.Encode(sum)
			fmt.Println(sum)
			if err != nil {
				fmt.Println(err)
				return
			}
			mycontent := "沐竹里、沐主里，DNSLOG~\r\nID：" + ek + "\r\nValue：" + ev + "\r\nSawa？"
			email := NewEmail("whilewhile@qq.com;",
				"[+] DnsLog 安全告警", mycontent)
			SendEmail(email)
		}
	}

	// filePtr, err := os.Create("info.json")
	// if err != nil {
	// 	fmt.Println("文件创建失败", err.Error())
	// 	return
	// }
	// defer filePtr.Close()
	// // 创建Json编码器
	// encoder := json.NewEncoder(filePtr)
	// err = encoder.Encode(dicEye)
	// if err != nil {
	// 	fmt.Println("EncodeError", err.Error())
	// }
	// if ID != "" || ID != gLocalId() {
	// 	mycontent := "木注里、木注里，DNSLOG cut cut，ID=" + "6cf8d5bb7be0f5553842b2bc356dc1e5"
	// 	email := NewEmail("whilewhile@qq.com;",
	// 		"[+] 安全告警", mycontent)
	// 	err := SendEmail(email)
	// 	fmt.Println(err)
	// }
}
