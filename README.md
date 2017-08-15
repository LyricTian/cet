# cet

> 四六级成绩查询，通过学信网提供的查询接口

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url]

## 下载并使用

``` bash
$ go get -u github.com/LyricTian/cet
```

``` go
result, err := cet.Query(nil, "370150162100108", "张三")
if err != nil {
    panic(err)
}

fmt.Println("姓名：", result.Name)
fmt.Println("学校：", result.University)
fmt.Println("考试级别：", result.Level)
fmt.Println("笔试准考证号：", result.WrittenTicket)
fmt.Println("总分：", result.Score)
fmt.Println("听力：", result.Listening)
fmt.Println("阅读：", result.Reading)
fmt.Println("写作和翻译：", result.WritingTranslation)
fmt.Println("口试准考证号：", result.OralTicket)
fmt.Println("口试等级：", result.OralLevel)
```

## 使用命令行工具

```
$ go get github.com/LyricTian/cet/cmd/cet
```

```
$ cet -name '张三' -ticket '370150162100108'
```

## MIT License

```
Copyright (c) 2016 LyricTian
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/LyricTian/fuh
[ReportCard-Image]: https://goreportcard.com/badge/github.com/LyricTian/fuh
[GoDoc-Url]: https://godoc.org/github.com/LyricTian/fuh
[GoDoc-Image]: https://godoc.org/github.com/LyricTian/fuh?status.svg