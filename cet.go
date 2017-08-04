package cet

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	query Querier = &querier{}
)

// Query 根据准考证号、姓名查询四六级成绩
func Query(ctx context.Context, ticket, name string) (*Result, error) {
	return query.Query(ctx, ticket, name)
}

// Querier 提供四六级成绩查询的接口
type Querier interface {
	// 根据准考证号、姓名查询
	Query(ctx context.Context, ticket, name string) (*Result, error)
}

// Result 查询结果
type Result struct {
	Name               string  // 姓名
	University         string  // 学校
	Level              string  // 考试级别
	WrittenTicket      string  // 笔试准考证号
	Score              float32 // 总分
	Listening          float32 // 听力
	Reading            float32 // 阅读
	WritingTranslation float32 // 写作和翻译
	OralTicket         string  // 口试准考证号
	OralLevel          string  // 口试等级
}

const (
	// 四六级查询的网址
	cetHost = "www.chsi.com.cn"
)

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// NewQuerier 创建Querier接口
func NewQuerier() Querier {
	return &querier{}
}

type querier struct {
}

func (q *querier) Query(ctx context.Context, ticket, name string) (result *Result, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	u := &url.URL{
		Scheme: "http",
		Host:   cetHost,
		Path:   "cet/query",
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return
	}

	uq := req.URL.Query()
	uq.Set("zkzh", ticket)
	uq.Set("xm", name)
	req.URL.RawQuery = uq.Encode()

	req.Header.Set("Referer", fmt.Sprintf("http://%s/cet/", cetHost))
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36")
	req.Header.Set("X-Forwarded-For", q.randomIP())

	err = q.httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		r, err := q.parse(resp)
		if err != nil {
			return err
		}
		result = r

		return nil
	})

	return
}

var (
	rd = rand.New(rand.NewSource(math.MaxInt64))
)

func (q *querier) randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rd.Intn(255), rd.Intn(255), rd.Intn(255), rd.Intn(255))
}

func (q *querier) httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func (q *querier) parse(resp *http.Response) (result *Result, err error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return
	}

	childs := doc.Find(".m_cnt_m").Children()
	if !childs.HasClass("cetTable") {
		err = ErrNotFound
		return
	}

	result = new(Result)
	scoreType := 0 // 成绩类型：1笔试成绩 2口试成绩

	childs.Find(".cetTable tr").Each(func(i int, s *goquery.Selection) {
		val := strings.TrimSpace(s.Find("td").Last().Text())
		if val == "--" {
			return
		}

		text := strings.TrimSpace(s.Find("th").Text())
		switch text {
		case "姓 名：":
			result.Name = val
		case "学 校：":
			result.University = val
		case "考试级别：":
			result.Level = val
		case "笔试成绩":
			scoreType = 1
		case "准考证号：":
			if scoreType == 1 {
				result.WrittenTicket = val
			} else if scoreType == 2 {
				result.OralTicket = val
			}
		case "总 分：":
			fv, _ := strconv.ParseFloat(val, 32)
			result.Score = float32(fv)
		case "听 力：":
			fv, _ := strconv.ParseFloat(val, 32)
			result.Listening = float32(fv)
		case "阅 读：":
			fv, _ := strconv.ParseFloat(val, 32)
			result.Reading = float32(fv)
		case "写作和翻译：":
			fv, _ := strconv.ParseFloat(val, 32)
			result.WritingTranslation = float32(fv)
		case "口试成绩":
			scoreType = 2
		case "等 级：":
			result.OralLevel = val
		}
	})

	return
}
