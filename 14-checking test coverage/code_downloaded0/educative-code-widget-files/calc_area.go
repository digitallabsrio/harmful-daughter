package main

import (
    "errors"
)

var errorMessage = "width and height must be positive"

// Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
    if w == 0 {
      return 0, errors.New("the width can't be zero")  
    }    
    if h == 0 {
      return 0, errors.New("the height can't be zero")  
    }
    if w < 0 {
      return 0, errors.New("the width can't be negative")  
    }
    
    if h < 0 {
      return 0, errors.New("the height can't be negative")  
    }

    return w * h, nil
}
