package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"net/http"
	"os"
)

// PrintLogo prints a simple ASCII art logo
func PrintLogo() {
	logo := `
	 _   _ ____  _  __
	| | | |  _ \| |/ /
	| | | | |_) | ' / 
	| |_| |  __/| . \ 
	 \___/|_|   |_|\_\
	`
	fmt.Println(logo)
}

func main() {
	// 解析命令行参数
	var targetIP string
	var pageSize int
	var pageNum int
	var showHelp bool

	flag.StringVar(&targetIP, "c", "", "Specify the IP address for encoding in base64")
	flag.IntVar(&pageSize, "n", 10, "Specify the number of items per page (allowed values: 10, 20, 50, 100)")
	flag.IntVar(&pageNum, "p", 1, "Specify the page number")
	flag.BoolVar(&showHelp, "h", false, "Show help message")

	flag.Parse()

	// 显示Logo
	PrintLogo()

	// 显示帮助信息
	if showHelp {
		flag.PrintDefaults()
		return
	}

	// 检查pageSize参数的合法性
	validPageSizes := map[int]bool{10: true, 20: true, 50: true, 100: true}
	if !validPageSizes[pageSize] {
		fmt.Println("Invalid page size. Allowed values: 10, 20, 50, 100.")
		os.Exit(1)
	}

	// 检查是否提供了-c参数
	if targetIP == "" {
		fmt.Println("Please provide the target IP address using the -c parameter.")
		os.Exit(1)
	}

	// 从配置文件读取API key
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Error reading configuration file: %v\n", err)
		os.Exit(1)
	}

	section := cfg.Section("api")
	apiKey := section.Key("api-key").String()

	// 对目标IP进行base64编码
	queryValue := base64.StdEncoding.EncodeToString([]byte(targetIP))

	apiURL := fmt.Sprintf("https://api.hunter.how/search?api-key=%s&query=%s&start_time=2023-01-01&end_time=2023-12-01&page=%d&page_size=%d", apiKey, queryValue, pageNum, pageSize)

	// 发送GET请求
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer response.Body.Close()

	// 解析JSON响应
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 提取域名信息
	if code, ok := result["code"].(float64); ok && code == 200 {
		data, ok := result["data"].(map[string]interface{})
		if ok {
			list, ok := data["list"].([]interface{})
			if ok {
				for _, item := range list {
					if itemMap, ok := item.(map[string]interface{}); ok {
						domain, ok := itemMap["domain"].(string)
						if ok {
							port := 0
							if portFloat, ok := itemMap["port"].(float64); ok {
								port = int(portFloat)
							}
							fmt.Printf("%s:%d\n", domain, port)
						}
					}
				}
			}
		}
		fmt.Printf("Total items: %d\n", int(data["total"].(float64)))
	} else {
		fmt.Println("API request failed. Code:", result["code"])
	}
}
