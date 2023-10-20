package main

import (
	"C"

	"github.com/laoflch/goav/avfilter"
	"github.com/laoflch/goav/avlaoflch"
)
import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"strconv"
)

func main() {

	//var logo_ctx **avfilter.Context

	//logo_ctx := new(*avfilter.Context)
	//logo_ctx2 := new(avfilter.Context)
	//println((*logo_ctx2).name)
	//println(*logo_ctx)
	logoImageFile, err := os.Open("/workspace/ffmpeg/FFmpeg/doc/examples/laoflch-mc-log.png")
	//logoImageFile, err := os.Open("/workspace/go/src/github.com/laoflch/goav/new_iamge.png")

	if err != nil {

	}

	logoImage, err := png.Decode(logoImageFile)

	if err != nil {

	}

	var frame *avfilter.Frame

	fmt.Println(reflect.TypeOf(logoImage))

	//logoImage2 := ImageTypeToRGBA64(&logoImage)
	switch imageColor := (logoImage).(type) {
	case *image.NRGBA:

		frame = avfilter.AvFrameAlloc()

		//if frame != nil {

		//frame.data[0] = unsafe.Pointer(&(imageColor.Pix))
		//frame.linesize[0] = imageColor.Stride
		frame.SetFrameFormat(26)

		frame.SetFrameSize(450, 85)

		//frame.SetPts(1)

		//frame2:=(*avfilter.Frame)frame

		avfilter.SetPictureNRGBA(frame, imageColor)

		//avfilter.SetPictureNRGBA64(frame, imageColor)
		fmt.Println(frame.GetFrameFormat())

	}

	//}
	ctx_Object := reflect.TypeOf(*frame)

	//ctx_value := reflect.ValueOf(**logo_ctx)

	//fmt.Println((**logo_ctx))

	for i := 0; i < ctx_Object.NumField(); i++ {
		//if ctx_Object.Field(i).Name == "name" {
		fmt.Printf("%s:%v:=\n", ctx_Object.Field(i).Name, ctx_Object.Field(i).Type)

		//fmt.Printf("%v\n", ctx_value.Field(i).Interface())

		//}
	}

	ret_chan := make(chan int)

	go func() {
		defer func() {

			//output_w = nil

		}()
		info := avlaoflch.VideoHandleProcessInfoAlloc()
		//sret := avlaoflch.PushVideo2RtspSubtitleLogo("/workspace/ffmpeg/c-ffmpeg/Se7en.1995.REMASTERED.1080p.BluRay.x264-SADPANDA.mkv", 0, 1, "/workspace/ffmpeg/c-ffmpeg/Se7en.1995.REMASTERED.1080p.BluRay.x264-SADPANDA.zh.ass", &frame, "rtsp://127.0.0.1:554/test", true, true, uint64(480), uint64(480*6), uint64(480*3), &info)
		sret := avlaoflch.PushVideo2RtspSubtitleLogo("http://192.168.71.247:8200/MediaItems/22167.mkv", 0, 1, "", &frame, "rtsp://127.0.0.1:554/test", true, true, uint64(480), uint64(480*6), uint64(480*3), &info)

		//	LOG.Warningln("do handle fail:", err.Error())

		ret_chan <- sret

	}()
	//time.Sleep(time.Duration(10) * time.Second)

	//for i := 2400 - 1; i > 0; i-- {

	//time.Sleep(time.Duration(10*10) * time.Millisecond)
	//	imageRbga64 := ImageTypeToRGBA64(&logoImage)

	//fmt.Printf("prec rate:%f", float64(i)/10)
	//	newRgb := OpactiyAdjust(imageRbga64, float64(i)/2400)
	//newRgb := OpactiyAdjust(imageRbga64, 0.8)

	//	avfilter.SetPictureRGBA64(frame, newRgb)
	//}
	//SaveImage("./new_iamge.png", newRgb)

	//return

	//fmt.Println(reflect.TypeOf(logoImage))
	//switch imageColor := (logoImage).(type) {
	//case *image.NRGBA:
	//println("########################################################################3")
	//frame := avfilter.AvFrameAlloc()

	//if frame != nil {

	//frame.data[0] = unsafe.Pointer(&(imageColor.Pix))
	//frame.linesize[0] = imageColor.Stride

	//frame.SetFrameFormat(26)

	//frame.SetFrameSize(450, 85)

	//frame2:=(*avfilter.Frame)frame

	//avfilter.SetPictureNRGBA(frame, imageColor)
	//	println((*(*avfilter.Context)(*logo_ctx)).GetName())

	//var int64 pts = 1

	/*for {

		//frame.

	/*	ret := (*logo_ctx).AvBuffersrcAddFrameFlags(frame, 4)

		fmt.Println("add logo frame", ret)
		//frame = avfilter.AvFrameAlloc()
		frame.SetPts(frame.NextPts())
		//if frame != nil {

		//frame.data[0] = unsafe.Pointer(&(imageColor.Pix))
		//frame.linesize[0] = imageColor.Stride
		//frame.SetFrameFormat(26)

		//frame.SetFrameSize(450, 85)

		//ret = (*logo_ctx).AvBuffersrcAddFrameFlags(nil, 4)
		//avfilter.SetPictureNRGBA(frame, logoImage.(*image.NRGBA))

		time.Sleep(time.Duration(40) * time.Millisecond)
	}
	*/
	//ret = (*logo_ctx).AvBuffersrcAddFrameFlags(nil, 4)

	//fmt.Println("add logo frame", ret)

	//ret = (*logo_ctx).AvBuffersrcClose(frame.NextPts(), 4)
	//fmt.Println("close context", ret)
	//}

	//fmt.Println(imageColor.ColorModel())

	//}

	select {

	case _, ok := <-ret_chan:
		if ok {
			//	return
		}

	}

	//go avlaoflch.PushVideo2RtspSubtitleLogo("/workspace/ffmpeg/c-ffmpeg/Se7en.1995.REMASTERED.1080p.BluRay.x264-SADPANDA.mkv", "/workspace/ffmpeg/c-ffmpeg/Se7en.1995.REMASTERED.1080p.BluRay.x264-SADPANDA.zh.ass", logo_ctx, "rtsp://127.0.0.1:554/test", true)

	//if logo_ctx != nil {
	//ctx_Object := reflect.TypeOf(**logo_ctx)

	//ctx_value := reflect.ValueOf(**logo_ctx)

	//fmt.Println((**logo_ctx))

	//for i := 0; i < ctx_Object.NumField(); i++ {
	//		if ctx_Object.Field(i).Name == "name" {
	//			fmt.Printf("%s:%v:=", ctx_Object.Field(i).Name, ctx_Object.Field(i).Type)

	//fmt.Printf("%v\n", ctx_value.Field(i).Interface())

	// 		}
	//}

	//avcodec.AvcodecFindDecoderByName("h264")
}

func SaveImage(fileName string, m *image.RGBA64) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()

	}()

	if err := png.Encode(f, m); err != nil {
		panic(err)
	}

}

func OpactiyAdjust(m *image.RGBA64, percentage float64) *image.RGBA64 {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	newRegba := image.NewRGBA64(bounds)

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()

			//fmt.Printf("opacity:%b", a)
			//fmt.Printf("float:%3.0f", float64(a)*percentage)
			opacity, _ := strconv.Atoi(fmt.Sprintf("%3.f", float64(a)*percentage))
			//opacity := int8(a * percentage)
			//fmt.Printf(" opacity:%b", uint(opacity))

			//fmt.Printf(" opacity:%b \n", uint16(opacity))
			v := newRegba.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(opacity)})
			rr, gg, bb, aa := v.RGBA()

			newRegba.SetRGBA64(i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
		}
	}

	return newRegba
}

func ImageTypeToRGBA64(m *image.Image) *image.RGBA64 {

	bounds := (*m).Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	newRgba := image.NewRGBA64(bounds)
	for i := 0; i < dx; i++ {

		for j := 0; j < dy; j++ {
			colorRgb := (*m).At(i, j)
			r, g, b, a := colorRgb.RGBA()
			nR := uint16(r)
			nG := uint16(g)
			nB := uint16(b)
			alpha := uint16(a)
			newRgba.SetRGBA64(i, j, color.RGBA64{R: nR, G: nG, B: nB, A: alpha})
		}

	}

	return newRgba

}
