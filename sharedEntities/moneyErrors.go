package SharedPackages

type CurrencyNotSupported struct {
}

func (*CurrencyNotSupported) Error() string {
	return ""
}

type CentAmountMustPresent struct {
}

func (*CentAmountMustPresent) Error() string {
	return ""
}

type TooHighFractionDigits struct {
}

func (*TooHighFractionDigits) Error() string {
	return ""
}

type TooLowFractionDigits struct {
}

func (*TooLowFractionDigits) Error() string {
	return ""
}
