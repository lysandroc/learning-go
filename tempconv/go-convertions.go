package tempconv

//CToF converts to a temperature from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts to a temperature from Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC convert to a temperature from Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k) - AbsoluteZeroC
}
