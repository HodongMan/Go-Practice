package tempconv

import "fmt"

type Celcius float64
type Fahrenheit float64
type Kelvin float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
    ZeroKelvin Celsius = 0
)

func (c Celcius) String() string { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k)}
