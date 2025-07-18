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

// DeviceMetricCalibrationState is documented here http://hl7.org/fhir/ValueSet/metric-calibration-state
type DeviceMetricCalibrationState int

const (
	DeviceMetricCalibrationStateNotCalibrated DeviceMetricCalibrationState = iota
	DeviceMetricCalibrationStateCalibrationRequired
	DeviceMetricCalibrationStateCalibrated
	DeviceMetricCalibrationStateUnspecified
)

func (code DeviceMetricCalibrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricCalibrationState) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-calibrated":
		*code = DeviceMetricCalibrationStateNotCalibrated
	case "calibration-required":
		*code = DeviceMetricCalibrationStateCalibrationRequired
	case "calibrated":
		*code = DeviceMetricCalibrationStateCalibrated
	case "unspecified":
		*code = DeviceMetricCalibrationStateUnspecified
	default:
		return fmt.Errorf("unknown DeviceMetricCalibrationState code `%s`", s)
	}
	return nil
}
func (code DeviceMetricCalibrationState) String() string {
	return code.Code()
}
func (code DeviceMetricCalibrationState) Code() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "not-calibrated"
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "calibration-required"
	case DeviceMetricCalibrationStateCalibrated:
		return "calibrated"
	case DeviceMetricCalibrationStateUnspecified:
		return "unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationState) Display() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "Not Calibrated"
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "Calibration Required"
	case DeviceMetricCalibrationStateCalibrated:
		return "Calibrated"
	case DeviceMetricCalibrationStateUnspecified:
		return "Unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationState) Definition() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "The metric has not been calibrated."
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "The metric needs to be calibrated."
	case DeviceMetricCalibrationStateCalibrated:
		return "The metric has been calibrated."
	case DeviceMetricCalibrationStateUnspecified:
		return "The state of calibration of this metric is unspecified."
	}
	return "<unknown>"
}
