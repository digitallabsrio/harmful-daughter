package main

import (
    "testing"
)

func TestCalcAreaSuccess(t *testing.T) {
    result, err := CalcArea(3, 5)
    if err != nil {
      t.Error("CalcArea(3, 5) returned an error")
    } else if result != 15 {
      t.Errorf("CalcArea(3, 5) returned %d. Expected 15", result)
    }
}
