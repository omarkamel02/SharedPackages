package SharedPackages

type CurrencyNotSupported struct {
}

func (*CurrencyNotSupported) Error() string {
	return "This Currency code isn't supported"
}

type BadMoneyValue struct {
}

func (*BadMoneyValue) Error() string {
	return "something is wrong with the money values"
}
