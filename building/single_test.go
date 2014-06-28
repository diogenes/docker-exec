package building

import "testing"

func TestSingleBuild(t *testing.T) {
	expected_string := "-p 8080"
	single_builder := NewSingleBuilder("-p", "8080")
	real_string := single_builder.Build()
	if expected_string != real_string {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_string, real_string)
	}
}
