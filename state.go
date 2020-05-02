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

// StateIdentity is an address for a single state
// Each gorue loop focuses on a single state, identified by a unique StateIdentity
type StateIdentity []byte

// SerializedState is a persistable representation of a gorue loop context.
// State informs the gorue components how to prompt the user and interpret their response.
type SerializedState []byte

// Storer stores a SerializedState identified by the provided identity
type Storer interface {
	Store(identity StateIdentity, state SerializedState) error
}

// Retriever is responsible for retrieving the context for a single iteration of the loop.
// identity identifies the state to be retrieved
type Retriever interface {
	Retrieve(identity StateIdentity) (SerializedState, error)
}

// StoreRetriever is the interface that groups the basic Store and Retrieve methods
type StoreRetriever interface {
	Storer
	Retriever
}
