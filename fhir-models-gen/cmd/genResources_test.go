package cmd

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/csotherden/golang-fhir-models/fhir-models-gen/fhir"
	"github.com/dave/jennifer/jen"
)

func TestUnmarshalResource(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expected      Resource
		expectError   bool
		errorContains string
	}{
		{
			name: "valid resource with all fields",
			input: []byte(`{
				"resourceType": "StructureDefinition",
				"url": "http://hl7.org/fhir/StructureDefinition/Patient",
				"version": "4.0.1",
				"name": "Patient"
			}`),
			expected: Resource{
				ResourceType: "StructureDefinition",
				Url:          stringPtr("http://hl7.org/fhir/StructureDefinition/Patient"),
				Version:      stringPtr("4.0.1"),
				Name:         stringPtr("Patient"),
			},
			expectError: false,
		},
		{
			name: "resource with minimal fields",
			input: []byte(`{
				"resourceType": "ValueSet"
			}`),
			expected: Resource{
				ResourceType: "ValueSet",
				Url:          nil,
				Version:      nil,
				Name:         nil,
			},
			expectError: false,
		},
		{
			name: "resource with partial fields",
			input: []byte(`{
				"resourceType": "CodeSystem",
				"url": "http://terminology.hl7.org/CodeSystem/v3-ActMood",
				"name": "ActMood"
			}`),
			expected: Resource{
				ResourceType: "CodeSystem",
				Url:          stringPtr("http://terminology.hl7.org/CodeSystem/v3-ActMood"),
				Version:      nil,
				Name:         stringPtr("ActMood"),
			},
			expectError: false,
		},
		{
			name:          "invalid JSON",
			input:         []byte(`{"resourceType": "StructureDefinition", "invalid": json}`),
			expected:      Resource{},
			expectError:   true,
			errorContains: "invalid character",
		},
		{
			name:        "empty JSON",
			input:       []byte(`{}`),
			expected:    Resource{ResourceType: ""},
			expectError: false,
		},
		{
			name:        "null JSON",
			input:       []byte(`null`),
			expected:    Resource{},
			expectError: false,
		},
		{
			name:          "empty input",
			input:         []byte(``),
			expected:      Resource{},
			expectError:   true,
			errorContains: "unexpected end of JSON input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnmarshalResource(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain '%s', got '%s'", tt.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result.ResourceType != tt.expected.ResourceType {
				t.Errorf("expected ResourceType '%s', got '%s'", tt.expected.ResourceType, result.ResourceType)
			}

			if !stringPtrEqual(result.Url, tt.expected.Url) {
				t.Errorf("expected Url %v, got %v", tt.expected.Url, result.Url)
			}

			if !stringPtrEqual(result.Version, tt.expected.Version) {
				t.Errorf("expected Version %v, got %v", tt.expected.Version, result.Version)
			}

			if !stringPtrEqual(result.Name, tt.expected.Name) {
				t.Errorf("expected Name %v, got %v", tt.expected.Name, result.Name)
			}
		})
	}
}

func TestFirstLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "uppercase first letter",
			input:    "Patient",
			expected: "patient",
		},
		{
			name:     "already lowercase first letter",
			input:    "patient",
			expected: "patient",
		},
		{
			name:     "single uppercase letter",
			input:    "P",
			expected: "p",
		},
		{
			name:     "single lowercase letter",
			input:    "p",
			expected: "p",
		},
		{
			name:     "mixed case",
			input:    "PatientRecord",
			expected: "patientRecord",
		},
		{
			name:     "all uppercase",
			input:    "PATIENT",
			expected: "pATIENT",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "number first",
			input:    "1Patient",
			expected: "1Patient",
		},
		{
			name:     "special character first",
			input:    "_Patient",
			expected: "_Patient",
		},
		{
			name:     "multibyte character",
			input:    "😀Patient",
			expected: "😀Patient",
		},
		{
			name:     "unicode character",
			input:    "Ñame",
			expected: "ñame",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstLower(tt.input)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestTypeCodeToTypeIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic types
		{
			name:     "base64Binary",
			input:    "base64Binary",
			expected: "string",
		},
		{
			name:     "boolean",
			input:    "boolean",
			expected: "bool",
		},
		{
			name:     "canonical",
			input:    "canonical",
			expected: "string",
		},
		{
			name:     "code",
			input:    "code",
			expected: "string",
		},
		{
			name:     "date",
			input:    "date",
			expected: "string",
		},
		{
			name:     "dateTime",
			input:    "dateTime",
			expected: "string",
		},
		{
			name:     "id",
			input:    "id",
			expected: "string",
		},
		{
			name:     "instant",
			input:    "instant",
			expected: "string",
		},
		{
			name:     "integer",
			input:    "integer",
			expected: "int",
		},
		{
			name:     "markdown",
			input:    "markdown",
			expected: "string",
		},
		{
			name:     "oid",
			input:    "oid",
			expected: "string",
		},
		{
			name:     "positiveInt",
			input:    "positiveInt",
			expected: "int",
		},
		{
			name:     "string",
			input:    "string",
			expected: "string",
		},
		{
			name:     "time",
			input:    "time",
			expected: "string",
		},
		{
			name:     "unsignedInt",
			input:    "unsignedInt",
			expected: "int",
		},
		{
			name:     "uri",
			input:    "uri",
			expected: "string",
		},
		{
			name:     "url",
			input:    "url",
			expected: "string",
		},
		{
			name:     "uuid",
			input:    "uuid",
			expected: "string",
		},
		{
			name:     "xhtml",
			input:    "xhtml",
			expected: "string",
		},
		{
			name:     "fhirpath string",
			input:    "http://hl7.org/fhirpath/System.String",
			expected: "string",
		},
		// Unknown/complex types should return as-is
		{
			name:     "unknown type",
			input:    "CustomType",
			expected: "CustomType",
		},
		{
			name:     "complex type",
			input:    "Address",
			expected: "Address",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := typeCodeToTypeIdentifier(tt.input)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestRequiredValueSetBinding(t *testing.T) {
	tests := []struct {
		name     string
		element  fhir.ElementDefinition
		expected *string
	}{
		{
			name: "required binding with valueSet",
			element: fhir.ElementDefinition{
				Binding: &fhir.ElementDefinitionBinding{
					Strength: fhir.BindingStrengthRequired,
					ValueSet: stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
				},
			},
			expected: stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
		},
		{
			name: "extensible binding should return nil",
			element: fhir.ElementDefinition{
				Binding: &fhir.ElementDefinitionBinding{
					Strength: fhir.BindingStrengthExtensible,
					ValueSet: stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
				},
			},
			expected: nil,
		},
		{
			name: "preferred binding should return nil",
			element: fhir.ElementDefinition{
				Binding: &fhir.ElementDefinitionBinding{
					Strength: fhir.BindingStrengthPreferred,
					ValueSet: stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
				},
			},
			expected: nil,
		},
		{
			name: "example binding should return nil",
			element: fhir.ElementDefinition{
				Binding: &fhir.ElementDefinitionBinding{
					Strength: fhir.BindingStrengthExample,
					ValueSet: stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
				},
			},
			expected: nil,
		},
		{
			name: "required binding without valueSet should return nil",
			element: fhir.ElementDefinition{
				Binding: &fhir.ElementDefinitionBinding{
					Strength: fhir.BindingStrengthRequired,
					ValueSet: nil,
				},
			},
			expected: nil,
		},
		{
			name:     "no binding should return nil",
			element:  fhir.ElementDefinition{Binding: nil},
			expected: nil,
		},
		{
			name:     "empty element should return nil",
			element:  fhir.ElementDefinition{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := requiredValueSetBinding(tt.element)
			if !stringPtrEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAppendLicenseComment(t *testing.T) {
	tests := []struct {
		name     string
		validate func(t *testing.T, code string)
	}{
		{
			name: "license comment is added",
			validate: func(t *testing.T, code string) {
				// Check that license text is present
				expectedLicenseContent := []string{
					"Copyright 2019 - 2025 The Samply Community",
					"Licensed under the Apache License, Version 2.0",
					"http://www.apache.org/licenses/LICENSE-2.0",
					"WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND",
				}

				for _, content := range expectedLicenseContent {
					if !strings.Contains(code, content) {
						t.Errorf("expected license content '%s' not found in generated code", content)
					}
				}

				// Check that comments are properly formatted (start with //)
				lines := strings.Split(code, "\n")
				licenseStarted := false
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if strings.Contains(line, "Copyright 2019 - 2025") {
						licenseStarted = true
					}
					if licenseStarted && line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "package") {
						if strings.Contains(line, "Apache License") || strings.Contains(line, "WITHOUT WARRANTIES") {
							t.Errorf("license line '%s' is not properly commented", line)
						}
					}
					if strings.HasPrefix(line, "package") {
						break
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := jen.NewFile("test")
			appendLicenseComment(file)

			code := file.GoString()
			if code == "" {
				t.Errorf("generated code is empty")
				return
			}

			tt.validate(t, code)
		})
	}
}

func TestAppendGeneratorComment(t *testing.T) {
	tests := []struct {
		name     string
		validate func(t *testing.T, code string)
	}{
		{
			name: "generator comment is added",
			validate: func(t *testing.T, code string) {
				expectedContent := []string{
					"THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models",
					"PLEASE DO NOT EDIT BY HAND",
				}

				for _, content := range expectedContent {
					if !strings.Contains(code, content) {
						t.Errorf("expected generator content '%s' not found in generated code", content)
					}
				}

				// Check that the generator comment is properly formatted
				if !strings.Contains(code, "// THIS FILE IS GENERATED BY") {
					t.Errorf("generator comment is not properly formatted as a comment")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := jen.NewFile("test")
			appendGeneratorComment(file)

			code := file.GoString()
			if code == "" {
				t.Errorf("generated code is empty")
				return
			}

			tt.validate(t, code)
		})
	}
}

// Test combining license and generator comments
func TestCombinedComments(t *testing.T) {
	file := jen.NewFile("test")
	appendLicenseComment(file)
	appendGeneratorComment(file)

	code := file.GoString()

	// Verify both comments are present
	if !strings.Contains(code, "Copyright 2019 - 2025") {
		t.Errorf("license comment not found")
	}
	if !strings.Contains(code, "THIS FILE IS GENERATED BY") {
		t.Errorf("generator comment not found")
	}

	// Verify proper ordering (license should come first in header, generator in body)
	lines := strings.Split(code, "\n")
	licenseIndex := -1
	generatorIndex := -1

	for i, line := range lines {
		if strings.Contains(line, "Copyright 2019 - 2025") {
			licenseIndex = i
		}
		if strings.Contains(line, "THIS FILE IS GENERATED BY") {
			generatorIndex = i
		}
	}

	if licenseIndex == -1 {
		t.Errorf("license comment not found in lines")
	}
	if generatorIndex == -1 {
		t.Errorf("generator comment not found in lines")
	}
}

// Edge case tests
func TestEdgeCases(t *testing.T) {
	t.Run("FirstLower with multibyte character - should now work correctly", func(t *testing.T) {
		// These tests verify that Unicode characters now work properly with rune handling
		testCases := []struct {
			input    string
			expected string
		}{
			{"😀Test", "😀Test"}, // Emoji shouldn't change (not uppercase)
			{"αTest", "αTest"}, // Greek letter alpha (already lowercase)
			{"ΑTest", "αTest"}, // Greek letter Alpha (uppercase to lowercase)
		}

		for _, tc := range testCases {
			result := FirstLower(tc.input)
			if result != tc.expected {
				t.Errorf("FirstLower(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		}
	})

	t.Run("UnmarshalResource with large JSON", func(t *testing.T) {
		// Test with a realistic large resource
		largeResource := map[string]interface{}{
			"resourceType": "StructureDefinition",
			"url":          "http://hl7.org/fhir/StructureDefinition/Patient",
			"version":      "4.0.1",
			"name":         "Patient",
			"description":  strings.Repeat("This is a long description. ", 100),
			"element": []map[string]interface{}{
				{"path": "Patient.id", "min": 0, "max": "1"},
				{"path": "Patient.name", "min": 0, "max": "*"},
			},
		}

		data, err := json.Marshal(largeResource)
		if err != nil {
			t.Fatalf("failed to marshal test data: %v", err)
		}

		result, err := UnmarshalResource(data)
		if err != nil {
			t.Errorf("unexpected error with large resource: %v", err)
		}

		if result.ResourceType != "StructureDefinition" {
			t.Errorf("expected ResourceType 'StructureDefinition', got %q", result.ResourceType)
		}
	})
}

// Helper functions

func stringPtrEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
