package main

import "testing"

func TestCommandsList(t *testing.T) {
  commands := CommandsCollection()
  if len(commands) == 0 {
    t.Error("No one command added to collection")
  }
}

func TestCommandsNames(t *testing.T) {
  commands := CommandsCollection()
  original_names := []string{"hook"}
  expected_names := make([]string, len(commands))
  for i, command := range commands {
    expected_names[i] = command.Name
  }

  if !testEqSlice(expected_names, original_names) {
    t.Errorf("Not all commands presents %v, %v", expected_names, original_names)
  }
}

func testEqSlice(a, b []string) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}
