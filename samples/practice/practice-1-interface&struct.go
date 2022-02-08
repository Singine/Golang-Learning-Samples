package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Report interface {
	writingReport() string
}

type Content struct {
	date string
	weather string
	dayPassing []int
	secondPassing string
}

func (content Content) writingReport() string{
	txt := "咕宝，现在是" + content.date + ",\n天气是" + content.weather + ",\n我已经没有咕宝" + strconv.Itoa(content.dayPassing[0]) + "天" + strconv.Itoa(content.dayPassing[1]) + "小时" + strconv.Itoa(content.dayPassing[2]) + "分钟" + content.secondPassing + "秒了。"
	return txt
}

func calcPassingDay() ([]int,string){
	now := time.Now()
	then := time.Date(2021, 03, 24, 00, 00, 00, 0, time.UTC)
	passingTime := now.Sub(then).String()
	timeData := timeSpliting(passingTime, []rune{'h','m','s'})

	h,_ := strconv.Atoi(timeData[0])
	m,_ := strconv.Atoi(timeData[1])
	s := strings.Split(timeData[2], ".")[0]

	result := [...]int{h/24,h%24,m}

	return result[:],s
}

func timeSpliting(str string, key []rune) []string{
	Split := func(r rune) bool{
		for _,v := range key {
			if v == r{
				return true
			}
		}
		return false
	}
	result := strings.FieldsFunc(str,Split)
	return result
}

func printing(reports []Report){
	for _,report := range reports{
		fmt.Println(report.writingReport())
	}
}

func main() {

	calc1,calc2 := calcPassingDay()

	today := Content{
		date: time.Now().Format("2006-01-02 15:04:05"),
		weather: "有云云但是他们是好好云云不太挡住太阳的还是有阳光的小太阳天咕",
		dayPassing:calc1,
		secondPassing:calc2,
	}
	report := []Report{today}
	printing(report)
}

