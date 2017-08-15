package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/LyricTian/cet"
)

var (
	ticket = flag.String("ticket", "", "准考证号")
	name   = flag.String("name", "", "姓名")
)

func main() {
	flag.Parse()

	if len(*ticket) == 0 || len(*name) == 0 {
		fmt.Println("请输入准考证号或姓名")
		return
	}

	fmt.Println("开始查询...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := cet.Query(ctx, *ticket, *name)
	if err != nil {
		fmt.Printf("查询发生错误：%s\n", err.Error())
		return
	}

	if result.OralTicket == "" {
		result.OralTicket = "--"
	}

	if result.OralLevel == "" {
		result.OralLevel = "--"
	}

	fmt.Println("查询结果：")
	fmt.Printf("\t姓   名：%s\n", result.Name)
	fmt.Printf("\t学   校：%s\n", result.University)
	fmt.Printf("\t考试级别：%s\n", result.Level)
	fmt.Println("=================笔试成绩=================")
	fmt.Printf("\t准考证号：%s\n", result.WrittenTicket)
	fmt.Printf("\t总   分：%g\n", result.Score)
	fmt.Printf("\t\t听     力：%g\n", result.Listening)
	fmt.Printf("\t\t阅     读：%g\n", result.Reading)
	fmt.Printf("\t\t写作和翻译：%g\n", result.WritingTranslation)
	fmt.Println("=================口试成绩=================")
	fmt.Printf("\t准考证号：%s\n", result.OralTicket)
	fmt.Printf("\t等   级：%s\n", result.OralLevel)
}
