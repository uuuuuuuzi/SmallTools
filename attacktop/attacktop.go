package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

// 排序，别问我为什么这样写，百度抄的。
func map_Sort(mp map[string]int32, nn int) {
	var lstPerson []peroson
	for k, v := range mp {
		lstPerson = append(lstPerson, peroson{k, v})
	}
	sort.Slice(lstPerson, func(i, j int) bool {
		return lstPerson[i].Age > lstPerson[j].Age // 降序
	})
	n := 0
	for _, v := range lstPerson {
		fmt.Println(v)
		n += 1
		if n == nn {
			break
		}
	}
}

type peroson struct {
	Name string
	Age  int32
}

func main() {
	// flags排序后的数据
	flags := map[string]int32{}
	// values所有数据
	values := []string{}

	var num int
	flag.IntVar(&num, "top", 10, "Attack top value,Example:-top [num]")
	Names := flag.String("filename", "ip.log", "Input Your Filename")
	flag.Parse()

	fi, err := os.Open(*Names)
	if err == nil {
		defer fi.Close()

		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			values = append(values, string(a))
		}
	}
	for _, v := range values {
		if flags[v] == 0 {
			flags[v] += 1
		} else {
			flags[v] += 1
		}
	}

	map_Sort(flags, num)

}
