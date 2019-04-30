package logAdapter

import (
	"fmt"
	l "github.com/nuveo/log"
	"log"
	"path/filepath"
	"runtime"
)

func LogAdapter(m l.MsgType, o l.OutType, config map[string]interface{}, msg ...interface{}) {
	if m == l.DebugLog && !l.DebugMode {
		return
	}
	var debugInfo, output string
	if l.DebugMode {
		_, fn, line, _ := runtime.Caller(5)
		fn = filepath.Base(fn)
		debugInfo = fmt.Sprintf("%s:%d ", fn, line)
	}
	if o == l.FormattedOut {
		output = fmt.Sprintf(msg[0].(string), msg[1:]...)
	} else {
		output = fmt.Sprint(msg...)
	}
	if l.EnableANSIColors {
		output = fmt.Sprintf("%s [%s] %s%s\033[0;00m",
			l.Colors[m],
			l.Prefixes[m],
			debugInfo,
			output)
	} else {
		output = fmt.Sprintf("[%s] %s%s",
			l.Prefixes[m],
			debugInfo,
			output)
	}
	if len(output) > l.MaxLineSize {
		output = output[:l.MaxLineSize] + "..."
	}
	log.Println(output)
}
