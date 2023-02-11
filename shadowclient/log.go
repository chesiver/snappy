package shadowclient

import (
	"bytes"
	"fmt"
	"log"
)

var b = bytes.NewBuffer(make([]byte, 0))
var logger = log.New(b, "", log.Lshortfile|log.LstdFlags)

func logf(f string, v ...interface{}) {
	if config.Verbose {
		logger.Output(2, fmt.Sprintf(f, v...))
	}
}

func DumpLogContent(truncate bool) string {
	res := b.String()
	if truncate {
		b.Truncate(0)
	}
	return res
}
