package avlaoflch

//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <complex_filter.h>
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
