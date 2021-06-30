package Common
// 解决一些原生Go方法没「有默认值」而直接panic的问题；「从晦涩到好用」。
// 用Go解释Go

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Log 自定义终端打印日志
func Log(txt ...interface{}) {
	if Config["log_debug"] == true || Config["log_debug"] == "true" {
		log.Println("Log：", txt)
	}
}
func Error(txt ...interface{}) {
	if Config["log_debug"] == true || Config["log_debug"] == "true" {
		log.Println("Error：", txt)
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

// GetUrlParam 获取url中的参数（非解码）
func GetUrlParam(_url string, _key string) (value string) {
	u, err := url.Parse(_url)
	values := u.Query()
	if err != nil {
		value = ""
	}else {
		value = values.Get(_key)
	}
	return
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

// StringToFloat string转float
func StringToFloat(_str string) float64 {
	_float, err := strconv.ParseFloat(_str, 64) // string转int
	if err != nil { // 报错则默认返回0
		_float = 0.0
		//fmt.Println("格式转换错误，默认为0。")
		//fmt.Println(err)
	}
	return _float
}

// FloatToString float转string
func FloatToString(_float float64) string {
	_str := strconv.FormatFloat(_float, 'e', 10, 64)
	return _str
}

// RandRange 获取指定范围内的可变随机整数数，正负都行。[a, b]
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
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
// Y年m月d号 H:i:s.MS.NS 星期W
func GetTimeDate(_format string) (date string) {
	if len(_format) == 0 {
		_format = "YmdHisMS"
	}
	date = _format

	// 时区
	//timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timer := time.Now().In(timeZone)

	var year int64 = int64(timer.Year())
	var month int64 = int64(timer.Month())
	var day int64 = int64(timer.Day())
	var hour int64 = int64(timer.Hour())
	var minute int64 = int64(timer.Minute())
	var second int64 = int64(timer.Second())
	var week int64 = int64(timer.Weekday())
	var ms int64 = int64(timer.UnixNano() / 1e6)
	var ns int64 = int64(timer.UnixNano() / 1e9)
	msTmp := IntToString(int64(math.Floor(float64(ms/1000))))
	nsTmp := IntToString(int64(math.Floor(float64(ns/1000000))))

	var _year string
	var _month string
	var _day string
	var _hour string
	var _minute string
	var _second string
	var _week string // 英文星期
	var _Week string // 中文星期
	var _ms string // 毫秒
	var _ns string // 纳秒

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
	_week = IntToString(week)
	WeekZh := [...]string{"日", "一", "二", "三", "四", "五", "六"} // 默认从"日"开始
	_Week = WeekZh[week]
	_ms = strings.Replace(IntToString(ms), msTmp, "", -1)
	_ns = strings.Replace(IntToString(ns), nsTmp, "", -1)

	// 替换关键词
	date = strings.Replace(date, "MS", _ms, -1)
	date = strings.Replace(date, "NS", _ns, -1)
	date = strings.Replace(date, "Y", _year, -1)
	date = strings.Replace(date, "m", _month, -1)
	date = strings.Replace(date, "d", _day, -1)
	date = strings.Replace(date, "H", _hour, -1)
	date = strings.Replace(date, "i", _minute, -1)
	date = strings.Replace(date, "s", _second, -1)
	date = strings.Replace(date, "W", _Week, -1)
	date = strings.Replace(date, "w", _week, -1)

	return
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

	timer := time.Now().In(timeZone)
	timeMS := int64(timer.UnixNano() / 1e6)
	return timeMS
}

// GetTimeNS 获取纳秒时间戳，ms
func GetTimeNS() int64 {
	// 时区
	timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	//timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timer := time.Now().In(timeZone)
	timeNS := int64(timer.UnixNano() / 1e9)
	return timeNS
}

// DateToTimeS 秒日期时间戳转时间戳，s
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

// TimeSToDate 秒时间戳转秒日期，ms
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

// GetBeforeTime 获取_day天前的秒时间戳、日期时间戳
// _day为负则代表取前几天，为正则代表取后几天，0则为今天
func GetBeforeTime(_day int) (int64, string) {
	// 时区
	timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	//timeZone := time.FixedZone("CST", 8*3600) // 东八区
	// 前n天
	nowTime := time.Now().In(timeZone)
	beforeTime := nowTime.AddDate(0, 0, _day)
	// 时间转换格式
	beforeTimeS := beforeTime.Unix() // 秒时间戳
	beforeDate := TimeSToDate(beforeTimeS, "YmdHis") // 日期时间戳
	return beforeTimeS, beforeDate
}

// GetDatesBetweenDay 计算两个日期之间有多少天间隔
// 支持YmdHis、Y-m-d H:i:s格式日期，startDate开始日期，endDate结束日期
// ±天数取决于开始时间和结束时间谁大
func GetDatesBetweenDay(startDate string, endDate string, format string) (day int64) {
	// 日期转秒时间戳
	startTime := DateToTimeS(startDate, format)
	endTime := DateToTimeS(endDate, format)
	// 获取日期秒之差
	dayTime := endTime - startTime
	// 时分秒将被忽略，只取天的部分
	day = int64(math.Floor(float64(dayTime / (24 * 60 * 60))))
	return
}

// FilterToLower 将html标签大写转小写
func FilterToLower(html string) string {
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = reg.ReplaceAllStringFunc(html, strings.ToLower)
	return html
}

// FilterHTML 过滤html标签、去除\n\t\r\n。
func FilterHTML(html string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = re.ReplaceAllString(html, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = re.ReplaceAllString(html, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllString(html, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	html = re.ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
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
	html = reg.ReplaceAllString(html, " ")
	return html
}

// FilterStyle 过滤html中的style
func FilterStyle(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = reg.ReplaceAllString(html, " ")
	return html
}

// FilterJS 过滤html中的js
func FilterJS(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = reg.ReplaceAllString(html, " ")
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

// FilterInput 过滤Input输入的值
// 转义%、"、'、(、)、!、/、^、*、.、
func FilterInput(_value string) string {
	value := _value

	blackArray := [...]string{ // 这些符号将被转义
		"%", "(",")", "!", "/", "^", "*", ".", "|", "=",
	}
	changArray := [...]string{ // 这些符号将被替代
		"select", "delete", "char", "insert", "count", "exec", "declare", "update",
	}

	for i:=0; i<len(blackArray); i++ {
		txt := blackArray[i]
		value = ReplaceString(value, txt, EncodeURL(txt))
	}
	for j:=0; j<len(changArray); j++ {
		txt := changArray[j]
		value = ReplaceString(value, txt, "_" + txt)
	}

	value = html.EscapeString(value) // xss

	return value
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

// MakeRedisKey 生成一个Redis的键
func MakeRedisKey(_array interface{}) string {
	key := ArrayInterfaceToString(_array)
	key = EncodeURL(key)
	return key
}

// MapInterfaceToJson map[string]interface{}转Json
func MapInterfaceToJson(_map map[string]interface{}) []byte {
	_json, _ := json.Marshal(_map)
	return _json
}

// JsonStringToMap String格式的json转Map，并不能直接转Json
func JsonStringToMap(_string string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(_string), &data)
	if err == nil {
		fmt.Println(data)
	}
	return data
}

// MD5 生成md5
func MD5(_string string) string {
	h := md5.New()
	io.WriteString(h, _string)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// EncodeBase64 生成base64（标准方式）
func EncodeBase64(_string string) string {
	res := base64.StdEncoding.EncodeToString([]byte(_string))
	return res
}

// DecodeBase64 解密base64（标准方式）
func DecodeBase64(_string string) string {
	res, err := base64.StdEncoding.DecodeString(_string)
	if err != nil {
		fmt.Printf("DecodeBase64 Error: %s ", err.Error())
		return ""
	}
	return string(res)
}

// EncodeUrlBase64 加密文件和url名安全型base64
func EncodeUrlBase64(_string string) string {
	res := base64.URLEncoding.EncodeToString([]byte(_string))
	return res
}

// DecodeUrlBase64 解密文件和url名安全型base64
func DecodeUrlBase64(_string string) string {
	res, err := base64.URLEncoding.DecodeString(_string)
	if err != nil {
		fmt.Printf("DecodeUrlBase64 Error: %s ", err.Error())
		return ""
	}
	return string(res)
}

// RequestGet GET请求
// 参考博客：https://blog.csdn.net/qq_29176323/article/details/109745009
func RequestGet(urlNoParams string, body map[string][]string) string {
	params := url.Values{}
	params = body
	//params.Set("aaa", "aaa") // 设置参数
	parseURL, err := url.Parse(urlNoParams)
	if err != nil {
		log.Println(err)
		return ""
	}
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(res)
}

// RemoveRepeatedStringArray String数组去重
func RemoveRepeatedStringArray(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// RemoveRepeatedIntArray Int数组去重
func RemoveRepeatedIntArray(array []int) []int {
	newArray := make([]int, 0)

	for _, i := range array {
		if len(newArray) == 0 {
			newArray = append(newArray, i)
		} else {
			for k, v := range newArray {
				if i == v {
					break
				}
				if k == len(newArray)-1 {
					newArray = append(newArray,i)
				}
			}
		}
	}
	return newArray
}

// MakeSMSCode 随机生成短信验证码
// _len 验证码多少个数字
func MakeSMSCode(_len int) (code string) {
	_array := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if _len <= 4 {
		_len = 4
	}
	for i:=0; i<_len; i++{
		theNumber := _array[RandRange(0, int64(len(_array)-1))]
		code = code + theNumber
	}
	return
}

// HideStringValue 隐藏/替换字符串中的某些字符
// 如隐藏手机号：185*******6，调用Common.HideStringValue("18512345506", 3, 10, "*")
func HideStringValue(_string string, start int, end int, replaceValue string) string {
	if len(replaceValue) == 0 {
		replaceValue = "*"
	}
	blackString := _string[start:end]
	replace := ""
	for i:=0;i<len(blackString);i++ {
		replace = replace + replaceValue
	}
	return strings.Replace(_string, blackString, replace, -1)
}
