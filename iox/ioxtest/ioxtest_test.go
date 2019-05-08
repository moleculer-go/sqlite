// Copyright (c) 2018 David Crawshaw <david@zentus.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package ioxtest

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestTester(t *testing.T) {
	f1, err := ioutil.TempFile("", "iotest")
	if err != nil {
		t.Fatal(err)
	}
	f2, err := ioutil.TempFile("", "iotest")
	if err != nil {
		t.Fatal(err)
	}
	ft := &Tester{T: t, F1: f1, F2: f2}
	ft.Run()
}

func TestBuffer(t *testing.T) {
	ft := &Tester{
		T:  t,
		F1: new(bytes.Buffer),
		F2: new(bytes.Buffer),
	}
	ft.Run()
}
