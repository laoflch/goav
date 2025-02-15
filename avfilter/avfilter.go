// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
		#include <libavfilter/buffersrc.h>
  #include <libavfilter/buffersink.h>
  #include <libavcodec/avcodec.h>
	#include <libavutil/imgutils.h>
	#include <libavutil/frame.h>
*/
import "C"
import (
	"image"
	"unsafe"
)

type (
	Filter     C.struct_AVFilter
	Context    C.struct_AVFilterContext
	Frame      C.struct_AVFrame
	Link       C.struct_AVFilterLink
	Graph      C.struct_AVFilterGraph
	Input      C.struct_AVFilterInOut
	InOut      C.struct_AVFilterInOut
	Pad        C.struct_AVFilterPad
	Dictionary C.struct_AVDictionary
	Class      C.struct_AVClass
	MediaType  C.enum_AVMediaType
)

const (
	AV_BUFFERSRC_FLAG_PUSH            = C.AV_BUFFERSRC_FLAG_PUSH
	AV_BUFFERSRC_FLAG_KEEP_REF        = C.AV_BUFFERSRC_FLAG_KEEP_REF
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT
)

//Return the LIBAvFILTER_VERSION_INT constant.
func AvfilterVersion() uint {
	return uint(C.avfilter_version())
}

//Return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//Return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//Get the number of elements in a NULL-terminated array of Pads (e.g.
func AvfilterPadCount(p *Pad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//Get the name of an Pad.
func AvfilterPadGetName(p *Pad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Get the type of an Pad.
func AvfilterPadGetType(p *Pad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Link two filters together.
func AvfilterLink(s *Context, sp uint, d *Context, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), ((*C.struct_AVFilterContext)(d)), C.uint(dp)))
}

//Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(l **Link) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//Get the number of channels of a link.
func AvfilterLinkGetChannels(l *Link) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//Set the closed field of a link.
func AvfilterLinkSetClosed(l *Link, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(f *Context) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//Make the filter instance process a command.
func AvfilterProcessCommand(f *Context, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}
func AvFrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

func (frame *Frame) SetFrameFormat(format_type int) {

	frame.format = C.int(format_type)

}

func (frame *Frame) GetFrameFormat() int {

	return int(frame.format)

}
func (frame *Frame) SetFrameSize(w int, h int) {

	frame.width = C.int(w)
	frame.height = C.int(h)

}

func SetPictureNRGBA(f **Frame, img *image.NRGBA) {
	//d := Data(f)
	// l := Linesize(f)
	// FIXME: Save the original pointers somewhere, this is a memory leak
	//d[0] = (*uint8)(unsafe.Pointer(&img.Pix[0]))

	C.av_image_fill_arrays(&((*f).data[0]), &((*f).linesize[0]), (*C.uchar)(&img.Pix[0]), int32((*f).format), (*f).width, (*f).height, 1)
	//fmt.Printf("logo_image len: %d %d \n", len(img.Pix), ret)
	// d[1] = (*uint8)(unsafe.Pointer(&img.Cb[0]))
}
func SetPictureNRGBA64(f *Frame, img *image.NRGBA64) {
	//d := Data(f)
	// l := Linesize(f)
	// FIXME: Save the original pointers somewhere, this is a memory leak
	//d[0] = (*uint8)(unsafe.Pointer(&img.Pix[0]))

	C.av_image_fill_arrays(&(f.data[0]), &(f.linesize[0]), (*C.uchar)(&img.Pix[0]), int32(f.format), f.width, f.height, 1)
	// d[1] = (*uint8)(unsafe.Pointer(&img.Cb[0]))
}
func SetPictureRGBA64(f *Frame, img *image.RGBA64) {
	//d := Data(f)
	// l := Linesize(f)
	// FIXME: Save the original pointers somewhere, this is a memory leak
	//d[0] = (*uint8)(unsafe.Pointer(&img.Pix[0]))

	C.av_image_fill_arrays(&(f.data[0]), &(f.linesize[0]), (*C.uchar)(&img.Pix[0]), int32(f.format), f.width, f.height, 1)
	// d[1] = (*uint8)(unsafe.Pointer(&img.Cb[0]))
}

//Initialize the filter system.
func AvfilterRegisterAll() {
	C.avfilter_register_all()
}

//Initialize a filter with the supplied parameters.
func (ctx *Context) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

func (ctx *Context) GetName() string {

	return C.GoString((*ctx).name)

}
func Data(f *Frame) (data [8]*uint8) {
	for i := range data {
		data[i] = (*uint8)(f.data[i])
	}
	return
}

//Initialize a filter with the supplied dictionary of options.
func (ctx *Context) AvfilterInitDict(o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Free a filter context.
func (ctx *Context) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

func (ctx *Context) AvBuffersrcAddFrameFlags(frame *Frame, flags C.int) int {
	return int(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), (*C.struct_AVFrame)(unsafe.Pointer(frame)), flags))
}

func (ctx *Context) AvBuffersinkGetFrameFlags(frame *Frame, flags C.int) int {
	return int(C.av_buffersink_get_frame_flags((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), (*C.struct_AVFrame)(unsafe.Pointer(frame)), flags))
}

func (ctx *Context) AvBuffersrcWriteFrame(frame *Frame) int {
	return int(C.av_buffersrc_write_frame((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

func (ctx *Context) AvBuffersrcAddFrames(frame *Frame) int {
	return int(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}
func (ctx *Context) AvBuffersrcClose(next_pts int64, flags uint) int {
	return int(C.av_buffersrc_close((*C.struct_AVFilterContext)(unsafe.Pointer(ctx)), (C.int64_t)(next_pts), C.uint(flags)))
}

//Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *Link, f *Context, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//avfilter_get_class
func AvfilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

//Allocate a single Input entry.
func AvfilterInoutAlloc() *Input {
	return (*Input)(C.avfilter_inout_alloc())
}

//Free the supplied list of Input and set *inout to NULL.
func AvfilterInoutFree(i *Input) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}
func (f *Frame) Pts() int64 {
	return int64(f.pts)
}
func (f *Frame) SetPts(pts int64) {
	f.pts = C.int64_t(pts)
}
func (f *Frame) NextPts() int64 {
	return int64(f.pts + f.pkt_duration)
}
