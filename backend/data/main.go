package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"info-go/server"
	"log"
	"os"
	"strings"
)

func main() {
	data := getData()
	// 写入 json
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func getData() []server.PersonInfo {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	res := make([]server.PersonInfo, 0)
	cellList := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	for index := 2; index <= 25; index++ {
		onePerson := make([]string, len(cellList))
		for cellIndex, cell := range cellList {
			value, err := f.GetCellValue("Sheet2", fmt.Sprintf("%s%d", cell, index))
			if err != nil {
				log.Println("获取值发生错误,", err)
				value = ""
			}
			//fmt.Println("value is ", value)
			onePerson[cellIndex] = strings.TrimSpace(value)
		}
		res = append(res, server.PersonInfo{
			Name:     onePerson[0],
			Gender:   onePerson[1],
			Id:       onePerson[2],
			Phone:    onePerson[3],
			Comment:  onePerson[4],
			Room:     onePerson[5],
			Amount:   onePerson[6],
			FamilyId: onePerson[7],
		})
	}

	for i := 0; i < len(res); i++ {
		log.Printf("%+v\n", res[i])
	}
	return res
}
