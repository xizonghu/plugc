package logutils

import (
	"bytes"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"

	//logger "fmt"
	"io"
	//logger "log"
	"path/filepath"
	"strings"
	"time"

	//logger "github.com/sirupsen/logrus"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *logger.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("[%s %s:%d ][routine-%d][%s] %s\n", timestamp, file, len, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GetConsoleWriter() io.Writer {
	return os.Stdout
}

func init() {
	//logger.SetFlags(logger.LstdFlags | logger.Lshortfile)
	//logger.SetFlags(logger.Lshortfile)
}