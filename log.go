package liblog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
	"golang.org/x/term"
)

const (
	LEVEL_ERROR = iota
	LEVEL_DONE
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
)

var Output io.Writer = os.Stdout
var IsTerminal = true
var UseColors = true
var LogLevel = LEVEL_INFO
var mutex sync.Mutex

func init() {
	IsTerminal = term.IsTerminal(int(os.Stdout.Fd()))
	UseColors = IsTerminal && os.Getenv("NO_COLOR") != "1"
}

func Debug(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_DEBUG {
		fmt.Fprintln(
			Output,
			optional(libescapes.TrueColor(128, 128, 128))+
				arrow(nesting)+
				fmt.Sprintf(format, args...)+
				optional(libescapes.ColorReset),
		)
	}
}

func Info(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_INFO {
		log(
			nesting,
			"",
			optional(libescapes.TextColorBrightBlue),
			fmt.Sprintf(format, args...),
		)
	}
}

func Warn(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"WARN: ",
			optional(libescapes.TextColorBrightYellow),
			fmt.Sprintf(format, args...),
		)
	}
}

func Done(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"DONE: ",
			optional(libescapes.TextColorBrightGreen),
			fmt.Sprintf(format, args...),
		)
	}
}

func Error(nesting uint, format string, args ...any) {
	if LogLevel >= LEVEL_WARN {
		log(
			nesting,
			"ERROR: ",
			optional(libescapes.TextColorBrightRed),
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
		optional(libescapes.ColorReset)+
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

func optional(ansi string) string {
	if UseColors {
		return ansi
	}
	return ""
}
