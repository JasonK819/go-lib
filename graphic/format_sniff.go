/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package graphic

import (
	"bufio"
	"os"
)

type format struct {
	name, magic string
}

// image formats:
var formats = []format{
	{"jpeg", "\xff\xd8"},
	{"bmp", "BM????\x00\x00\x00\x00"},
	{"png", "\x89PNG\r\n\x1a\n"},
	{"tiff", "MM\x00\x2A"}, // little-endian
	{"tiff", "II\x2A\x00"}, // big-endian
	{"gif", "GIF8?a"},
}

// Sniff determines the format of r's data.
func sniff(r *bufio.Reader) format {
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

func SniffImageFormat(file string) (string, error) {
	fh, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fh.Close()

	reader := bufio.NewReader(fh)
	format := sniff(reader)
	return format.name, nil
}
