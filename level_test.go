// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package loghook

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name       string
		level      Level
		reqmessage string
	}{
		{
			name:       "success info level",
			level:      InfoLevel,
			reqmessage: "info",
		},
		{
			name:       "success error level",
			level:      ErrorLevel,
			reqmessage: "error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.reqmessage != tt.level.String() {
				t.Error("level to string err")
			}
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name   string
		reqlvl Level
		reqstr string
	}{
		{
			name:   "success info level",
			reqlvl: InfoLevel,
			reqstr: "info",
		},
		{
			name:   "success error level",
			reqlvl: ErrorLevel,
			reqstr: "error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res, err := ParseLevel(tt.reqstr)
			if err != nil {
				t.Error(err)
			}
			if tt.reqlvl != res {
				t.Errorf("parse level err want %v, got %v", tt.reqlvl, res)
			}
		})
	}
}

func TestMarshalText(t *testing.T) {
	tests := []struct {
		name   string
		reqlvl Level
		reqstr string
	}{
		{
			name:   "success info level",
			reqlvl: InfoLevel,
			reqstr: "info",
		},
		{
			name:   "success error level",
			reqlvl: ErrorLevel,
			reqstr: "error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := tt.reqlvl.MarshalText()
			if tt.reqstr != string(res) {
				t.Errorf("marshal text err want %v, got %v", tt.reqstr, string(res))
			}
		})
	}
}
