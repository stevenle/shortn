// Copyright 2016 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

var infolog *log.Logger
var errorlog *log.Logger
var fatallog *log.Logger

func init() {
	infolog = log.New(&LogWriter{
		TimeFormat: time.RFC3339,
		Prefix:     color.GreenString("INFO"),
	}, "" /* prefix */, 0 /* flags */)
	errorlog = log.New(&LogWriter{
		TimeFormat: time.RFC3339,
		Prefix:     color.RedString("ERROR"),
	}, "" /* prefix */, 0 /* flags */)
	fatallog = log.New(&LogWriter{
		TimeFormat: time.RFC3339,
		Prefix:     color.RedString("FATAL"),
	}, "" /* prefix */, 0 /* flags */)
}

type LogWriter struct {
	TimeFormat string
	Prefix     string
}

func (w *LogWriter) Write(bytes []byte) (int, error) {
	now := time.Now().Format(w.TimeFormat)
	return fmt.Print(now + " " + w.Prefix + " " + string(bytes))
}

func Infof(format string, args ...interface{}) {
	infolog.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	errorlog.Printf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	fatallog.Fatalf(format, args...)
}
