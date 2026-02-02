package liblog_test

import (
	"testing"

	liblog "github.com/bbfh-dev/lib-log"
)

func TestDemo(t *testing.T) {
	liblog.Output = t.Output()
	liblog.LogLevel = liblog.LEVEL_DEBUG
	liblog.UseColors = true

	liblog.Debug(0, "This is a message")
	liblog.Cached(0, "This is a message")
	liblog.Info(0, "This is a message")
	liblog.Warn(0, "This is a message")
	liblog.Done(0, "This is a message")
	liblog.Error(0, "This is a message")
}
