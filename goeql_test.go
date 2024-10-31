package goeql

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Test EncryptedText Serialization
func TestEncryptedText_Serialize(t *testing.T) {
	et := EncryptedText("Hello, World!")
	table := "test_table"
	column := "test_column"

	serializedData, err := et.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	deserializedData, err := et.Deserialize(serializedData)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if !reflect.DeepEqual(deserializedData, et) {
		t.Errorf("Expected deserialized value to be '%s', got '%s'", et, deserializedData)
	}
}

// Test EncryptedText Deserialization
func TestEncryptedText_Deserialize(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "Hello, World!",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var et EncryptedText
	deserialized, err := et.Deserialize(data)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if deserialized != EncryptedText("Hello, World!") {
		t.Errorf("Expected deserialized value to be 'Hello, World!', got '%s'", deserialized)
	}
}

// Test EncryptedJsonb Serialization
func TestEncryptedJsonb_Serialize(t *testing.T) {
	// You must cast any int to float64 to get the correct JSON output
	// Deserialization will always return a float64 for ints as json.Unmarshal will
	// convert them to float64 by default
	ej := EncryptedJsonb{
		"name":      "Alice",
		"age":       float64(30),
		"is_member": true,
	}

	table := "test_table"
	column := "test_column"

	serializedData, err := ej.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	deserializedData, err := ej.Deserialize(serializedData)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if !reflect.DeepEqual(deserializedData, ej) {
		t.Errorf("Expected deserialized value to be '%s', got '%s'", ej, deserializedData)
	}
}
func TestNullEncryptedJsonb_Serialize(t *testing.T) {
	ej := EncryptedJsonb(nil)

	table := "test_table"
	column := "test_column"

	serializedData, err := ej.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	deserializedData, err := ej.Deserialize(serializedData)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if !reflect.DeepEqual(deserializedData, ej) {
		t.Errorf("Expected deserialized value to be '%s', got '%s'", ej, deserializedData)
	}
}
func TestNullEncryptedText_Serialize(t *testing.T) {
	ej := EncryptedText("")

	table := "test_table"
	column := "test_column"

	serializedData, err := ej.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	deserializedData, err := ej.Deserialize(serializedData)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if !reflect.DeepEqual(deserializedData, ej) {
		t.Errorf("Expected deserialized value to be '%s', got '%s'", ej, deserializedData)
	}
}

// Test EncryptedJsonb Deserialization
func TestEncryptedJsonb_Deserialize(t *testing.T) {
	originalData := map[string]interface{}{
		"name":      "Alice",
		"age":       float64(30),
		"is_member": true,
	}

	jsonString, err := json.Marshal(originalData)
	if err != nil {
		t.Fatalf("Error marshaling original data: %v", err)
	}

	ec := EncryptedColumn{
		K: "pt",
		P: string(jsonString),
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var ej EncryptedJsonb
	deserialized, err := ej.Deserialize(data)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if !reflect.DeepEqual(deserialized, EncryptedJsonb(originalData)) {
		t.Errorf("Deserialized data does not match original data")
	}
}

// Test EncryptedInt Serialization
func TestEncryptedInt_Serialize(t *testing.T) {
	ei := EncryptedInt(42)
	table := "test_table"
	column := "test_column"

	serializedData, err := ei.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	expectedP := "42"
	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}
}

// Test EncryptedInt Deserialization
func TestEncryptedInt_Deserialize(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "42",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var ei EncryptedInt
	deserialized, err := ei.Deserialize(data)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if deserialized != EncryptedInt(42) {
		t.Errorf("Expected deserialized value to be 42, got %d", deserialized)
	}
}

// Test EncryptedBool Serialization
func TestEncryptedBool_Serialize(t *testing.T) {
	eb := EncryptedBool(true)
	table := "test_table"
	column := "test_column"

	serializedData, err := eb.Serialize(table, column)
	if err != nil {
		t.Fatalf("Serialize returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != "true" {
		t.Errorf("Expected P to be 'true', got '%s'", ec.P)
	}
}

// Test EncryptedBool Deserialization
func TestEncryptedBool_Deserialize(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "true",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var eb EncryptedBool
	deserialized, err := eb.Deserialize(data)
	if err != nil {
		t.Fatalf("Deserialize returned error: %v", err)
	}

	if deserialized != EncryptedBool(true) {
		t.Errorf("Expected deserialized value to be true, got %v", deserialized)
	}
}

func TestMatchQuerySerialization(t *testing.T) {
	value := "test_string"
	table := "table1"
	column := "column1"
	expectedP := "test_string"
	expectedQ := "match"

	serializedData, err := MatchQuery(value, table, column)
	if err != nil {
		t.Fatalf("SerializeQuery returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}

	if ec.Q != expectedQ {
		t.Errorf("Expected Q to be '%s', got '%s'", expectedQ, ec.Q)
	}
}
func TestOreQuerySerialization(t *testing.T) {
	value := 123
	table := "table1"
	column := "column1"
	expectedP := "123"
	expectedQ := "ore"

	serializedData, err := OreQuery(value, table, column)
	if err != nil {
		t.Fatalf("SerializeQuery returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}

	if ec.Q != expectedQ {
		t.Errorf("Expected Q to be '%s', got '%s'", expectedQ, ec.Q)
	}
}

func TestUniqueQuerySerialization(t *testing.T) {
	value := true
	table := "table1"
	column := "column1"
	expectedP := "true"
	expectedQ := "unique"

	serializedData, err := UniqueQuery(value, table, column)
	if err != nil {
		t.Fatalf("SerializeQuery returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}

	if ec.Q != expectedQ {
		t.Errorf("Expected Q to be '%s', got '%s'", expectedQ, ec.Q)
	}
}

func TestJsonbQuerySerialization(t *testing.T) {
	value := map[string]interface{}{"key": "value"}
	table := "table1"
	column := "column1"
	expectedP := `{"key":"value"}`
	expectedQ := "ste_vec"

	serializedData, err := JsonbQuery(value, table, column)
	if err != nil {
		t.Fatalf("SerializeQuery returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}
	if ec.Q != expectedQ {
		t.Errorf("Expected Q to be '%s', got '%s'", expectedQ, ec.Q)
	}
}

func TestEJsonPathQueryQuerySerialization(t *testing.T) {
	value := "$.top"
	table := "table1"
	column := "column1"
	expectedP := "$.top"
	expectedQ := "ejson_path"

	serializedData, err := EJsonPathQuery(value, table, column)
	if err != nil {
		t.Fatalf("SerializeQuery returned error: %v", err)
	}

	var ec EncryptedColumn
	if err := json.Unmarshal(serializedData, &ec); err != nil {
		t.Fatalf("Error unmarshaling serialized data: %v", err)
	}

	if ec.P != expectedP {
		t.Errorf("Expected P to be '%s', got '%s'", expectedP, ec.P)
	}
	if ec.Q != expectedQ {
		t.Errorf("Expected Q to be '%s', got '%s'", expectedQ, ec.Q)
	}
}

// Test ToEncryptedColumn Function
func TestToEncryptedColumn(t *testing.T) {
	tests := []struct {
		value     interface{}
		table     string
		column    string
		expectedP string
	}{
		{value: "test_string", table: "table1", column: "column1", expectedP: "test_string"},
		{value: 123, table: "table2", column: "column2", expectedP: "123"},
		{value: 123.456, table: "table3", column: "column3", expectedP: "123.456000"},
		{value: true, table: "table4", column: "column4", expectedP: "true"},
		{value: map[string]interface{}{"key": "value"}, table: "table5", column: "column5", expectedP: `{"key":"value"}`},
	}

	for _, tt := range tests {
		ec, err := ToEncryptedColumn(tt.value, tt.table, tt.column, nil)
		if err != nil {
			t.Fatalf("ToEncryptedColumn returned error: %v", err)
		}

		if ec.P != tt.expectedP {
			t.Errorf("Expected P to be '%s', got '%s'", tt.expectedP, ec.P)
		}
	}
}

// Test convertToString Function
func TestConvertToString(t *testing.T) {
	// the ptr value for unitptr will change between tests
	// set a mock value to use to confirm the convertToString function works on this uintptr.
	var mockPtr uintptr = 1374390189136
	tests := []struct {
		value       interface{}
		expectedStr string
		expectError bool
	}{
		{value: "test_string", expectedStr: "test_string", expectError: false},
		{value: 123, expectedStr: "123", expectError: false},
		{value: 123.456, expectedStr: "123.456000", expectError: false},
		{value: int(1), expectedStr: "1", expectError: false},
		{value: int8(-128), expectedStr: "-128", expectError: false},
		{value: int8(127), expectedStr: "127", expectError: false},
		{value: int16(-32768), expectedStr: "-32768", expectError: false},
		{value: int16(32767), expectedStr: "32767", expectError: false},
		{value: int32(-2147483648), expectedStr: "-2147483648", expectError: false},
		{value: int32(2147483647), expectedStr: "2147483647", expectError: false},
		{value: int64(-9223372036854775808), expectedStr: "-9223372036854775808", expectError: false},
		{value: int64(9223372036854775807), expectedStr: "9223372036854775807", expectError: false},
		{value: uint(1), expectedStr: "1", expectError: false},
		{value: uint8(255), expectedStr: "255", expectError: false},
		{value: uint16(65535), expectedStr: "65535", expectError: false},
		{value: uint32(4294967295), expectedStr: "4294967295", expectError: false},
		{value: uint64(18446744073709551615), expectedStr: "18446744073709551615", expectError: false},
		{value: float32(123.456), expectedStr: "123.456001", expectError: false},
		{value: float32(-10.543), expectedStr: "-10.543000", expectError: false},
		{value: float64(3.1425), expectedStr: "3.142500", expectError: false},
		{value: float64(-2.7182), expectedStr: "-2.718200", expectError: false},
		{value: mockPtr, expectedStr: "1374390189136", expectError: false}, //uinttpr type
		{value: true, expectedStr: "true", expectError: false},
		{value: map[string]interface{}{"key": "value"}, expectedStr: `{"key":"value"}`, expectError: false},
		{value: []int{1, 2, 3}, expectedStr: "[1, 2, 3]", expectError: false},
		{value: []float64{1.1, 2.2, 3.3}, expectedStr: "[1.100000, 2.200000, 3.300000]", expectError: false},
		{value: []string{"hello", "world"}, expectedStr: "[hello, world]", expectError: false},
		{value: []bool{true, false, true}, expectedStr: "[true, false, true]", expectError: false},
	}

	for _, tt := range tests {
		str, err := convertToString(tt.value)
		if tt.expectError {
			if err == nil {
				t.Errorf("Expected error for value: %v, but got none", tt.value)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for value: %v, error: %v", tt.value, err)
			} else if str != tt.expectedStr {
				t.Errorf("Expected '%s', got '%s' for value: %v for type: '%s'", tt.expectedStr, str, tt.value, reflect.TypeOf(tt.value))
			}
		}
	}
}

// Test EncryptedInt Deserialization Error
func TestEncryptedInt_Deserialize_Error(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "not_an_integer",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var ei EncryptedInt
	_, err = ei.Deserialize(data)
	if err == nil {
		t.Errorf("Expected error during Deserialize, but got none")
	}
}

// Test EncryptedBool Deserialization Error
func TestEncryptedBool_Deserialize_Error(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "not_a_boolean",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var eb EncryptedBool
	_, err = eb.Deserialize(data)
	if err == nil {
		t.Errorf("Expected error during Deserialize, but got none")
	}
}

// Test EncryptedJsonb Deserialization Error
func TestEncryptedJsonb_Deserialize_Error(t *testing.T) {
	ec := EncryptedColumn{
		K: "pt",
		P: "invalid_json",
		I: TableColumn{T: "test_table", C: "test_column"},
		V: 1,
	}

	data, err := json.Marshal(ec)
	if err != nil {
		t.Fatalf("Error marshaling EncryptedColumn: %v", err)
	}

	var ej EncryptedJsonb
	_, err = ej.Deserialize(data)
	if err == nil {
		t.Errorf("Expected error during Deserialize, but got none")
	}
}
