package liblog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

const (
	LEVEL_ERROR = iota
	LEVEL_DONE
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_CACHE
	LEVEL_DEBUG
)

var Output io.Writer = os.Stdout
var LogLevel = LEVEL_CACHE
var mutex sync.Mutex

func Debug(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_DEBUG {
		fmt.Fprintln(
			Output,
			libescapes.Optional(libescapes.TrueColor(128, 128, 128))+
				arrow(nesting)+
				fmt.Sprintf(format, args...)+
				libescapes.Optional(libescapes.ColorReset),
		)
	}
}

func Cached(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_CACHE {
		log(
			nesting,
			"CACHED: ",
			libescapes.Optional(libescapes.TextColorBrightMagenta),
			fmt.Sprintf(format, args...),
		)
	}
}

func Info(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_INFO {
		log(
			nesting,
			"",
			libescapes.Optional(libescapes.TextColorBrightBlue),
			fmt.Sprintf(format, args...),
		)
	}
}

func Warn(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"WARN: ",
			libescapes.Optional(libescapes.TextColorBrightYellow),
			fmt.Sprintf(format, args...),
		)
	}
}

func Done(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"DONE: ",
			libescapes.Optional(libescapes.TextColorBrightGreen),
			fmt.Sprintf(format, args...),
		)
	}
}

func Error(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"ERROR: ",
			libescapes.Optional(libescapes.TextColorBrightRed),
			fmt.Sprintf(format, args...),
		)
	}
}

func log(nesting uint, prefix, color, body string) {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Fprintln(Output, color+
		arrow(nesting)+
		prefix+
		libescapes.Optional(libescapes.ColorReset)+
		body)
}

func arrow(nesting uint) string {
	switch nesting {
	case 0:
		return "==> "
	case 1:
		return " -> "
	}

	return strings.Repeat("  ", int(nesting)) + "-> "
}
