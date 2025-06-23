package cmd

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/csotherden/golang-fhir-models/fhir-models-gen/fhir"
	"github.com/dave/jennifer/jen"
)

func TestGenerateValueSet(t *testing.T) {
	tests := []struct {
		name          string
		resources     ResourceMap
		valueSet      fhir.ValueSet
		expectError   bool
		errorContains string
	}{
		{
			name: "successful generation",
			resources: ResourceMap{
				"CodeSystem": {
					"http://hl7.org/fhir/administrative-gender": createValidCodeSystemJSON(),
				},
			},
			valueSet:    createValidValueSet(),
			expectError: false,
		},
		{
			name:          "ValueSet without name",
			resources:     ResourceMap{},
			valueSet:      fhir.ValueSet{},
			expectError:   true,
			errorContains: "ValueSet without name",
		},
		{
			name:      "ValueSet without compose",
			resources: ResourceMap{},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
			},
			expectError:   true,
			errorContains: "doens't include any CodeSystems",
		},
		{
			name:      "ValueSet with empty includes",
			resources: ResourceMap{},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{},
				},
			},
			expectError:   true,
			errorContains: "doens't include any CodeSystems",
		},
		{
			name:      "ValueSet with multiple includes",
			resources: ResourceMap{},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{System: stringPtr("http://system1.com")},
						{System: stringPtr("http://system2.com")},
					},
				},
			},
			expectError:   true,
			errorContains: "includes more than one CodeSystem",
		},
		{
			name:      "ValueSet with empty canonical URL",
			resources: ResourceMap{},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{}, // empty include
					},
				},
			},
			expectError:   true,
			errorContains: "doens't include any CodeSystems",
		},
		{
			name: "missing CodeSystem in resources",
			resources: ResourceMap{
				"CodeSystem": {}, // empty map
			},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{System: stringPtr("http://missing.system")},
					},
				},
			},
			expectError:   true,
			errorContains: "missing CodeSystem with canonical URL",
		},
		{
			name: "invalid CodeSystem JSON",
			resources: ResourceMap{
				"CodeSystem": {
					"http://invalid.system": []byte("invalid json"),
				},
			},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{System: stringPtr("http://invalid.system")},
					},
				},
			},
			expectError:   true,
			errorContains: "", // will be JSON unmarshal error
		},
		{
			name: "CodeSystem with no concepts",
			resources: ResourceMap{
				"CodeSystem": {
					"http://empty.system": createEmptyCodeSystemJSON(),
				},
			},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{System: stringPtr("http://empty.system")},
					},
				},
			},
			expectError:   true,
			errorContains: "has no codes",
		},
		{
			name: "system with version",
			resources: ResourceMap{
				"CodeSystem": {
					"http://hl7.org/fhir/administrative-gender|4.0.1": createValidCodeSystemJSON(),
				},
			},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Url:  stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{
							System:  stringPtr("http://hl7.org/fhir/administrative-gender"),
							Version: stringPtr("4.0.1"),
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "nested concepts",
			resources: ResourceMap{
				"CodeSystem": {
					"http://test.system": createNestedCodeSystemJSON(),
				},
			},
			valueSet: fhir.ValueSet{
				Name: stringPtr("TestValueSet"),
				Url:  stringPtr("http://test.valueset"),
				Compose: &fhir.ValueSetCompose{
					Include: []fhir.ValueSetComposeInclude{
						{System: stringPtr("http://test.system")},
					},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := generateValueSet(tt.resources, tt.valueSet)

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

			if file == nil {
				t.Errorf("expected file to be non-nil")
				return
			}

			// Verify the generated code can be rendered
			code := file.GoString()
			if code == "" {
				t.Errorf("generated code is empty")
				return
			}

			// Basic structure checks
			if tt.valueSet.Name != nil {
				expectedType := "type " + *tt.valueSet.Name + " int"
				if !strings.Contains(code, expectedType) {
					t.Errorf("expected type definition '%s' not found in generated code", expectedType)
				}

				// Check for generated methods
				expectedMethods := []string{
					"func (code " + *tt.valueSet.Name + ") MarshalJSON()",
					"func (code *" + *tt.valueSet.Name + ") UnmarshalJSON(",
					"func (code " + *tt.valueSet.Name + ") String()",
					"func (code " + *tt.valueSet.Name + ") Code()",
					"func (code " + *tt.valueSet.Name + ") Display()",
					"func (code " + *tt.valueSet.Name + ") Definition()",
				}

				for _, method := range expectedMethods {
					if !strings.Contains(code, method) {
						t.Errorf("expected method '%s' not found in generated code", method)
					}
				}
			}
		})
	}
}

func TestCanonical(t *testing.T) {
	tests := []struct {
		name     string
		include  fhir.ValueSetComposeInclude
		expected string
	}{
		{
			name: "system without version",
			include: fhir.ValueSetComposeInclude{
				System: stringPtr("http://hl7.org/fhir/administrative-gender"),
			},
			expected: "http://hl7.org/fhir/administrative-gender",
		},
		{
			name: "system with version",
			include: fhir.ValueSetComposeInclude{
				System:  stringPtr("http://hl7.org/fhir/administrative-gender"),
				Version: stringPtr("4.0.1"),
			},
			expected: "http://hl7.org/fhir/administrative-gender|4.0.1",
		},
		{
			name:     "empty include",
			include:  fhir.ValueSetComposeInclude{},
			expected: "",
		},
		{
			name: "nil system",
			include: fhir.ValueSetComposeInclude{
				System: nil,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canonical(tt.include)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestCodeIdentifier(t *testing.T) {
	tests := []struct {
		name         string
		valueSetName string
		code         string
		expected     string
	}{
		{
			name:         "equals operator",
			valueSetName: "TestValueSet",
			code:         "=",
			expected:     "TestValueSetEquals",
		},
		{
			name:         "not equals operator",
			valueSetName: "TestValueSet",
			code:         "!=",
			expected:     "TestValueSetNotEquals",
		},
		{
			name:         "greater than operator",
			valueSetName: "TestValueSet",
			code:         ">",
			expected:     "TestValueSetGreaterThan",
		},
		{
			name:         "less than operator",
			valueSetName: "TestValueSet",
			code:         "<",
			expected:     "TestValueSetLessThan",
		},
		{
			name:         "greater or equals operator",
			valueSetName: "TestValueSet",
			code:         ">=",
			expected:     "TestValueSetGreaterOrEquals",
		},
		{
			name:         "less or equals operator",
			valueSetName: "TestValueSet",
			code:         "<=",
			expected:     "TestValueSetLessOrEquals",
		},
		{
			name:         "regular code",
			valueSetName: "TestValueSet",
			code:         "male",
			expected:     "TestValueSetMale",
		},
		{
			name:         "code with hyphens",
			valueSetName: "TestValueSet",
			code:         "some-code",
			expected:     "TestValueSetSomeCode",
		},
		{
			name:         "code with dots",
			valueSetName: "TestValueSet",
			code:         "some.code",
			expected:     "TestValueSetSome_Code",
		},
		{
			name:         "mixed case code",
			valueSetName: "TestValueSet",
			code:         "mixedCase",
			expected:     "TestValueSetMixedCase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := codeIdentifier(tt.valueSetName, tt.code)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// Helper functions for creating test data

func stringPtr(s string) *string {
	return &s
}

func createValidValueSet() fhir.ValueSet {
	return fhir.ValueSet{
		Name: stringPtr("AdministrativeGender"),
		Url:  stringPtr("http://hl7.org/fhir/ValueSet/administrative-gender"),
		Compose: &fhir.ValueSetCompose{
			Include: []fhir.ValueSetComposeInclude{
				{
					System: stringPtr("http://hl7.org/fhir/administrative-gender"),
				},
			},
		},
	}
}

func createValidCodeSystemJSON() []byte {
	codeSystem := fhir.CodeSystem{
		Concept: []fhir.CodeSystemConcept{
			{
				Code:       "male",
				Display:    stringPtr("Male"),
				Definition: stringPtr("Male gender"),
			},
			{
				Code:       "female",
				Display:    stringPtr("Female"),
				Definition: stringPtr("Female gender"),
			},
		},
	}
	data, _ := json.Marshal(codeSystem)
	return data
}

func createEmptyCodeSystemJSON() []byte {
	codeSystem := fhir.CodeSystem{
		Concept: []fhir.CodeSystemConcept{}, // empty concepts
	}
	data, _ := json.Marshal(codeSystem)
	return data
}

func createNestedCodeSystemJSON() []byte {
	codeSystem := fhir.CodeSystem{
		Concept: []fhir.CodeSystemConcept{
			{
				Code:       "parent1",
				Display:    stringPtr("Parent 1"),
				Definition: stringPtr("Parent concept 1"),
				Concept: []fhir.CodeSystemConcept{
					{
						Code:       "child1",
						Display:    stringPtr("Child 1"),
						Definition: stringPtr("Child concept 1"),
					},
					{
						Code:       "child2",
						Display:    stringPtr("Child 2"),
						Definition: stringPtr("Child concept 2"),
					},
				},
			},
			{
				Code:       "parent2",
				Display:    stringPtr("Parent 2"),
				Definition: stringPtr("Parent concept 2"),
			},
		},
	}
	data, _ := json.Marshal(codeSystem)
	return data
}

// Benchmark tests

func BenchmarkGenerateValueSet(b *testing.B) {
	resources := ResourceMap{
		"CodeSystem": {
			"http://hl7.org/fhir/administrative-gender": createValidCodeSystemJSON(),
		},
	}
	valueSet := createValidValueSet()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := generateValueSet(resources, valueSet)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCanonical(b *testing.B) {
	include := fhir.ValueSetComposeInclude{
		System:  stringPtr("http://hl7.org/fhir/administrative-gender"),
		Version: stringPtr("4.0.1"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = canonical(include)
	}
}

func BenchmarkCodeIdentifier(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = codeIdentifier("TestValueSet", "test-code.value")
	}
}

// Table-driven tests for edge cases

func TestGenerateValueSetEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		setupFunc   func() (ResourceMap, fhir.ValueSet)
		expectError bool
		validate    func(t *testing.T, file *jen.File, err error)
	}{
		{
			name: "special characters in codes",
			setupFunc: func() (ResourceMap, fhir.ValueSet) {
				specialCodeSystem := fhir.CodeSystem{
					Concept: []fhir.CodeSystemConcept{
						{Code: "=", Display: stringPtr("Equals")},
						{Code: "!=", Display: stringPtr("Not Equals")},
						{Code: ">", Display: stringPtr("Greater Than")},
						{Code: "<", Display: stringPtr("Less Than")},
						{Code: ">=", Display: stringPtr("Greater Or Equals")},
						{Code: "<=", Display: stringPtr("Less Or Equals")},
					},
				}
				data, _ := json.Marshal(specialCodeSystem)

				return ResourceMap{
						"CodeSystem": {
							"http://special.system": data,
						},
					}, fhir.ValueSet{
						Name: stringPtr("SpecialCodes"),
						Url:  stringPtr("http://special.valueset"),
						Compose: &fhir.ValueSetCompose{
							Include: []fhir.ValueSetComposeInclude{
								{System: stringPtr("http://special.system")},
							},
						},
					}
			},
			expectError: false,
			validate: func(t *testing.T, file *jen.File, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
					return
				}
				code := file.GoString()

				// Check that special operators are properly converted
				expectedConstants := []string{
					"SpecialCodesEquals",
					"SpecialCodesNotEquals",
					"SpecialCodesGreaterThan",
					"SpecialCodesLessThan",
					"SpecialCodesGreaterOrEquals",
					"SpecialCodesLessOrEquals",
				}

				for _, constant := range expectedConstants {
					if !strings.Contains(code, constant) {
						t.Errorf("expected constant '%s' not found in generated code", constant)
					}
				}
			},
		},
		{
			name: "concepts with missing display and definition",
			setupFunc: func() (ResourceMap, fhir.ValueSet) {
				minimalCodeSystem := fhir.CodeSystem{
					Concept: []fhir.CodeSystemConcept{
						{Code: "code1"}, // no display or definition
						{Code: "code2", Display: stringPtr("Display 2")},       // no definition
						{Code: "code3", Definition: stringPtr("Definition 3")}, // no display
					},
				}
				data, _ := json.Marshal(minimalCodeSystem)

				return ResourceMap{
						"CodeSystem": {
							"http://minimal.system": data,
						},
					}, fhir.ValueSet{
						Name: stringPtr("MinimalCodes"),
						Url:  stringPtr("http://minimal.valueset"),
						Compose: &fhir.ValueSetCompose{
							Include: []fhir.ValueSetComposeInclude{
								{System: stringPtr("http://minimal.system")},
							},
						},
					}
			},
			expectError: false,
			validate: func(t *testing.T, file *jen.File, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
					return
				}
				code := file.GoString()

				// Verify that the code generation handles missing fields gracefully
				if !strings.Contains(code, "MinimalCodesCode1") {
					t.Errorf("expected constant 'MinimalCodesCode1' not found")
				}
				if !strings.Contains(code, "MinimalCodesCode2") {
					t.Errorf("expected constant 'MinimalCodesCode2' not found")
				}
				if !strings.Contains(code, "MinimalCodesCode3") {
					t.Errorf("expected constant 'MinimalCodesCode3' not found")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resources, valueSet := tt.setupFunc()
			file, err := generateValueSet(resources, valueSet)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if tt.validate != nil {
				tt.validate(t, file, err)
			}
		})
	}
}

// Additional unit tests for helper functions

func TestConstsRoot(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{
			Code: "first",
			Concept: []fhir.CodeSystemConcept{
				{Code: "child1"},
				{Code: "child2"},
			},
		},
		{Code: "second"},
	}

	file := jen.NewFile("test")
	file.Const().DefsFunc(constsRoot("TestType", concepts))

	code := file.GoString()

	// Check that iota is used correctly
	if !strings.Contains(code, "TestTypeFirst TestType = iota") {
		t.Errorf("expected iota declaration not found in: %s", code)
	}

	// Check that all constants are generated
	expectedConstants := []string{"TestTypeFirst", "TestTypeChild1", "TestTypeChild2", "TestTypeSecond"}
	for _, constant := range expectedConstants {
		if !strings.Contains(code, constant) {
			t.Errorf("expected constant '%s' not found in generated code", constant)
		}
	}
}

func TestConsts(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "first"},
		{Code: "second"},
	}

	file := jen.NewFile("test")
	file.Const().DefsFunc(consts("TestType", concepts))

	code := file.GoString()

	expectedConstants := []string{"TestTypeFirst", "TestTypeSecond"}
	for _, constant := range expectedConstants {
		if !strings.Contains(code, constant) {
			t.Errorf("expected constant '%s' not found in generated code", constant)
		}
	}
}

func TestUnmarshalRoot(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "test1"},
		{Code: "test2"},
	}

	file := jen.NewFile("test")
	file.Func().Id("testFunc").Params().Block(
		jen.Switch(jen.Id("s")).BlockFunc(unmarshalRoot("TestType", concepts)),
	)

	code := file.GoString()

	// Check that cases are generated
	if !strings.Contains(code, `case "test1"`) {
		t.Errorf("expected case for test1 not found")
	}
	if !strings.Contains(code, `case "test2"`) {
		t.Errorf("expected case for test2 not found")
	}

	// Check that default case with error is generated
	if !strings.Contains(code, "default:") {
		t.Errorf("expected default case not found")
	}
	if !strings.Contains(code, "fmt.Errorf") {
		t.Errorf("expected error generation not found")
	}
}

func TestUnmarshal(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "test1"},
		{Code: "test2"},
	}

	file := jen.NewFile("test")
	file.Func().Id("testFunc").Params().Block(
		jen.Switch(jen.Id("s")).BlockFunc(unmarshal("TestType", concepts)),
	)

	code := file.GoString()

	// Check that cases are generated
	if !strings.Contains(code, `case "test1"`) {
		t.Errorf("expected case for test1 not found")
	}
	if !strings.Contains(code, `*code = TestTypeTest1`) {
		t.Errorf("expected assignment for test1 not found")
	}
}

func TestCodes(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "test1"},
		{Code: "test2"},
	}

	file := jen.NewFile("test")
	file.Func().Id("testFunc").Params().Block(
		jen.Switch(jen.Id("code")).BlockFunc(codes("TestType", concepts)),
	)

	code := file.GoString()

	// Check that cases return the correct codes
	if !strings.Contains(code, "case TestTypeTest1:") {
		t.Errorf("expected case for TestTypeTest1 not found")
	}
	if !strings.Contains(code, `return "test1"`) {
		t.Errorf("expected return for test1 not found")
	}
}

func TestDisplays(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "test1", Display: stringPtr("Display 1")},
		{Code: "test2"}, // no display
	}

	file := jen.NewFile("test")
	file.Func().Id("testFunc").Params().Block(
		jen.Switch(jen.Id("code")).BlockFunc(displays("TestType", concepts)),
	)

	code := file.GoString()

	// Check that only concepts with display generate cases
	if !strings.Contains(code, "case TestTypeTest1:") {
		t.Errorf("expected case for TestTypeTest1 not found")
	}
	if !strings.Contains(code, `return "Display 1"`) {
		t.Errorf("expected return for Display 1 not found")
	}

	// test2 should not have a case since it has no display
	if strings.Contains(code, "case TestTypeTest2:") {
		t.Errorf("unexpected case for TestTypeTest2 found")
	}
}

func TestDefinitions(t *testing.T) {
	concepts := []fhir.CodeSystemConcept{
		{Code: "test1", Definition: stringPtr("Definition 1")},
		{Code: "test2"}, // no definition
	}

	file := jen.NewFile("test")
	file.Func().Id("testFunc").Params().Block(
		jen.Switch(jen.Id("code")).BlockFunc(definitions("TestType", concepts)),
	)

	code := file.GoString()

	// Check that only concepts with definition generate cases
	if !strings.Contains(code, "case TestTypeTest1:") {
		t.Errorf("expected case for TestTypeTest1 not found")
	}
	if !strings.Contains(code, `return "Definition 1"`) {
		t.Errorf("expected return for Definition 1 not found")
	}

	// test2 should not have a case since it has no definition
	if strings.Contains(code, "case TestTypeTest2:") {
		t.Errorf("unexpected case for TestTypeTest2 found")
	}
}

// Integration test for complete code generation
func TestGenerateValueSetIntegration(t *testing.T) {
	// Create a comprehensive CodeSystem with all features
	codeSystem := fhir.CodeSystem{
		Concept: []fhir.CodeSystemConcept{
			{
				Code:       "active",
				Display:    stringPtr("Active"),
				Definition: stringPtr("The resource is active"),
				Concept: []fhir.CodeSystemConcept{
					{
						Code:       "fully-active",
						Display:    stringPtr("Fully Active"),
						Definition: stringPtr("The resource is fully active"),
					},
				},
			},
			{
				Code:       "inactive",
				Display:    stringPtr("Inactive"),
				Definition: stringPtr("The resource is inactive"),
			},
			{
				Code: "pending", // no display or definition
			},
		},
	}
	data, _ := json.Marshal(codeSystem)

	resources := ResourceMap{
		"CodeSystem": {
			"http://example.com/status": data,
		},
	}

	valueSet := fhir.ValueSet{
		Name: stringPtr("ResourceStatus"),
		Url:  stringPtr("http://example.com/ValueSet/status"),
		Compose: &fhir.ValueSetCompose{
			Include: []fhir.ValueSetComposeInclude{
				{System: stringPtr("http://example.com/status")},
			},
		},
	}

	file, err := generateValueSet(resources, valueSet)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	code := file.GoString()

	// Verify complete structure
	expectedElements := []string{
		"type ResourceStatus int",
		"const (",
		"ResourceStatusActive ResourceStatus = iota",
		"ResourceStatusFullyActive",
		"ResourceStatusInactive",
		"ResourceStatusPending",
		"func (code ResourceStatus) MarshalJSON()",
		"func (code *ResourceStatus) UnmarshalJSON(",
		"func (code ResourceStatus) String()",
		"func (code ResourceStatus) Code()",
		"func (code ResourceStatus) Display()",
		"func (code ResourceStatus) Definition()",
		`case "active":`,
		`case "fully-active":`,
		`case "inactive":`,
		`case "pending":`,
		`return "Active"`,
		`return "Inactive"`,
		`return "The resource is active"`,
		`return "The resource is inactive"`,
	}

	for _, element := range expectedElements {
		if !strings.Contains(code, element) {
			t.Errorf("expected element '%s' not found in generated code", element)
		}
	}
}

// Test error handling in generated methods
func TestGeneratedCodeErrorHandling(t *testing.T) {
	codeSystem := fhir.CodeSystem{
		Concept: []fhir.CodeSystemConcept{
			{Code: "test"},
		},
	}
	data, _ := json.Marshal(codeSystem)

	resources := ResourceMap{
		"CodeSystem": {
			"http://test.system": data,
		},
	}

	valueSet := fhir.ValueSet{
		Name: stringPtr("TestType"),
		Url:  stringPtr("http://test.valueset"),
		Compose: &fhir.ValueSetCompose{
			Include: []fhir.ValueSetComposeInclude{
				{System: stringPtr("http://test.system")},
			},
		},
	}

	file, err := generateValueSet(resources, valueSet)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	code := file.GoString()

	// Check that unknown cases return appropriate defaults
	if !strings.Contains(code, `return "<unknown>"`) {
		t.Errorf("expected default return value '<unknown>' not found")
	}

	// Check that UnmarshalJSON has proper error handling
	if !strings.Contains(code, "fmt.Errorf") {
		t.Errorf("expected error handling in UnmarshalJSON not found")
	}
}
