package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	var idx int
	var result string
	preItems := []string{"Heart A", "Spade K", "Dia J", "Club Q"}
	ans := "Dia J"

	for {
		curItems := append([]string{}, preItems...)
		// curItems = append(curItems, s)
		prompt := promptui.Select{
			// 選択肢のタイトル
			Label: "Select",
			// 選択肢の配列
			Items: curItems,
		}

		_idx, _result, err := prompt.Run() //入力を受け取る

		if err != nil {
			panic(err)
		}
		idx = _idx
		result = _result
		if result == ans {
			break
		} else {
			fmt.Println("やりなおし〜")
		}
	}

	fmt.Printf("正解！　脱出成功〜 idx: %d result: %q\n", idx, result)
}
