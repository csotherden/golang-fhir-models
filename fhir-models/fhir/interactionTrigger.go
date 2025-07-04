// Copyright 2019 - 2025 The Samply Community
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

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// InteractionTrigger is documented here http://hl7.org/fhir/ValueSet/interaction-trigger
type InteractionTrigger int

const (
	InteractionTriggerRead InteractionTrigger = iota
	InteractionTriggerVread
	InteractionTriggerUpdate
	InteractionTriggerPatch
	InteractionTriggerDelete
	InteractionTriggerHistory
	InteractionTriggerHistoryInstance
	InteractionTriggerHistoryType
	InteractionTriggerHistorySystem
	InteractionTriggerCreate
	InteractionTriggerSearch
	InteractionTriggerSearchType
	InteractionTriggerSearchSystem
	InteractionTriggerSearchCompartment
	InteractionTriggerCapabilities
	InteractionTriggerTransaction
	InteractionTriggerBatch
	InteractionTriggerOperation
)

func (code InteractionTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *InteractionTrigger) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "read":
		*code = InteractionTriggerRead
	case "vread":
		*code = InteractionTriggerVread
	case "update":
		*code = InteractionTriggerUpdate
	case "patch":
		*code = InteractionTriggerPatch
	case "delete":
		*code = InteractionTriggerDelete
	case "history":
		*code = InteractionTriggerHistory
	case "history-instance":
		*code = InteractionTriggerHistoryInstance
	case "history-type":
		*code = InteractionTriggerHistoryType
	case "history-system":
		*code = InteractionTriggerHistorySystem
	case "create":
		*code = InteractionTriggerCreate
	case "search":
		*code = InteractionTriggerSearch
	case "search-type":
		*code = InteractionTriggerSearchType
	case "search-system":
		*code = InteractionTriggerSearchSystem
	case "search-compartment":
		*code = InteractionTriggerSearchCompartment
	case "capabilities":
		*code = InteractionTriggerCapabilities
	case "transaction":
		*code = InteractionTriggerTransaction
	case "batch":
		*code = InteractionTriggerBatch
	case "operation":
		*code = InteractionTriggerOperation
	default:
		return fmt.Errorf("unknown InteractionTrigger code `%s`", s)
	}
	return nil
}
func (code InteractionTrigger) String() string {
	return code.Code()
}
func (code InteractionTrigger) Code() string {
	switch code {
	case InteractionTriggerRead:
		return "read"
	case InteractionTriggerVread:
		return "vread"
	case InteractionTriggerUpdate:
		return "update"
	case InteractionTriggerPatch:
		return "patch"
	case InteractionTriggerDelete:
		return "delete"
	case InteractionTriggerHistory:
		return "history"
	case InteractionTriggerHistoryInstance:
		return "history-instance"
	case InteractionTriggerHistoryType:
		return "history-type"
	case InteractionTriggerHistorySystem:
		return "history-system"
	case InteractionTriggerCreate:
		return "create"
	case InteractionTriggerSearch:
		return "search"
	case InteractionTriggerSearchType:
		return "search-type"
	case InteractionTriggerSearchSystem:
		return "search-system"
	case InteractionTriggerSearchCompartment:
		return "search-compartment"
	case InteractionTriggerCapabilities:
		return "capabilities"
	case InteractionTriggerTransaction:
		return "transaction"
	case InteractionTriggerBatch:
		return "batch"
	case InteractionTriggerOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code InteractionTrigger) Display() string {
	switch code {
	case InteractionTriggerRead:
		return "read"
	case InteractionTriggerVread:
		return "vread"
	case InteractionTriggerUpdate:
		return "update"
	case InteractionTriggerPatch:
		return "patch"
	case InteractionTriggerDelete:
		return "delete"
	case InteractionTriggerHistory:
		return "history"
	case InteractionTriggerHistoryInstance:
		return "history-instance"
	case InteractionTriggerHistoryType:
		return "history-type"
	case InteractionTriggerHistorySystem:
		return "history-system"
	case InteractionTriggerCreate:
		return "create"
	case InteractionTriggerSearch:
		return "search"
	case InteractionTriggerSearchType:
		return "search-type"
	case InteractionTriggerSearchSystem:
		return "search-system"
	case InteractionTriggerSearchCompartment:
		return "search-compartment"
	case InteractionTriggerCapabilities:
		return "capabilities"
	case InteractionTriggerTransaction:
		return "transaction"
	case InteractionTriggerBatch:
		return "batch"
	case InteractionTriggerOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code InteractionTrigger) Definition() string {
	switch code {
	case InteractionTriggerRead:
		return "Read the current state of the resource."
	case InteractionTriggerVread:
		return "Read the state of a specific version of the resource."
	case InteractionTriggerUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case InteractionTriggerPatch:
		return "Update an existing resource by posting a set of changes to it."
	case InteractionTriggerDelete:
		return "Delete a resource."
	case InteractionTriggerHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case InteractionTriggerHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case InteractionTriggerHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case InteractionTriggerHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case InteractionTriggerCreate:
		return "Create a new resource with a server assigned id."
	case InteractionTriggerSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case InteractionTriggerSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case InteractionTriggerSearchSystem:
		return "Search all resources based on some filter criteria."
	case InteractionTriggerSearchCompartment:
		return "Search resources in a compartment based on some filter criteria."
	case InteractionTriggerCapabilities:
		return "Get a Capability Statement for the system."
	case InteractionTriggerTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case InteractionTriggerBatch:
		return "perform a set of a separate interactions in a single http operation"
	case InteractionTriggerOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
