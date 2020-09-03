package main

import (
	"flag"
	"fmt"
)

// Celsius contiene la temperatura en celsius
type Celsius float64

// Fahrenheit contiene la temperatura en Fahrenheit
type Fahrenheit float64

//CToF .
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

//FToC .
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

//Metodo String es parte del interface que usa Flag
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{ Celsius }

//Con esto *celsiusFlag ya implementa completamente el interface
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	//*celsiusFlag implementa el interface que es el primer argumento de este método
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
