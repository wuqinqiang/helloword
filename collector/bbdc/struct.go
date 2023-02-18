package bbdc

import (
	"errors"

	"github.com/wuqinqiang/helloword/dao/model"
)

type Response struct {
	ResultCode  int    `json:"result_code"`
	DataKind    string `json:"data_kind"`
	DataVersion string `json:"data_version"`
	ErrorBody   struct {
		UserMessage string `json:"user_message"`
		Info        string `json:"info"`
	} `json:"error_body"`

	DataBody struct {
		WordList []struct {
			SentenceList []struct {
				Id                 float64 `json:"id"`
				Word               string  `json:"word"`
				WordOriginal       string  `json:"wordOriginal"`
				OriginalContext    string  `json:"originalContext"`
				TranslationContext string  `json:"translationContext"`
				WordLength         float64 `json:"wordLength"`
				WordNum            float64 `json:"wordNum"`
				CourseId           string  `json:"courseId"`
				SentenceId         string  `json:"sentenceId"`
				Url                string  `json:"url"`
				SortLevel          float64 `json:"sortLevel"`
				SortBy             float64 `json:"sortBy"`
			} `json:"sentenceList"`
			Interpret  string `json:"interpret"`
			Ukpron     string `json:"ukpron"`
			Updatetime string `json:"updatetime"`
			Word       string `json:"word"`
			Uspron     string `json:"uspron"`
		} `json:"wordList"`
		PageInfo struct {
			TotalRecord float64 `json:"totalRecord"`
			PageSize    float64 `json:"pageSize"`
			TotalPage   float64 `json:"totalPage"`
			CurrentPage float64 `json:"currentPage"`
		} `json:"pageInfo"`
	} `json:"data_body"`
}

func (resp *Response) Ok() error {
	if resp.ResultCode == 200 {
		return nil
	}
	return errors.New(resp.ErrorBody.UserMessage)
}

func (resp *Response) End() bool {
	return resp.DataBody.PageInfo.CurrentPage == resp.DataBody.PageInfo.TotalPage
}

func (resp *Response) TotalPage() int {
	return int(resp.DataBody.PageInfo.TotalPage)
}

func (resp *Response) GetWords() (words model.Words) {
	items := resp.DataBody.WordList
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		word := model.NewWord(item.Word)
		word.SetDefinition(item.Interpret)
		word.SetPhonetic(item.Ukpron)
		words = append(words, word)
	}
	return
}
