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

package errors

import (
	juju "github.com/juju/errors"
)

var trace = false

func SetTrace(t bool) {
	trace = t
}

func Trace(err error) error {
	if !trace {
		return err
	}

	if err == nil {
		return nil
	}

	e := juju.Trace(err).(*juju.Err)
	e.SetLocation(1)
	return e
}

func Real(err error) error {
	if !trace {
		return err
	}

	if err == nil {
		return nil
	}

	if e, ok := err.(*juju.Err); ok {
		return Real(e.Underlying())
	}

	return err
}

func ErrorStack(err error) string {
	return juju.ErrorStack(err)
}
