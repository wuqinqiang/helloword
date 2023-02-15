package bbdc

import (
	"context"
	. "helloword/tools"
	"strconv"
	"time"
)

var newWordsURI = "https://bbdc.cn/api/user-new-word"

type BBDC struct {
	page   int64
	cookie string
}

func New(cookie string) *BBDC {
	return &BBDC{
		page:   0,
		cookie: cookie,
	}
}

func (b *BBDC) Collect(ctx context.Context) error {
	now := time.Now()
	var resp Response

	for {
		if _, err := Resty.R().SetHeader("Cookie", b.cookie).
			SetHeader("Accept", "application/json").
			SetQueryParam("page", strconv.Itoa(int(b.page))).
			SetQueryParam("time", strconv.Itoa(int(now.Unix()))).
			SetResult(&resp).
			Get(newWordsURI); err != nil {
			return err
		}
		if err := resp.Ok(); err != nil {
			return err
		}
		if resp.End() {
			break
		}
	}
	return nil
}
