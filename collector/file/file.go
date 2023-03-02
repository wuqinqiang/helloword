package file

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/wuqinqiang/helloword/logging"

	"github.com/wuqinqiang/helloword/tools/fx"

	"github.com/wuqinqiang/helloword/dao/model"
)

type File struct {
	list []string
}

func New(files string) *File {
	file := new(File)
	file.list = append(file.list, strings.Split(files, ",")...)
	return file
}

func (f *File) Name() string {
	return "file"
}

func (f *File) Collect(ctx context.Context) (model.Words, error) {
	var words model.Words

	fx.From(func(source chan<- interface{}) {
		for _, file := range f.list {
			source <- file
		}

	}).Walk(func(item interface{}, pipe chan<- interface{}) {
		file, err := os.Open(item.(string))
		if err != nil {
			logging.Errorf("Open file:%s err:%v", item.(string), err)
			return
		}
		defer file.Close() //nolint

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var tmp model.Words
		for scanner.Scan() {
			// etc.wrap [ræp] vt.裹，包，捆 n.披肩
			items := strings.Split(scanner.Text(), " ")
			if len(items) < 3 {
				continue
			}
			word := model.NewWord(items[0])
			word.SetPhonetic(items[1])
			word.SetDefinition(items[2])
			tmp = append(tmp, word)
		}
		if len(tmp) <= 0 {
			return
		}
		pipe <- tmp

	}).ForEach(func(item interface{}) {
		words = append(words, item.(model.Words)...)
	})

	return words, nil
}
