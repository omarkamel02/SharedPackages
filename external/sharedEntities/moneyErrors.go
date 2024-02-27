package SharedPackages

type CurrencyNotSupported struct {
}

func (*CurrencyNotSupported) Error() string {
	return ""
}

type BadMoneyValue struct {
}

func (*BadMoneyValue) Error() string {
	return ""
}
