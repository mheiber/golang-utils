package stringwriter_test

import (
	"fmt"
	"github.com/mheiber/golang-utils/stringwriter"
	"reflect"
	"testing"
)

func TestStringw(t *testing.T) {
	sw := new(stringwriter.StringW)

	fmt.Fprint(sw, "  hello\nstringW\n\n  ")

	expected := []string{"hello", "stringW"}
	if !reflect.DeepEqual(sw.Val(), expected) {
		t.Errorf("Expected %v to equal %v", sw.Val(), expected)
	}
}
