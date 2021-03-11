package format

import (
	"github.com/cgoder/vdk/av/avutil"
	"github.com/cgoder/vdk/format/aac"
	"github.com/cgoder/vdk/format/flv"
	"github.com/cgoder/vdk/format/mp4"
	"github.com/cgoder/vdk/format/rtmp"
	"github.com/cgoder/vdk/format/rtsp"
	"github.com/cgoder/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
