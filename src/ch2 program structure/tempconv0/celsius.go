// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 39.
//!+

// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

//Define dos tipos custom
type Celsius float64
type Fahrenheit float64

//Define constantes con valores expresados en los tipos custom
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//Define una función que usa tipos customs
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
