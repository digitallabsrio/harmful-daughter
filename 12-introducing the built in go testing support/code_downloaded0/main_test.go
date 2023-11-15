package main
 
import (
    "errors"
    "testing"
)
 
var errorMessage = "width and height must be positive"
 
// Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
    if w < 1 || h < 1 {
      return 0, errors.New(errorMessage)
    }
    return w * h, nil
}

func TestCalcAreaSuccess(t *testing.T) {
    result, err := CalcArea(3, 5)
    if err != nil {
      t.Error("Expected CalcArea(3, 5) to succeed")
    } else if result != 15 {
      t.Errorf("CalcArea(3, 5) returned %d. Expected 15", result)
    }
}

func TestCalcAreaFail(t *testing.T) {
    _, err := CalcArea(-3, 6)
    if err == nil {
      t.Error("Expected CalcArea(-3, 6) to return an error")
    }
        
    if err.Error() != errorMessage {
      t.Error("Expected error to be: " + errorMessage)      
    } 
}
