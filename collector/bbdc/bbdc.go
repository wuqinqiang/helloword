package bbdc

import (
	"context"
	"strconv"
	"time"

	"github.com/wuqinqiang/helloword/logging"

	"github.com/wuqinqiang/helloword/tools/fx"

	"github.com/wuqinqiang/helloword/dao/model"

	. "github.com/wuqinqiang/helloword/tools"
)

var newWordsURI = "https://bbdc.cn/api/user-new-word"

type BBDC struct {
	cookie string
}

func New(cookie string) *BBDC {
	return &BBDC{
		cookie: cookie,
	}
}

func (b *BBDC) Name() string {
	return "BBDC" // 不背单词啦
}

func (b *BBDC) Collect(ctx context.Context) (words model.Words, err error) {
	//get the first page
	resp, err := b.request(ctx, 0)
	if err != nil {
		return nil, err
	}
	if err = resp.Ok(); err != nil {
		return
	}
	words = append(words, model.NewWords(resp.GetWords())...)

	// end of page
	if resp.End() {
		return
	}

	// total pagesize
	totalPage := resp.TotalPage()

	fx.From(func(source chan<- interface{}) {
		for i := 1; i < totalPage; i++ {
			source <- i
		}
	}).Walk(func(item interface{}, pipe chan<- interface{}) {
		resp, err := b.request(ctx, item.(int))
		if err != nil {
			logging.Errorf("[BBDC] request page:%d,err:%v", item.(int), err)
			return
		}

		if err = resp.Ok(); err != nil {
			logging.Errorf("[BBDC] request page errmsg :%d,err:%v", item.(int), err)
			return
		}

		pipe <- resp.GetWords()
	}).ForEach(func(item interface{}) {
		words = append(words, model.NewWords(item.([]string))...)
	})

	return
}

func (b *BBDC) request(ctx context.Context, page int) (resp *Response, err error) {
	now := time.Now()
	resp = new(Response)
	_, err = Resty.R().SetHeader("Cookie", b.cookie).
		SetHeader("Accept", "application/json").
		SetQueryParam("page", strconv.Itoa(page)).
		SetQueryParam("time", strconv.Itoa(int(now.Unix()))).
		SetResult(&resp).
		Get(newWordsURI)
	return
}
