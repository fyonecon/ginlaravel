package Common

import (
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

// 代码公共函数
func Test(txt string) string {
	if len(txt) == 0 {
		txt = "txt-nil"
	}
	Config["test"] = txt
	return txt
}

// 转义url
func EncodeURL(_url string) string {
	return url.QueryEscape(_url)
}
// 解析转义url
func DecodeURL(_url string) (string, error) {
	return url.QueryUnescape(_url)
}

// string转int
func StringToInt(_str string) int {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil { // 报错则默认返回0
		_int = 0
		fmt.Println("格式转换错误：")
		fmt.Println(err)
	}
	return int(_int)
}

// int转string
func IntToString(_int int) string {
	_str := strconv.FormatInt(int64(_int),10)
	return _str
}

// 获取指定范围内的可变随机整数数，正负都行
func RandRange(_min int, _max int) int {
	var _rand int
	if _min >= _max {
		_rand = 0
	}else {
		rand.Seed(time.Now().UnixNano())
		_rand = rand.Intn(_max - _min) + _min
	}
	return _rand
}

// 生成指定长度的字符串
func RandString(_length int) string {
	var length int
	if _length >= 1 {
		length = _length
	}else {
		length = 1
	}
	str := "0123456789-abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成分页数据
// (数据总条数，每页多少条数据，当前第几页)
// 首页1、上一页N-1、N-2、N-1、本页N、N+1、N+2、下一页N+1、最后一页
func MakePaging(_total int, _limit int, _page int) map[string]int{
	// 总页数
	pageTotal := int(math.Ceil(float64(_total / _limit)))
	// 第一页
	firstPage := 1
	// 最后一页
	lastPage := pageTotal
	// 上一页
	beforePage := _page - 1
	// 下一页
	afterPage := _page + 1

	var back = map[string]int{
		"total":       _total,
		"page":        _page,
		"limit":       _limit,
		"page_total":  pageTotal,
		"first_page":  firstPage,
		"last_page":   lastPage,
		"before_page": beforePage,
		"after_page":  afterPage,
	}
	return back
}

// 获取日期时间戳，s
func GetTimeDate(_format string) string {
	timer := time.Now()

	var year int = timer.Year()
	var month int = int(timer.Month())
	var day int = timer.Day()
	var hour int = timer.Hour()
	var minute int = timer.Minute()
	var second int = timer.Second()

	var _year string
	var _month string
	var _day string
	var _hour string
	var _minute string
	var _second string

	_year = IntToString(year)
	if month < 10 {
		_month = IntToString(month)
		_month = "0" + _month
	}else {
		_month = IntToString(month)
	}
	if day < 10 {
		_day = IntToString(day)
		_day = "0" + _day
	}else {
		_day = IntToString(day)
	}
	if hour < 10 {
		_hour = IntToString(hour)
		_hour = "0" + _hour
	}else {
		_hour = IntToString(hour)
	}
	if minute < 10 {
		_minute = IntToString(minute)
		_minute = "0" + _minute
	}else {
		_minute = IntToString(minute)
	}
	if second < 10 {
		_second = IntToString(second)
		_second = "0" + _second
	}else {
		_second = IntToString(second)
	}

	_year1 := IntToString(year)
	_month1 := IntToString(month)
	_day1 := IntToString(day)
	_hour1 := IntToString(hour)
	_minute1 := IntToString(minute)
	_second1 := IntToString(second)

	var _date string

	switch _format {
	case "YmdHis":
		_date = _year + "" + _month + "" + _day + "" + _hour + "" + _minute + "" + _second
		break
	case "Y-m-d H:i:s":
		_date = _year + "-" + _month + "-" + _day + " " + _hour + ":" + _minute + ":" + _second
		break
	case "y-m-d h:i:s":
		_date = _year1 + "-" + _month1 + "-" + _day1 + " " + _hour1 + ":" + _minute1 + ":" + _second1
		break
	default:
		_date = _year + "" + _month + "" + _day + "" + _hour + "" + _minute + "" + _second
		break
	}

	return _date
}

// 获取秒时间戳
func GetTimeS() int64 {
	return time.Now().Unix()
}

// 获取毫秒时间戳，ms
func GetTimeMS() int64 {
	timeNS := time.Now().UnixNano() // 纳秒
	timeMS := math.Floor(float64(timeNS / 1000000))
	return int64(timeMS)
}

// 日期时间戳转时间戳，s
func DateToTimeS(_date string, format string) int64 {
	var date string
	if len(_date) == 0 { //给一个默认值
		date = GetTimeDate("YmdHis")
	}else {
		date = _date
	}

	var layout string
	if format == "YmdHis" || format == "" {
		layout = "20060102150405" // 转化所需内定模板
	}else if format == "Y-m-d H:i:s" {
		layout = "2006-01-02 15:04:05"
	}else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	}else {
		layout = "20060102150405"
	}

	//日期转化为时间戳
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(layout, date, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}

// 秒时间戳转日期，ms
func TimeSToDate(_timeS int64, format string) string {
	var timeS int64
	if _timeS == 0 { //给一个默认值
		timeS = GetTimeS()
	}else {
		timeS = _timeS
	}

	var layout string
	if format == "YmdHis" || format == "" {
		layout = "20060102150405" // 转化所需内定模板
	}else if format == "Y-m-d H:i:s" {
		layout = "2006-01-02 15:04:05"
	}else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	}else {
		layout = "20060102150405"
	}

	date := time.Unix(timeS, 0).Format(layout)
	return date
}

// 将日期时间戳YmdHis转成日期时间戳Y-m-d H:i:s
func DateToDate(_date string) string {
	var date string
	if len(_date) == 0 {
		date = GetTimeDate("YmdHis")
	}else {
		date = _date
	}

	timeS := DateToTimeS(date, "")

	return TimeSToDate(timeS, "Y-m-d H:i:s")
}

// GET请求
func RequestGet(requestUrl string)  {
	//
	
}