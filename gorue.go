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

// Start starts the Gorue read/eval/print loop for a single state identified by the stateIdentity.
func Start(stateIdentity StateIdentity, retriever Retriever, describer Describer, prompter Prompter, interpreter Interpreter, storer Storer) error {
	shouldContinue := true
	for shouldContinue {
		// load the state
		state, err := retriever.Retrieve(stateIdentity)
		if err != nil {
			return err
		}

		// describe the loaded state
		description, err := describer.Describe(state)
		if err != nil {
			return err
		}

		// send the description to the user, accept a response
		response, err := prompter.Prompt(description)
		if err != nil {
			return err
		}

		// interpret the response, modify the state based on that interpretation
		var updatedState SerializedState
		updatedState, shouldContinue, err = interpreter.Interpret(response, state)
		if err != nil {
			return err
		}

		// store the updated state
		err = storer.Store(stateIdentity, updatedState)
		if err != nil {
			return err
		}
	}
	return nil
}
