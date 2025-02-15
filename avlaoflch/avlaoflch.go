package avlaoflch

//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <complex_filter.h>
//#include <push_stream.h>
//#include <push_stream_cuda.h>
//#include <libavfilter/avfilter.h>
import "C"
import (
	"unsafe"

	"github.com/laoflch/goav/avfilter"
)

type (
	TaskHandleProcessInfo C.struct_TaskHandleProcessInfo
)

func TaskHandleProcessInfoAlloc() *TaskHandleProcessInfo {
	return (*TaskHandleProcessInfo)(unsafe.Pointer(C.task_handle_process_info_alloc()))
}

func TaskHandleProcessInfoFree(info *TaskHandleProcessInfo) {
	//return (*VideoHandleProcessInfo)(unsafe.Pointer(C.video_handle_process_info_alloc()))

	C.task_handle_process_info_free((*C.struct_TaskHandleProcessInfo)(unsafe.Pointer(info)))
}
func PushVideo2RtspSubtitleLogo(VideoFilePath string, VideoIndex int, AudioIndex int, SubtitleFilePath string, LogoFrame **avfilter.Frame, RtspPushPath string, if_hw bool, if_logo_fade bool, duration_frames uint64, interval_frames uint64, present_frames uint64, task_Handle_process_info *TaskHandleProcessInfo) int {
	//return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.CString(SubtitleFilePath), C.CString(LogoFilePath), C.CString(RtspPushPath), C.bool(if_hw)))

	//if SubtitleFilePath == "" {
	//	return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), unsafe.Pointer, (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (**C.struct_VideoHandleProcessInfo)(unsafe.Pointer(video_Handle_process_info))))
	//} else {
	return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), C.CString(SubtitleFilePath), (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (*C.struct_TaskHandleProcessInfo)(unsafe.Pointer(task_Handle_process_info))))

	//}
}
func Push2rtspSubLogo(VideoFilePath string, VideoIndex int, AudioIndex int, SubtitleIndex int, SubtitleFilePath string, LogoFrame **avfilter.Frame, RtspPushPath string, if_hw bool, if_logo_fade bool, duration_frames uint64, interval_frames uint64, present_frames uint64, task_Handle_process_info *TaskHandleProcessInfo) int {
	//return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.CString(SubtitleFilePath), C.CString(LogoFilePath), C.CString(RtspPushPath), C.bool(if_hw)))

	//if SubtitleFilePath == "" {
	//	return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), unsafe.Pointer, (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (**C.struct_VideoHandleProcessInfo)(unsafe.Pointer(video_Handle_process_info))))
	//} else {
	return int(C.push2rtsp_sub_logo(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), C.int(SubtitleIndex), C.CString(SubtitleFilePath), (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (*C.struct_TaskHandleProcessInfo)(unsafe.Pointer(task_Handle_process_info))))

	//}
}
func Push2rtspSubLogoCUDA(VideoFilePath string, VideoIndex int, AudioIndex int, SubtitleIndex int, SubtitleFilePath string, LogoFrame **avfilter.Frame, RtspPushPath string, if_hw bool, if_logo_fade bool, duration_frames uint64, interval_frames uint64, present_frames uint64, task_Handle_process_info *TaskHandleProcessInfo) int {
	//return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.CString(SubtitleFilePath), C.CString(LogoFilePath), C.CString(RtspPushPath), C.bool(if_hw)))

	//if SubtitleFilePath == "" {
	//	return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), unsafe.Pointer, (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (**C.struct_VideoHandleProcessInfo)(unsafe.Pointer(video_Handle_process_info))))
	//} else {
	return int(C.push2rtsp_sub_logo_cuda(C.CString(VideoFilePath), C.int(VideoIndex), C.int(AudioIndex), C.int(SubtitleIndex), C.CString(SubtitleFilePath), (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames), (*C.struct_TaskHandleProcessInfo)(unsafe.Pointer(task_Handle_process_info))))

	//}
}
func (info *TaskHandleProcessInfo) GetPassDuration() int64 {

	return int64(C.int64_t(info.pass_duration))

}

func (info *TaskHandleProcessInfo) GetTotalDuration() int64 {

	return int64(C.uint64_t(info.total_duration))

}

func (info *TaskHandleProcessInfo) GetHandledRate() float32 {

	return float32(C.float(info.handled_rate))

}

func (info *TaskHandleProcessInfo) SetVideoCodecID(d int) {

	//return uint64((*info).total_duration)

	info.video_codec_id = (C.enum_AVCodecID)(d)

}

func (info *TaskHandleProcessInfo) SetVideoSize(w int, h int) {

	//return uint64((*info).total_duration)

	info.width = (C.int)(w)
	info.height = (C.int)(h)

}
func (info *TaskHandleProcessInfo) SeekSeconds(s int64) {

	//return uint64((*info).total_duration)

	info.control.seek_time = (C.int64_t)(s)
	//info.height = (C.int)(h)

}

//设置外挂字幕时间轴的偏移量，单位毫秒(ms)
func (info *TaskHandleProcessInfo) SetSubtitleTimeOffset(offset int64) {

	info.control.subtitle_time_offset = (C.int64_t)(offset)

}

//设置外挂字幕的字符编码格式
func (info *TaskHandleProcessInfo) SetSubtitleCharenc(charenc string) {

	info.control.subtitle_charenc = (C.CString)(charenc)

}

func (info *TaskHandleProcessInfo) TaskCancel() {

	//return uint64((*info).total_duration)

	info.control.task_cancel = (C.bool)(true)
	//info.height = (C.int)(h)

}
func (info *TaskHandleProcessInfo) TaskPause(pause bool) {

	//return uint64((*info).total_duration)

	info.control.task_pause = (C.bool)(pause)
	//info.height = (C.int)(h)

}

func (info *TaskHandleProcessInfo) GetTaskPause() bool {

	return bool(C.bool(info.control.task_pause))

}

func (info *TaskHandleProcessInfo) SetAudioEncodecQ(q_min int, q_max int) {

	info.encodec_para.aq_min = (C.int)(q_min)
	info.encodec_para.aq_max = (C.int)(q_max)

}

func (info *TaskHandleProcessInfo) SetVideoEncodecQ(q_min int, q_max int) {

	info.encodec_para.vq_min = (C.int)(q_min)
	info.encodec_para.vq_max = (C.int)(q_max)

}

func (info *TaskHandleProcessInfo) SetVideoEncodecGopSize(v_gop_size int) {

	info.encodec_para.v_gop_size = (C.int)(v_gop_size)

}

func (info *TaskHandleProcessInfo) SetVideoEncodecMaxBFrames(v_max_b_frames int) {

	info.encodec_para.v_max_b_frames = (C.int)(v_max_b_frames)

}

func (info *TaskHandleProcessInfo) SetVideoEncodecPreset(preset string) {

	info.encodec_para.v_preset = (C.CString)(preset)

}

func (info *TaskHandleProcessInfo) SetRtspTransport(transport string) {

	info.push_rtsp_para.rtsp_transport = (C.CString)(transport)

}

func (info *TaskHandleProcessInfo) SetRtspBufferSize(buffer_size string) {

	info.push_rtsp_para.buffer_size = (C.CString)(buffer_size)

}

func (info *TaskHandleProcessInfo) SetRtspMaxDelay(max_delay string) {

	info.push_rtsp_para.max_delay = (C.CString)(max_delay)

}

func (info *TaskHandleProcessInfo) SetRtspTimeout(timeout string) {

	info.push_rtsp_para.timeout = (C.CString)(timeout)

}

func (info *TaskHandleProcessInfo) SetRtspSTimeout(stimeout string) {

	info.push_rtsp_para.stimeout = (C.CString)(stimeout)

}
