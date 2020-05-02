// Copyright 2020 Kyle Stafford

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gorue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockGorue struct {
	mock.Mock
	state SerializedState
}

func (m *mockGorue) Retrieve(identity StateIdentity) (SerializedState, error) {
	args := m.Called(identity)
	return args.Get(0).(SerializedState), args.Error(1)
}

func (m *mockGorue) Describe(state SerializedState) (string, error) {
	args := m.Called(state)
	return args.String(0), args.Error(1)
}

func (m *mockGorue) Prompt(out string) (string, error) {
	args := m.Called(out)
	return args.String(0), args.Error(1)
}

func (m *mockGorue) Interpret(command string, state SerializedState) (SerializedState, bool, error) {
	args := m.Called(command, state)
	return args.Get(0).(SerializedState), args.Bool(1), args.Error(2)
}

func (m *mockGorue) Store(identity StateIdentity, state SerializedState) error {
	args := m.Called(identity, state)
	return args.Error(0)
}

func TestGorueOneTurn(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var description = "description"
	var response = "response"
	var updatedState = SerializedState("updated state")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(description, nil).Once()
	mock.On("Prompt", description).Return(response, nil).Once()
	mock.On("Interpret", response, initialState).Return(updatedState, false, nil).Once()
	mock.On("Store", identity, updatedState).Return(nil).Once()

	Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
}

func TestGorueTwoTurns(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var firstDescription = "first description"
	var firstResponse = "first response"
	var secondState = SerializedState("second state")
	var secondDescription = "second description"
	var secondResponse = "second response"
	var thirdState = SerializedState("third state")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(firstDescription, nil).Once()
	mock.On("Prompt", firstDescription).Return(firstResponse, nil).Once()
	mock.On("Interpret", firstResponse, initialState).Return(secondState, true, nil).Once()
	mock.On("Store", identity, secondState).Return(nil).Once()

	mock.On("Retrieve", identity).Return(secondState, nil).Once()
	mock.On("Describe", secondState).Return(secondDescription, nil).Once()
	mock.On("Prompt", secondDescription).Return(secondResponse, nil).Once()
	mock.On("Interpret", secondResponse, secondState).Return(thirdState, false, nil).Once()
	mock.On("Store", identity, thirdState).Return(nil).Once()

	Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
}

func TestGorueRetrieveError(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var expectedError = errors.New("womp womp")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, expectedError).Once()

	var err = Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
	if err != expectedError {
		t.Error("Retrieve error was not returned")
	}
}

func TestGorueDescribeError(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var description = "description"
	var expectedError = errors.New("womp womp")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(description, expectedError).Once()

	var err = Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
	if err != expectedError {
		t.Error("Describe error was not returned")
	}
}

func TestGoruePromptError(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var description = "description"
	var response = "response"
	var expectedError = errors.New("womp womp")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(description, nil).Once()
	mock.On("Prompt", description).Return(response, expectedError).Once()

	var err = Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
	if err != expectedError {
		t.Error("Prompt error was not returned")
	}
}

func TestGorueInterpretError(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var description = "description"
	var response = "response"
	var updatedState = SerializedState("updated state")
	var expectedError = errors.New("womp womp")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(description, nil).Once()
	mock.On("Prompt", description).Return(response, nil).Once()
	mock.On("Interpret", response, initialState).Return(updatedState, false, expectedError).Once()

	var err = Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
	if err != expectedError {
		t.Error("Interpret error was not returned")
	}
}

func TestGorueStoreError(t *testing.T) {
	var identity = StateIdentity("test state")
	var initialState = SerializedState("state")
	var description = "description"
	var response = "response"
	var updatedState = SerializedState("updated state")
	var expectedError = errors.New("womp womp")
	mock := new(mockGorue)

	mock.On("Retrieve", identity).Return(initialState, nil).Once()
	mock.On("Describe", initialState).Return(description, nil).Once()
	mock.On("Prompt", description).Return(response, nil).Once()
	mock.On("Interpret", response, initialState).Return(updatedState, false, nil).Once()
	mock.On("Store", identity, updatedState).Return(expectedError).Once()

	var err = Start(identity, mock, mock, mock, mock, mock)

	mock.AssertExpectations(t)
	if err != expectedError {
		t.Error("Store error was not returned")
	}
}
