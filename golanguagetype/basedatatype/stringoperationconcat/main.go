package main

// Example of string concat operation.

import (
	"bytes"
	"strings"
)

func plusConcat(n int, str string) string {
	// +号拼接
}

func sprintfConcat(n int, str string) string {
	// fmt.Sprintf拼接
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func builderGrowConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	// 与builderConcat相同

}

func bufferGrowConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	buf.Grow(n * len(s))
	// 与bufferConcat相同
}
