// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package liblink1

import (
	"bufio"
	"time"
)

var start time.Time

func Cputime() float64 {
	if start.IsZero() {
		start = time.Now()
	}
	return time.Since(start).Seconds()
}

type Biobuf struct {
	unget     int
	haveUnget bool
	r         *bufio.Reader
	w         *bufio.Writer
}

func (b *Biobuf) Write(p []byte) (int, error) {
	return b.w.Write(p)
}

func (b *Biobuf) Flush() error {
	return b.w.Flush()
}

func Bwrite(b *Biobuf, p []byte) (int, error) {
	return b.w.Write(p)
}

func Bputc(b *Biobuf, c byte) {
	b.w.WriteByte(c)
}

func Bgetc(b *Biobuf) int {
	if b.haveUnget {
		b.haveUnget = false
		return int(b.unget)
	}
	c, err := b.r.ReadByte()
	if err != nil {
		b.unget = -1
		return -1
	}
	b.unget = int(c)
	return int(c)
}

func Bungetc(b *Biobuf) {
	b.haveUnget = true
}
