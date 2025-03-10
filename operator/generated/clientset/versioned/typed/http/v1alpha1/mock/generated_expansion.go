// /*
// Copyright 2023 The KEDA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: operator/generated/clientset/versioned/typed/http/v1alpha1/generated_expansion.go
//
// Generated by this command:
//
//	mockgen -copyright_file=hack/boilerplate.go.txt -destination=operator/generated/clientset/versioned/typed/http/v1alpha1/mock/generated_expansion.go -package=mock -source=operator/generated/clientset/versioned/typed/http/v1alpha1/generated_expansion.go
//

// Package mock is a generated GoMock package.
package mock

import (
	gomock "go.uber.org/mock/gomock"
)

// MockHTTPScaledObjectExpansion is a mock of HTTPScaledObjectExpansion interface.
type MockHTTPScaledObjectExpansion struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScaledObjectExpansionMockRecorder
	isgomock struct{}
}

// MockHTTPScaledObjectExpansionMockRecorder is the mock recorder for MockHTTPScaledObjectExpansion.
type MockHTTPScaledObjectExpansionMockRecorder struct {
	mock *MockHTTPScaledObjectExpansion
}

// NewMockHTTPScaledObjectExpansion creates a new mock instance.
func NewMockHTTPScaledObjectExpansion(ctrl *gomock.Controller) *MockHTTPScaledObjectExpansion {
	mock := &MockHTTPScaledObjectExpansion{ctrl: ctrl}
	mock.recorder = &MockHTTPScaledObjectExpansionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScaledObjectExpansion) EXPECT() *MockHTTPScaledObjectExpansionMockRecorder {
	return m.recorder
}
