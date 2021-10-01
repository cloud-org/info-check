package main

import (
	"encoding/json"
	"fmt"
	"info-go/server"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
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
	cellList := []string{"B", "C", "D", "E", "F"}
	for index := 3; index <= 24; index++ {
		onePerson := make([]string, len(cellList))
		for cellIndex, cell := range cellList {
			value, err := f.GetCellValue("Sheet1", fmt.Sprintf("%s%d", cell, index))
			if err != nil {
				log.Println("获取值发生错误,", err)
				value = ""
			}
			//fmt.Println("value is ", value)
			onePerson[cellIndex] = value
		}
		res = append(res, server.PersonInfo{
			Name:     onePerson[0],
			Gender:   onePerson[1],
			Id:       onePerson[2],
			Phone:    onePerson[3],
			Comment:  onePerson[4],
			Room:     "",
			FamilyId: "",
		})
	}

	for i := 0; i < len(res); i++ {
		log.Printf("%+v\n", res[i])
	}
	return res
}
