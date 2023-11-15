package main

import (
    "errors"
)

var errorMessage = "width and height must be positive"

// Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
    if w < 1 || h < 1 {
      return 0, errors.New(errorMessage)
    }
    return w * h, nil
}
