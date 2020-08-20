// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package log

// DefaultNopLogger returns a nop logger.
func DefaultNopLogger() Logger {
	return nopLogger{}
}

type nopLogger struct{}

func (l nopLogger) WithFields(fields Fields) Logger                 { return l }
func (l nopLogger) WithField(name string, value interface{}) Logger { return l }
func (l nopLogger) WithError(err error) Logger                      { return l }
func (l nopLogger) Debug(args ...interface{})                       {}
func (l nopLogger) Info(args ...interface{})                        {}
func (l nopLogger) Warn(args ...interface{})                        {}
func (l nopLogger) Error(args ...interface{})                       {}
func (l nopLogger) Debugf(format string, args ...interface{})       {}
func (l nopLogger) Infof(format string, args ...interface{})        {}
func (l nopLogger) Warnf(format string, args ...interface{})        {}
func (l nopLogger) Errorf(format string, args ...interface{})       {}
