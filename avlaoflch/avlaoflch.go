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

func PushVideo2RtspSubtitleLogo(VideoFilePath string, SubtitleFilePath string, LogoFrame **avfilter.Frame, RtspPushPath string, if_hw bool, if_logo_fade bool, duration_frames uint64, interval_frames uint64, present_frames uint64) int {
	//return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.CString(SubtitleFilePath), C.CString(LogoFilePath), C.CString(RtspPushPath), C.bool(if_hw)))
	return int(C.push_video_to_rtsp_subtitle_logo(C.CString(VideoFilePath), C.CString(SubtitleFilePath), (**C.struct_AVFrame)(unsafe.Pointer(LogoFrame)), C.CString(RtspPushPath), C.bool(if_hw), C.bool(if_logo_fade), C.uint64_t(duration_frames), C.uint64_t(interval_frames), C.uint64_t(present_frames)))

}
