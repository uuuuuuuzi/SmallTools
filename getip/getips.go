package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"time"
)

type result struct {
	Origin string `json:"origin"`
}

func getip(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res result
	_ = json.Unmarshal(body, &res)
	return res.Origin
}

func DoWriteFile(_filePath string) error {
	nowtime := time.Now()
	n1 := nowtime.Format("2006-01-02 15:04")

	_file, _err := os.OpenFile(_filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if _err != nil {
		return _err
	}
	//提前关闭文件
	defer _file.Close()
	//写入文件
	_writer := bufio.NewWriter(_file)
	_writer.WriteString("[-] " + n1 + " - " + getip("http://httpbin.org/ip") + "\n")
	_writer.Flush()
	return nil
}

// 执行cmd
func cmder(cmd string) error {
	cmdv := exec.Command(cmd)
	cmdv.Stdout = os.Stdout
	return cmdv.Run()
}

// 判断是否存在启动项
func checkFile(path string) {
	_, err := os.Stat(path)
	// if os.IsExist(err) {
	// 	fmt.Println("6666")
	// } else {
	// 	os.Mkdir(path, os.ModePerm)
	// 	_, err := os.Create(path + "ips.log")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	if err == nil {
		_, errr := os.Stat(path + "ip.log")
		// 判断ip.log是否存在
		if os.IsNotExist(errr) {
			_, err := os.Create(path + "ip.log")
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
		_, err := os.Create(path + "ip.log")
		if err != nil {
			fmt.Println(err)
		}
	}

}
func eWinMsconfig() {
	//

}
func eMacMsconfig(path string) {
	Crotab := `3 * * * * /Users/iii/.getip/getip`
	fmt.Println(Crotab)
}

// 一些帮助函数,支持-s检索日志文件该IP是否出现过。
func ehelps() {

}

/*
生成日志文件后将当天的日期写到第一行
每隔5分钟写入IP
支持检索IP是否出现过
*/
func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	checkFile(u.HomeDir + "/.getip/")

	// DoWriteFile(u.HomeDir + "/.getip/ip.log")

	//fmt.Println(cmder("ifconfig"))

}
