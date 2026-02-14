package converter

import "fmt"

type Fahrenheit float64
type Celsius float64
type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g F\n", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%g C\n", c) }
func (m Meter) String() string      { return fmt.Sprintf("%g M\n", m) }
func (fe Feet) String() string      { return fmt.Sprintf("%g Fe\n", fe) }
func (p Pound) String() string      { return fmt.Sprintf("%g P\n", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g k\n", k) }

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }
func MtoFe(m Meter) Feet        { return Feet(m * 3.28084) }
func FetoM(fe Feet) Meter       { return Meter(fe * 0.3048) }
func PtoK(p Pound) Kilogram     { return Kilogram(p * 0.453592) }
func KtoP(kg Kilogram) Pound    { return Pound(kg * 2.205) }
