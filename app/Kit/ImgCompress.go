package Kit

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"
)

// CompressImg 压缩图片
// _imgName图片的名字+格式；
// _filepath图片存放的绝对路径文件夹；
// _saveFilepath和新生成的绝对路径文件夹，/结尾。
/*
	// 压缩图片示例
	filepath := Common.ServerInfo["storage_path"] + "upload/" // 在服务器里面的绝对路径文件夹
	dateFile := Common.GetTimeDate("Cmp-Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := Common.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹=[%v]\n", err)
		}
	}
	// 压缩图片
	img := Kit.CompressImg("color.png", filepath, saveFilepath)
*/
func CompressImg(_imgName string, _filepath string, _saveFilepath string)  (backImgName string) {
	has, _ := Common.HasFile(_filepath+_imgName)
	if has == false {
		backImgName = "null-img"
	}else {
		backImgName = Compress(_imgName, _filepath, _saveFilepath)
	}
	return
}

// Compress filepath + img.png
func Compress(imgName string, filepath string, saveFilepath string) (defaultImgName string) {
	var err error
	var file *os.File
	var imgFile string = filepath + imgName
	var width int // 图片动态宽度
	var suffix string // 格式后缀
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg)|(jpeg))$`)

	if !reg.MatchString(imgFile) {
		fmt.Println("压缩图片格式仅限[png, jpg, jpeg]")
		return imgName
	}
	if file, err = os.Open(imgFile); err != nil {
		fmt.Println(err)
		return "null1"
	}
	defer file.Close()
	name := file.Name()
	var img image.Image
	switch {
		case strings.HasSuffix(name, ".png"):
			suffix = "png"
			if img, err = png.Decode(file); err != nil {
				return "null2"
			}
		case strings.HasSuffix(name, ".jpg"):
			suffix = "jpg"
			if img, err = jpeg.Decode(file); err != nil {
				return "null3"
			}
		case strings.HasSuffix(name, ".jpeg"):
			suffix = "jpeg"
			if img, err = jpeg.Decode(file); err != nil {
				return "null4"
			}
		default:
			fmt.Printf("图片名字不正确=%s", name)
			return "null5"
	}

	// 图片的宽
	imgWidth := uint(img.Bounds().Dx()) // px

	// 图片新名字
	newImgName := newImgName()

	// 制作4种尺寸的图
	for p:=1; p<=4; p++ {

		var x string // 缩放倍率，[1,2,3,4]
		x = Common.IntToString(int64(p))

		switch p {
			case 4:
				width = int(imgWidth)
			case 3:
				width = 800
			case 2:
				width = 640
			case 1:
				width = 320
			default:
				width = int(imgWidth)
		}

		resizeImg := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		newImg := newImgName + x + "_" + suffix + "_gl." + suffix
		//fmt.Println(newImg, x, width)
		if x == "4" { // 默认返回原图
			defaultImgName = newImg
		}
		if outFile, err := os.Create(saveFilepath + newImg); err != nil {
			fmt.Println(err)
			return "null6"
		} else {
			defer outFile.Close()
			err = jpeg.Encode(outFile, resizeImg, nil)
			if err != nil {
				fmt.Println(err)
				return "null7"
			}
		}

	}

	return
}

// 重新起名
func newImgName() string {
	return fmt.Sprintf(Common.GetTimeDate("YmdHis")+"_"+Common.RandString(Common.RandRange(13, 15))+"_x")
}