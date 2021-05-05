package Common

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Log 自定义终端打印日志
func Log(txt interface{}) {
	if Config["debug"] == true {
		log.Println(txt)
	}
}

// EncodeURL 转义url或转义其他字符
func EncodeURL(_url string) string {
	return url.QueryEscape(_url)
}

// DecodeURL 解义url
func DecodeURL(_url string) (string, error) {
	return url.QueryUnescape(_url)
}

// ValueInterfaceToString interface转string，非map[string]interface{}
func ValueInterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// ValueInterfaceToInt interface转int，map[string]interface{}
func ValueInterfaceToInt(_value interface{}) int64 {
	return StringToInt(ValueInterfaceToString(_value))
}

// MapInterfaceToString interface转string，针对map[string]interface{}的某个键
func MapInterfaceToString(_map map[string]interface{}, _key string) string {
	value := _map[_key].(string)
	return value
}

// ArrayInterfaceToString interface转string，准对一维数组[]string{}或[]int{}
func ArrayInterfaceToString(_array interface{}) string {
	value := fmt.Sprintf("%v", _array)
	return value
}

// StringToInt string转int
func StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil { // 报错则默认返回0
		_int = 0
		//fmt.Println("格式转换错误，默认为0。")
		//fmt.Println(err)
	}
	return _int
}

// IntToString int转string
func IntToString(_int int64) string {
	_str := strconv.FormatInt(_int,10)
	return _str
}

// RandRange 获取指定范围内的可变随机整数数，正负都行
func RandRange(_min int64, _max int64) int64 {
	var _rand int64
	if _min >= _max {
		_rand = 0
	}else {
		rand.Seed(time.Now().UnixNano())
		_rand = rand.Int63n(_max - _min) + _min
	}
	return _rand
}

// RandString 生成指定长度的字符串
func RandString(_length int64) string {
	var length int64
	if _length >= 1 {
		length = _length
	}else {
		length = 1
	}
	str := "0123456789-abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(length); i++ {
		result = append(result, bytes[r.Int63n(int64(len(bytes)))])
	}
	return string(result)
}

// MakePaging 生成分页数据
// (数据总条数，每页多少条数据，当前第几页)
// 首页1、上一页N-1、N-2、N-1、本页N、N+1、N+2、下一页N+1、最后一页
func MakePaging(_total int, _limit int, _page int) map[string]interface{}{
	// 总页数
	pageTotal := int(math.Ceil(float64(_total / _limit)))
	if pageTotal < 1 {
		pageTotal = 1
	}
	// 第一页
	firstPage := 1
	// 最后一页
	lastPage := pageTotal
	// 上一页
	beforePage := _page
	if beforePage < 1 {
		beforePage = 1
	}
	// 当前页
	nowPage := _page + 1
	// 下一页
	afterPage := _page + 2
	// 数字分页：1，2，3，4，当前，6，7，8，9
	footPage := ""
	footLen := 4
	// 前4页
	for a:=0; a<footLen; a++ {
		p := nowPage - footLen + a
		if p < nowPage && p >= 1 {
			footPage = footPage + IntToString(int64(p)) + ","
		}
	}
	//footPage = footPage + IntToString(int64(nowPage)) + ","
	// 后4页
	for b:=0; b<footLen; b++ {
		p := nowPage + b
		if p >= 1 && p <= pageTotal {
			footPage = footPage + IntToString(int64(p)) + ","
		}
	}

	var back = map[string]interface{}{
		"total":       _total,
		"page":        _page+1,
		"limit":       _limit,
		"calc": map[string]interface{}{
			"total_page":  pageTotal,
			"first_page":  firstPage,
			"last_page":   lastPage,
			"before_page": beforePage,
			"now_page": nowPage,
			"after_page":  afterPage,
			"foot_page": footPage,
		},
	}
	return back
}

// GetTimeDate 获取日期时间戳，s
func GetTimeDate(_format string) string {
	// 时区
	timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	//timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timer := time.Now().In(timeZone)

	var year int64 = int64(timer.Year())
	var month int64 = int64(timer.Month())
	var day int64 = int64(timer.Day())
	var hour int64 = int64(timer.Hour())
	var minute int64 = int64(timer.Minute())
	var second int64 = int64(timer.Second())

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
	case "Y-m-d":
		_date = _year1 + "-" + _month + "-" + _day
		break
	case "H:i:s":
		_date = _hour + ":" + _minute + ":" + _second
		break
	default:
		_date = _year + "" + _month + "" + _day + "" + _hour + "" + _minute + "" + _second
		break
	}

	return _date
}

// GetTimeS 获取秒时间戳
func GetTimeS() int64 {
	// 时区
	timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	//timeZone := time.FixedZone("CST", 8*3600) // 东八区

	return time.Now().In(timeZone).Unix()
}

// GetTimeMS 获取毫秒时间戳，ms
func GetTimeMS() int64 {
	// 时区
	timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	//timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timeNS := time.Now().In(timeZone).UnixNano() // 纳秒
	timeMS := math.Floor(float64(timeNS / 1000000))
	return int64(timeMS)
}

// DateToTimeS 日期时间戳转时间戳，s
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

// TimeSToDate 秒时间戳转日期，ms
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

// DateToDate 将日期时间戳YmdHis转成日期时间戳Y-m-d H:i:s
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

// FilterToLower 将html标签大写转小写
func FilterToLower(html string) string {
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = reg.ReplaceAllStringFunc(html, strings.ToLower)
	return html
}

// FilterIframe 过滤iframe
func FilterIframe(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<iframe[\\S\\s]+?\\</iframe\\>")
	html = reg.ReplaceAllString(html, "<p class='style'></p>")
	return html
}

// FilterXML 过滤xml
func FilterXML(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<?xml[\\S\\s]+?\\?\\>")
	html = reg.ReplaceAllString(html, "<p class='xml'></p>")
	return html
}

// FilterStyle 过滤html中的style
func FilterStyle(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = reg.ReplaceAllString(html, "<p class='style'></p>")
	return html
}

// FilterJS 过滤html中的js
func FilterJS(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = reg.ReplaceAllString(html, "<p class='js'></p>")
	return html
}

// ReplaceString 全部替换字符串中的某词
func ReplaceString(text string, _old string, _new string) string {
	if len(text) == 0 {
		return ""
	}
	if len(_old) == 0 {
		return text
	}
	if len(_new) == 0 {
		_new = "「小嘴抹了蜜」"
	}
	text = strings.Replace(text, _old, _new, -1)
	return text
}

// ReplaceRangeString 替换字符串几位到几位
func ReplaceRangeString(text string,_start int, _end int, _new string) string {
	if len(text) <= _end {
		_end = len(text)-1
	}
	if len(_new)==0 {
		_new = "**"
	}
	return text[:_start] + _new + text[_end:]
}

// ShuffleArray 打乱数组(字符串型数组)
func ShuffleArray(strings []string) string {
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < len(strings); i++ {
		str += strings[i]
	}
	return str
}

// HasFile 判断文件或文件夹是否存在
func HasFile(filePath string) (bool, string) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, filePath
	}else {
		return false, "FileChecker:::NotFound " + filePath
	}
}

// RequestGet GET请求
func RequestGet(requestUrl string)  {
	//
	
}