package utils

import (
	"encoding/json"
	"strings"
	"testing"
)

func PrettyPrintJSON(t *testing.T, jsonStr string, testName string) {
	t.Helper()

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		t.Fatalf("Invalid JSON in %s: %v", testName, err)
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("Failed to format JSON in %s: %v", testName, err)
	}

	t.Logf("\n=== %s JSON OUTPUT ===\n%s\n=== END %s ===\n",
		strings.ToUpper(testName), string(prettyJSON), strings.ToUpper(testName))
}

func ValidateJSONStructure(t *testing.T, jsonStr string, expectedFields []string) {
	t.Helper()

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}

	for _, field := range expectedFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Missing required field: %s", field)
		}
	}
}
