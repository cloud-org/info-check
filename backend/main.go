package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"info-go/server"
	"log"
	"os"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `info-go
Usage: info-go [-h help] [-c ./test.json]
Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&filePath, "c", "./info.json", "数据文件")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	data, err := getData(filePath)
	//log.Printf("data is %+v\n", data)

	engine, err := server.CreateEngine(data)
	if err != nil {
		log.Printf("[main] create engine err: %v", err)
		return
	}

	if err := engine.Start("127.0.0.1:12555"); err != nil {
		log.Printf("[main] start echo server err: %v", err) // 这里不要用 Fatal 不然优雅关停会直接退出
		return
	}
	log.Printf("[main] echo server is start at: %v", ":12555")

}

func getData(filePath string) ([]server.PersonInfo, error) {
	filePtr, err := os.Open(filePath)
	if err != nil {
		log.Printf("文件打开失败 [Err:%s]", err.Error())
		return nil, err
	}
	defer filePtr.Close()
	var res []server.PersonInfo
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&res)
	if err != nil {
		log.Printf("解码失败, err: %v\n", err.Error())
		return nil, err
	} else {
		fmt.Println("解码成功")
		return res, nil
	}
}
