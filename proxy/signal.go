// Copyright 2016 ByteDance, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	. "github.com/bytedance/dbatman/log"
	"os"
)

type SignalHandler func(s os.Signal, arg interface{}) error

type SignalSet struct {
	M map[os.Signal]SignalHandler
}

func NewSignalSet() *SignalSet {
	s := new(SignalSet)
	s.M = make(map[os.Signal]SignalHandler)
	return s
}

func (s *SignalSet) Register(sig os.Signal, handler SignalHandler) {
	if _, exist := s.M[sig]; !exist {
		s.M[sig] = handler
	}
}

func (s *SignalSet) Handle(sig os.Signal, arg interface{}) error {
	if handler, exist := s.M[sig]; exist {
		return handler(sig, arg)
	} else {
		SysLog.Warn("no available handler for signal %v, ignore!", sig)
		return nil
	}
}

func init() {

}

/* vim: set expandtab ts=4 sw=4 */
