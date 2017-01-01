package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoliingC      Celsius    = 100
	AbsoluteZeroK Kelvin     = 0
	FreezingK     Kelvin     = 273.15
	BoliingK      Kelvin     = 373.15
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoliingF      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%7.2f˚C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%7.2f˚F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%7.2f˚K", k) }
