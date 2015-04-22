//stub for testing implements io.Writer
//string value accessible via Val() method

package stringwriter

import (
	"bytes"
	"strings"
)

type StringW struct {
	bytes.Buffer
}

//turns newline-delimited []byte data into []string
func (sw *StringW) Val() (value []string) {
	str := strings.TrimSpace(sw.String())
	value = strings.Split(str, "\n")
	return
}
