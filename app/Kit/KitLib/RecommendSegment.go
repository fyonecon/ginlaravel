package KitLib
// 生成分词词组

import (
	"ginvel.com/app/Common"
	"github.com/go-ego/gse"
	"strings"
)

// MakeSeg 生成文章的分词
func MakeSeg(_text string) []string {
	// 只需要纯文本
	_text = Common.FilterXML(_text) // 过滤xml
	_text = Common.FilterHTML(_text) // 过滤html

	text := []byte(_text)

	// 分词插件1（最短分词）
	//var seg sego.Segmenter
	//seg.LoadDictionary(Common.ServerInfo["storage_path"] + "sys_file/seg/data/dictionary.txt")
	//segments := seg.Segment(text)
	//wordSeg := sego.SegmentsToString(segments, false) // 支持普通模式和搜索模式两种分词，false普通模式

	// 分词插件2（最短分词）
	var seg gse.Segmenter
	seg.LoadDict(Common.ServerInfo["storage_path"] + "sys_file/gse/data/dict/zh/dict.txt")
	seg.LoadStop()
	//seg.LoadDictEmbed()
	//seg.LoadStopEmbed()
	wordSeg := seg.String(string(text), false) // 支持普通模式和搜索模式两种分词，false普通模式

	// 过滤分词
	wordArray := strings.Split(wordSeg, " ")
	wordArray = Common.RemoveRepeatedStringArray(wordArray)
	wordArray = FilterLenStringArray(wordArray, 6)
	wordArray = FilterValueStringArray(wordArray, "的/uj")
	wordArray = FilterValueStringArray(wordArray, "得/ud")

	//wordArray = FilterValueStringArray(wordArray, " /x")
	//wordArray = FilterValueStringArray(wordArray, "./x")
	//wordArray = FilterValueStringArray(wordArray, ",/x")
	//wordArray = FilterValueStringArray(wordArray, ";/x")
	//wordArray = FilterValueStringArray(wordArray, "。/x")
	//wordArray = FilterValueStringArray(wordArray, "，/x")
	//wordArray = FilterValueStringArray(wordArray, "\n/x")
	//wordArray = FilterValueStringArray(wordArray, "、/x")
	//wordArray = FilterValueStringArray(wordArray, "?/x")
	//wordArray = FilterValueStringArray(wordArray, "？/x")
	//wordArray = FilterValueStringArray(wordArray, "!/x")
	//wordArray = FilterValueStringArray(wordArray, "！/x")
	//wordArray = FilterValueStringArray(wordArray, ":/x")
	//wordArray = FilterValueStringArray(wordArray, "：/x")

	return wordArray
}


// FilterValueStringArray 删除数组中不能要的值
func FilterValueStringArray(arr []string, filterValue string) (newArr []string)  {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		theValue := arr[i]
		if theValue != filterValue && len(filterValue) > 0 {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// FilterLenStringArray 删除数组中不能要的值
func FilterLenStringArray(arr []string, _len int) (newArr []string)  {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		theValue := arr[i]
		if len(theValue) >= _len && _len > 0 {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
