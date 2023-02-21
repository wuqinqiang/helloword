package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/wuqinqiang/helloword/games"

	"github.com/wuqinqiang/helloword/dao"

	"github.com/urfave/cli/v2"
)

var GamesCmd = &cli.Command{
	Name:  "games",
	Usage: "import your own words",
	Subcommands: []*cli.Command{
		wordChainCmd,
	},
}

var wordChainCmd = &cli.Command{
	Name: "chain",
	Action: func(cctx *cli.Context) error {

		list, err := dao.NewWord().GetList(cctx.Context)
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return errors.New("please import some word data first")
		}
		startWord := list.RandomPick()
		chain := games.NewWordChain(list, startWord)

		fmt.Println("Game start")
		fmt.Println("Start word:", chain.StartWord().Word, "  ", chain.StartWord().Definition)

		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print("> ")
			scanner.Scan()
			word := scanner.Text()

			if word == "" {
				fmt.Println("Invalid word. Game over!")
				break
			}

			if word == "exit" {
				break
			}

			if !strings.HasSuffix(chain.PrevWord().Word, word[0:1]) {
				fmt.Println("Invalid word. Game over!")
				break
			}

			if !chain.SetPrevWord(word) {
				fmt.Println("the word has already been used. Game over!")
				break
			}

			nextWord, ok := chain.Pick()
			if !ok {
				fmt.Println("Congratulations, you win!")
				break
			}
			fmt.Println("Next word:", nextWord.Word, "  ", nextWord.Definition)
		}
		return nil
	},
}
