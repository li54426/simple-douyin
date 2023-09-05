package service

import (
	 "bytes"
	 "fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
    "github.com/disintegration/imaging"
	"os"
    "strings"
    )
// ----------------------------TO DO------------------
// 抽取某一帧作为图片
func Vedio2Jpeg(inFileName string, frameNum int) {
	buf := bytes.NewBuffer(nil)

	// 使用 ffmpeg 命令行工具提取视频的指定帧作为 JPEG 图像
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return
	}

	index := strings.LastIndex(inFileName, ".")
	// var outFileName strings
	if index != -1 {
		outFileName := strings.Join([]string{inFileName[:index+1], "jpeg"}, "")
    fmt.Print(outFileName)
		err = imaging.Save(img, outFileName)
		if err != nil {
			return 
		} else {
			return
		}
	}
}