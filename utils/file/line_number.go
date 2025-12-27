package file

import (
	"bytes"
	"fmt"
	"slices"
)

func AttachLineNumbers(content []byte) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(content, []byte("\n"))
	var buf bytes.Buffer
	for i, line := range lines {
		fmt.Fprintf(&buf, "%4d: %s\n", i+1, line)
	}
	return buf.Bytes()
}

type InsertLineAfterOptions struct {
	Number int
	Line   string
}

func InsertLineAfter(content []byte, opts InsertLineAfterOptions) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(content, []byte("\n"))
	var buf bytes.Buffer
	for i, line := range lines {
		if i+1 == opts.Number {
			fmt.Fprintf(&buf, "%s\n%s\n", line, opts.Line)
		} else {
			fmt.Fprintf(&buf, "%s\n", line)
		}
	}
	return buf.Bytes()
}

type RemoveLineOptions struct {
	Number int
}

func RemoveLine(content []byte, opts RemoveLineOptions) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(content, []byte("\n"))
	var buf bytes.Buffer
	for i, line := range lines {
		if i+1 == opts.Number {
			// Skip this line
			continue
		} else {
			fmt.Fprintf(&buf, "%s\n", line)
		}
	}
	return buf.Bytes()
}

type ReplaceLineOptions struct {
	Number int
	Line   string
}

func ReplaceLine(content []byte, opts ReplaceLineOptions) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(content, []byte("\n"))
	var buf bytes.Buffer
	for i, line := range lines {
		if i+1 == opts.Number {
			fmt.Fprintf(&buf, "%s\n", opts.Line)
		} else {
			fmt.Fprintf(&buf, "%s\n", line)
		}
	}
	return buf.Bytes()
}

func ReplaceLines(content []byte, opts ...ReplaceLineOptions) []byte {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(content, []byte("\n"))
	var buf bytes.Buffer
	for i, line := range lines {
		index := slices.IndexFunc(opts, func(o ReplaceLineOptions) bool {
			return o.Number == i+1
		})
		if index != -1 {
			fmt.Fprintf(&buf, "%s\n", opts[index].Line)
		} else {
			fmt.Fprintf(&buf, "%s\n", line)
		}
	}
	return buf.Bytes()
}
