package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Gopher/converter"
)

func main() {
	for _, args := range os.Args[1:] {
		arg, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Converter: %v\n", err)
		}

		f := converter.Fahrenheit(arg)
		c := converter.Celsius(arg)
		fe := converter.Feet(arg)
		m := converter.Meter(arg)
		p := converter.Pound(arg)
		kg := converter.Kilogram(arg)

		fmt.Printf("%s = %s\n", f, converter.CtoF(c))
		fmt.Printf("%s = %s\n", c, converter.FtoC(f))
		fmt.Printf("%s = %s\n", fe, converter.MtoFe(m))
		fmt.Printf("%s = %s\n", m, converter.FetoM(fe))
		fmt.Printf("%s = %s\n", p, converter.KtoP(kg))
		fmt.Printf("%s = %s\n", kg, converter.PtoK(p))
	}
}
