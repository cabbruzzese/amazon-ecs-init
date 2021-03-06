// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Source: backoff.go in package docker
// Automatically generated by MockGen. DO NOT EDIT!

package docker

import (
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// Mock of Backoff interface
type MockBackoff struct {
	ctrl     *gomock.Controller
	recorder *_MockBackoffRecorder
}

// Recorder for MockBackoff (not exported)
type _MockBackoffRecorder struct {
	mock *MockBackoff
}

func NewMockBackoff(ctrl *gomock.Controller) *MockBackoff {
	mock := &MockBackoff{ctrl: ctrl}
	mock.recorder = &_MockBackoffRecorder{mock}
	return mock
}

func (_m *MockBackoff) EXPECT() *_MockBackoffRecorder {
	return _m.recorder
}

func (_m *MockBackoff) Duration() time.Duration {
	ret := _m.ctrl.Call(_m, "Duration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockBackoffRecorder) Duration() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Duration")
}

func (_m *MockBackoff) ShouldRetry() bool {
	ret := _m.ctrl.Call(_m, "ShouldRetry")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockBackoffRecorder) ShouldRetry() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShouldRetry")
}
