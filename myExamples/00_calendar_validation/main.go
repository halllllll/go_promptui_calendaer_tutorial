package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {

	var years []int
	for i := 2020; i < 2100; i++ {
		years = append(years, i)
	}
	months := []string{"1月 - Jan", "2月 - Feb", "3月 - Mar", "4月 - Apr", "5月 - May", "6月 - Jun", "7月 - Jul", "8月 - Aug", "9月 - Sep", "10月 - Oct", "11月 - Nov", "12月 - Dec"}
	hours := []int{}
	for i := 0; i < 24; i++ {
		hours = append(hours, i)
	}
	minutes := []int{}
	for i := 0; i < 60; i++ {
		minutes = append(minutes, i)
	}

	var fromUnixTime, toUnixTime int64

	fmt.Println("どこから？")
	for {
		now := time.Now()
		curYearIdx := sort.SearchInts(years, now.Year())
		fromYearPrompt := promptui.Select{
			Label:     "年  Select Year",
			Items:     years,
			CursorPos: curYearIdx,
			Size:      1,
		}
		fromYearIdx, fromYearResult, err := fromYearPrompt.Run()

		if err != nil {
			panic(err)
		}

		fromMonthPrompt := promptui.Select{
			Label:     "月 Select Month",
			Items:     months,
			Size:      1,
			CursorPos: 0,
		}

		fromMonthIdx, _, err := fromMonthPrompt.Run()

		if err != nil {
			panic(err)
		}

		// 年、月から妥当な日を生成
		fromDays := []int{}
		for i := 1; i <= 31; i++ {
			d := strconv.Itoa(i)
			_, err := time.Parse("2006 1 2", fmt.Sprintf("%s %d %s", fromYearResult, fromMonthIdx+1, d))
			if err != nil {
				break
			}
			fromDays = append(fromDays, i)
		}
		fromDayPrompt := promptui.Select{
			Label:     "日 Select Day",
			Items:     fromDays,
			CursorPos: now.Day() + 1,
		}

		_, fromDayResult, err := fromDayPrompt.Run()

		if err != nil {
			panic(err)
		}

		fromHourPrompt := promptui.Select{
			Label:     "時間 Select Hour",
			Items:     hours,
			CursorPos: 0,
		}

		fromHourIdx, _, err := fromHourPrompt.Run()

		if err != nil {
			panic(err)
		}

		fromMinutesPrompt := promptui.Select{
			Label:     "分 Select Minute",
			Items:     minutes,
			CursorPos: 0,
		}

		fromMinuteIdx, _, err := fromMinutesPrompt.Run()

		fromYear, _ := strconv.ParseInt(fromYearResult, 10, 64)
		fromDay, _ := strconv.ParseInt(fromDayResult, 10, 64)

		fromTime := time.Date(int(fromYear), time.Month(fromMonthIdx)+1, int(fromDay), hours[fromHourIdx], minutes[fromMinuteIdx], 0, 0, time.FixedZone("Asia/Tokyo", 0))
		if err != nil {
			panic(err)
		}
		fromConfirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("FROM: %d/%d/%d %d:%d からでいいですか？", fromTime.Year(), fromTime.Local().Month(), fromTime.Day(), fromTime.Hour(), fromTime.Local().Minute()),
			IsConfirm: true,
		}
		yes, err := fromConfirmPrompt.Run()
		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}
		if yes == "y" {
			years = years[fromYearIdx:]
			fromUnixTime = fromTime.Unix()
			break
		} else {
			continue
		}
	}
	fmt.Println("いつまで？")
	for {
		now := time.Now()
		curYearIdx := sort.SearchInts(years, now.Year())
		toYearPrompt := promptui.Select{
			Label:     "年  Select Year",
			Items:     years,
			CursorPos: curYearIdx,
			Size:      1,
		}
		_, toYearResult, err := toYearPrompt.Run() //入力を受け取る

		if err != nil {
			panic(err)
		}

		toMonthPrompt := promptui.Select{
			Label:     "月 Select Month",
			Items:     months,
			Size:      1,
			CursorPos: 0,
		}

		toMonthIdx, _, err := toMonthPrompt.Run()

		if err != nil {
			panic(err)
		}

		// 年、月から妥当な日を生成
		toDays := []int{}
		for i := 1; i <= 31; i++ {
			d := strconv.Itoa(i)
			_, err := time.Parse("2006 1 2", fmt.Sprintf("%s %d %s", toYearResult, toMonthIdx+1, d))
			if err != nil {
				break
			}
			toDays = append(toDays, i)
		}
		toDayPrompt := promptui.Select{
			Label:     "日 Select Day",
			Items:     toDays,
			CursorPos: now.Day() + 1,
		}

		_, toDayResult, err := toDayPrompt.Run()

		if err != nil {
			panic(err)
		}

		toHourPrompt := promptui.Select{
			Label:     "時間 Select Hour",
			Items:     hours,
			CursorPos: now.Hour(),
		}

		toHourIdx, _, err := toHourPrompt.Run()

		if err != nil {
			panic(err)
		}

		toMinutesPrompt := promptui.Select{
			Label:     "分 Select Minute",
			Items:     minutes,
			CursorPos: now.Minute(),
		}

		toMinuteIdx, _, err := toMinutesPrompt.Run()

		toYear, _ := strconv.ParseInt(toYearResult, 10, 64)
		toDay, _ := strconv.ParseInt(toDayResult, 10, 64)

		toTime := time.Date(int(toYear), time.Month(toMonthIdx)+1, int(toDay), hours[toHourIdx], minutes[toMinuteIdx], 0, 0, time.FixedZone("Asia/Tokyo", 0))
		if err != nil {
			panic(err)
		}

		toUnixTime = toTime.Unix()

		if fromUnixTime >= toUnixTime {
			fmt.Println("!! 過去は選べないよ !!")
			continue
		}

		toConfirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("TO: %d/%d/%d %d:%d まででいいですか？", toTime.Year(), toTime.Local().Month(), toTime.Day(), toTime.Hour(), toTime.Local().Minute()),
			IsConfirm: true,
		}
		yes, err := toConfirmPrompt.Run()
		if err != nil {
			fmt.Printf("failed %v\n", err)
			return
		}
		if yes == "y" {
			break
		} else {
			continue
		}
	}

}
