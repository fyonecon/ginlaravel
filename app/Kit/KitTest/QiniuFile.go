package KitTest

// 上传本地绝对路径的文件
func QNUploadFile(_file string) (int64, string, string, string)  {
	var state int64
	var msg string
	var filename string
	var domain string

	return state, msg, filename, domain
}

// 上传图片
// png、jpg、gif
// (文件绝对路径，默认输出压缩度照片等级)
func QNUploadImage(_file string, _compress string) (int64, string, string, string)  {
	var state int64
	var msg string
	var filename string
	var domain string

	if len(_compress) == 0 {
		_compress = "x3"
	}

	return state, msg, filename, domain
}

// 上传url图片
func QNUploadUrlImage(_file string, _compress string) (int64, string, string, string)  {
	var state int64
	var msg string
	var filename string
	var domain string

	if len(_compress) == 0 {
		_compress = "x3"
	}

	return state, msg, filename, domain
}

// ----------------------------------------------

// 压缩图片
func QNCompressImage(_file string) string {

	return ""
}

// 统一生成文件新名
func QNNewName(_filename string) string {
	return ""
}