package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
	"log"
)

func searchInFile(filePath, keyword string) bool {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("无法读取文件 %s: %v\n", filePath, err)
		return false
	}

	// 查找关键字
	if strings.Contains(string(content), keyword) {
		return true
	}
	return false
}

func findFilesAndSearch(root, keyword string, extensions []string) {
	// 遍历目录及子目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 错误处理
		if err != nil {
			log.Printf("访问路径 %s 时出错: %v\n", path, err)
			return err
		}

		// 仅处理指定的扩展名文件
		for _, ext := range extensions {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
				// 如果文件包含指定的关键词，输出文件路径
				if searchInFile(path, keyword) {
					fmt.Println(path)
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("遍历目录时出错: %v\n", err)
	}
}

func main() {
	// 检查命令行参数是否提供了关键词
	if len(os.Args) < 2 {
		fmt.Println("请提供要搜索的关键字")
		os.Exit(1)
	}

	// 获取命令行输入的关键词
	keyword := os.Args[1]

	// 写死的搜索目录
	rootDir := "/Users/iii/Documents/GitHub/notepad"  // 这里指定你要搜索的目录

	// 支持的文件扩展名
	extensions := []string{".md", ".yml"}

	// 在指定路径下查找包含关键字的 .md 和 .yml 文件
	findFilesAndSearch(rootDir, keyword, extensions)
}
