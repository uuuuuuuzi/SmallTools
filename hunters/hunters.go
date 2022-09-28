package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type gResult struct {
	Code int `json:"code"`
	Data struct {
		AccountType string `json:"account_type"`
		Total       int    `json:"total"`
		Time        int    `json:"time"`
		Arr         []struct {
			IsRisk         string `json:"is_risk"`
			URL            string `json:"url"` // 提取值
			IP             string `json:"ip"`  // 提取值
			Port           int    `json:"port"`
			WebTitle       string `json:"web_title"` // 提取值
			Domain         string `json:"domain"`
			IsRiskProtocol string `json:"is_risk_protocol"`
			Protocol       string `json:"protocol"`
			BaseProtocol   string `json:"base_protocol"`
			StatusCode     int    `json:"status_code"`
			Component      []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"component"`
			Os        string `json:"os"`
			Company   string `json:"company"`
			Number    string `json:"number"`
			Country   string `json:"country"`
			Province  string `json:"province"`
			City      string `json:"city"`
			UpdatedAt string `json:"updated_at"`
			IsWeb     string `json:"is_web"`
			AsOrg     string `json:"as_org"`
			Isp       string `json:"isp"`
			Banner    string `json:"banner"`
		} `json:"arr"`
		ConsumeQuota string `json:"consume_quota"`
		RestQuota    string `json:"rest_quota"`
		SyntaxPrompt string `json:"syntax_prompt"`
	} `json:"data"`
	Message string `json:"message"`
}

type MainConfig struct {
	Key string
}

// load config
func loadConf() string {
	// 打开文件
	str, _ := os.Getwd()
	file, _ := os.Open(str + "/conf.json")

	// 关闭文件
	defer file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)

	conf := MainConfig{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf.Key
}

// Excel处理
func gExcel(ips, urls, number []string, filename string) {
	file := xlsx.NewFile()
	style := xlsx.NewStyle()
	var row *xlsx.Row
	var cell *xlsx.Cell
	sheet, err := file.AddSheet("Hunters")
	style.Alignment = xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "序号"
	cell = row.AddCell()
	cell.Value = "IP"
	cell = row.AddCell()
	cell.Value = "URL"
	cell = row.AddCell()
	cell.Value = "备案号"

	for i := 0; i < len(ips); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = strconv.Itoa(i + 1)
		cell = row.AddCell()
		cell.Value = ips[i]
		cell = row.AddCell()
		cell.Value = urls[i]
		cell = row.AddCell()
		cell.Value = number[i]
	}

	err = file.Save(filename + ".xlsx")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var tSyntax string
	flag.StringVar(&tSyntax, "s", "", "Hunter语法,如:ip=\"123.56.83.203\"")
	flag.Parse()
	if tSyntax == "" {
		fmt.Println("请通过-s输入查询语句!")
		return
	}
	var sIP = []string{}
	var sURL = []string{}
	var sNumber = []string{}

	key := loadConf()

	syntax := base64.URLEncoding.EncodeToString([]byte(tSyntax))
	page := 50
	for i := 0; i < page; i++ {
		sTime := "2000-01-21+00%3A00%3A00"
		uri := "https://hunter.qianxin.com/openApi/search?api-key=" + key + "&search=" + syntax + "&page=" + strconv.Itoa(i+1) + "&page_size=100&is_web=1&status_code=200&start_time=" + sTime + "&end_time=2099-08-19+23%3A59%3A59"
		fmt.Println(uri)
		// 因为爬快了Hunter会掉数据,所以延迟3
		time.Sleep(5)
		request, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Println("NewRequest Error:", err)
			return
		}

		client := &http.Client{
			// 取消https认证
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}

		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Do Request Error:", err)
			return
		}
		defer response.Body.Close()

		result, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("ReadAll Error:", err)
			return
		}
		// 通过构造体获取到所有json数据
		var gr gResult
		err = json.Unmarshal(result, &gr)
		for n, keys := range gr.Data.Arr {
			fmt.Println(keys.IP, keys.URL, keys.Number)
			sIP = append(sIP, gr.Data.Arr[n].IP)
			sURL = append(sURL, gr.Data.Arr[n].URL)
			sNumber = append(sNumber, gr.Data.Arr[n].Number)
		}
		time.Sleep(2)
		if int(float64(gr.Data.Total/100))+1 == i+1 {
			// 根据需求调用生成Excel
			// 结束
			t1 := time.Now()
			timeStamp := t1.Unix()
			t := time.Unix(int64(timeStamp), 0) // time.Time

			gExcel(sIP, sURL, sNumber, t.Format("2006-01-02_15_04"))
			fmt.Println("[+] ", i+1, "页, 总共:", gr.Data.Total, ":条数据!")
			fmt.Println("[+] -= Game Over =- [+]")
			return
		}
	}
}
