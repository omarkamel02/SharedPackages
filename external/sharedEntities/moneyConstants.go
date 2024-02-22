package SharedPackages

var CurrencyFractions = map[string]int64{
	"AED": 2, // United Arab Emirates Dirham
	"AFN": 2, // Afghan Afghani
	"ALL": 2, // Albanian Lek
	"AMD": 2, // Armenian Dram
	"ANG": 2, // Netherlands Antillean Guilder
	"AOA": 2, // Angolan Kwanza
	"ARS": 2, // Argentine Peso
	"AUD": 2, // Australian Dollar
	"AWG": 2, // Aruban Florin
	"AZN": 2, // Azerbaijani Manat
	"BAM": 2, // Bosnia-Herzegovina Convertible Mark
	"BBD": 2, // Barbadian Dollar
	"BDT": 2, // Bangladeshi Taka
	"BGN": 2, // Bulgarian Lev
	"BHD": 3, // Bahraini Dinar
	"BIF": 0, // Burundian Franc
	"BMD": 2, // Bermudian Dollar
	"BND": 2, // Brunei Dollar
	"BOB": 2, // Bolivian Boliviano
	"BRL": 2, // Brazilian Real
	"BSD": 2, // Bahamian Dollar
	"BTN": 2, // Bhutanese Ngultrum
	"BWP": 2, // Botswana Pula
	"BYN": 2, // Belarusian Ruble
	"BZD": 2, // Belize Dollar
	"CAD": 2, // Canadian Dollar
	"CDF": 2, // Congolese Franc
	"CHF": 2, // Swiss Franc
	"CLP": 0, // Chilean Peso
	"CNY": 2, // Chinese Yuan
	"COP": 2, // Colombian Peso
	"CRC": 2, // Costa Rican Colón
	"CUC": 2, // Cuban Convertible Peso
	"CUP": 2, // Cuban Peso
	"CVE": 0, // Cape Verdean Escudo
	"CZK": 2, // Czech Koruna
	"DJF": 0, // Djiboutian Franc
	"DKK": 2, // Danish Krone
	"DOP": 2, // Dominican Peso
	"DZD": 2, // Algerian Dinar
	"EGP": 2, // Egyptian Pound
	"ERN": 2, // Eritrean Nakfa
	"ETB": 2, // Ethiopian Birr
	"EUR": 2, // Euro
	"FJD": 2, // Fijian Dollar
	"FKP": 2, // Falkland Islands Pound
	"GBP": 2, // British Pound Sterling
	"GEL": 2, // Georgian Lari
	"GGP": 2, // Guernsey Pound
	"GHS": 2, // Ghanaian Cedi
	"GIP": 2, // Gibraltar Pound
	"GMD": 2, // Gambian Dalasi
	"GNF": 0, // Guinean Franc
	"GTQ": 2, // Guatemalan Quetzal
	"GYD": 2, // Guyanaese Dollar
	"HKD": 2, // Hong Kong Dollar
	"HNL": 2, // Honduran Lempira
	"HRK": 2, // Croatian Kuna
	"HTG": 2, // Haitian Gourde
	"HUF": 2, // Hungarian Forint
	"IDR": 2, // Indonesian Rupiah
	"ILS": 2, // Israeli New Sheqel
	"IMP": 2, // Isle of Man Pound
	"INR": 2, // Indian Rupee
	"IQD": 3, // Iraqi Dinar
	"IRR": 2, // Iranian Rial
	"ISK": 0, // Icelandic Króna
	"JEP": 2, // Jersey Pound
	"JMD": 2, // Jamaican Dollar
	"JOD": 3, // Jordanian Dinar
	"JPY": 0, // Japanese Yen
	"KES": 2, // Kenyan Shilling
	"KGS": 2, // Kyrgystani Som
	"KHR": 2, // Cambodian Riel
	"KID": 2, // Kiribati Dollar
	"KMF": 0, // Comorian Franc
	"KRW": 0, // South Korean Won
	"KWD": 3, // Kuwaiti Dinar
	"KYD": 2, // Cayman Islands Dollar
	"KZT": 2, // Kazakhstani Tenge
	"LAK": 2, // Laotian Kip
	"LBP": 2, // Lebanese Pound
	"LKR": 2, // Sri Lankan Rupee
	"LRD": 2, // Liberian Dollar
	"LSL": 2, // Lesotho Loti
	"LYD": 3, // Libyan Dinar
	"MAD": 2, // Moroccan Dirham
	"MDL": 2, // Moldovan Leu
	"MGA": 2, // Malagasy Ariary
	"MKD": 2, // Macedonian Denar
	"MMK": 2, // Myanma Kyat
	"MNT": 2, // Mongolian Tugrik
	"MOP": 2, // Macanese Pataca
	"MRO": 2, // Mauritanian Ouguiya
	"MRU": 2, // Mauritanian Ouguiya
	"MUR": 2, // Mauritian Rupee
	"MVR": 2, // Maldivian Rufiyaa
	"MWK": 2, // Malawian Kwacha
	"MXN": 2, // Mexican Peso
	"MYR": 2, // Malaysian Ringgit
	"MZN": 2, // Mozambican Metical
	"NAD": 2, // Namibian Dollar
	"NGN": 2, // Nigerian Naira
	"NIO": 2, // Nicaraguan Córdoba
	"NOK": 2, // Norwegian Krone
	"NPR": 2, // Nepalese Rupee
	"NZD": 2, // New Zealand Dollar
	"OMR": 3, // Omani Rial
	"PAB": 2, // Panamanian Balboa
	"PEN": 2, // Peruvian Nuevo Sol
	"PGK": 2, // Papua New Guinean Kina
	"PHP": 2, // Philippine Peso
	"PKR": 2, // Pakistani Rupee
	"PLN": 2, // Polish Zloty
	"PYG": 0, // Paraguayan Guarani
	"QAR": 2, // Qatari Rial
	"RON": 2, // Romanian Leu
	"RSD": 2, // Serbian Dinar
	"RUB": 2, // Russian Ruble
	"RWF": 0, // Rwandan Franc
	"SAR": 2, // Saudi Riyal
	"SBD": 2, // Solomon Islands Dollar
	"SCR": 2, // Seychellois Rupee
	"SDG": 2, // Sudanese Pound
	"SEK": 2, // Swedish Krona
	"SGD": 2, // Singapore Dollar
	"SHP": 2, // Saint Helena Pound
	"SLL": 2, // Sierra Leonean Leone
	"SOS": 2, // Somali Shilling
	"SRD": 2, // Surinamese Dollar
	"SSP": 2, // South Sudanese Pound
	"STN": 2, // São Tomé and Príncipe Dobra
	"SVC": 2, // Salvadoran Colón
	"SYP": 2, // Syrian Pound
	"SZL": 2, // Swazi Lilangeni
	"THB": 2, // Thai Baht
	"TJS": 2, // Tajikistani Somoni
	"TMT": 2, // Turkmenistani Manat
	"TND": 3, // Tunisian Dinar
	"TOP": 2, // Tongan Pa'anga
	"TRY": 2, // Turkish Lira
	"TTD": 2, // Trinidad and Tobago Dollar
	"TWD": 2, // New Taiwan Dollar
	"TZS": 2, // Tanzanian Shilling
	"UAH": 2, // Ukrainian Hryvnia
	"UGX": 0, // Ugandan Shilling
	"USD": 2, // United States Dollar
	"UYU": 2, // Uruguayan Peso
	"UZS": 2, // Uzbekistan Som
	"VES": 2, // Venezuelan Bolívar Soberano
	"VND": 0, // Vietnamese Dong
	"VUV": 0, // Vanuatu Vatu
	"WST": 2, // Samoan Tala
	"XAF": 0, // CFA Franc BEAC
	"XAG": 0, // Silver Ounce
	"XAU": 0, // Gold Ounce
	"XCD": 2, // East Caribbean Dollar
	"XDR": 0, // Special Drawing Rights
	"XOF": 0, // CFA Franc BCEAO
	"XPD": 0, // Palladium Ounce
	"XPF": 0, // CFP Franc
	"XPT": 0, // Platinum Ounce
	"YER": 2, // Yemeni Rial
	"ZAR": 2, // South African Rand
	"ZMW": 2, // Zambian Kwacha
	"ZWL": 2, // Zimbabwean Dollar
}

const (
	AED = "AED" // United Arab Emirates Dirham
	AFN = "AFN" // Afghan Afghani
	ALL = "ALL" // Albanian Lek
	AMD = "AMD" // Armenian Dram
	ANG = "ANG" // Netherlands Antillean Guilder
	AOA = "AOA" // Angolan Kwanza
	ARS = "ARS" // Argentine Peso
	AUD = "AUD" // Australian Dollar
	AWG = "AWG" // Aruban Florin
	AZN = "AZN" // Azerbaijani Manat
	BAM = "BAM" // Bosnia-Herzegovina Convertible Mark
	BBD = "BBD" // Barbadian Dollar
	BDT = "BDT" // Bangladeshi Taka
	BGN = "BGN" // Bulgarian Lev
	BHD = "BHD" // Bahraini Dinar
	BIF = "BIF" // Burundian Franc
	BMD = "BMD" // Bermudian Dollar
	BND = "BND" // Brunei Dollar
	BOB = "BOB" // Bolivian Boliviano
	BRL = "BRL" // Brazilian Real
	BSD = "BSD" // Bahamian Dollar
	BTC = "BTC" // Bitcoin
	BTN = "BTN" // Bhutanese Ngultrum
	BWP = "BWP" // Botswana Pula
	BYN = "BYN" // Belarusian Ruble
	BZD = "BZD" // Belize Dollar
	CAD = "CAD" // Canadian Dollar
	CDF = "CDF" // Congolese Franc
	CHF = "CHF" // Swiss Franc
	CLP = "CLP" // Chilean Peso
	CNY = "CNY" // Chinese Yuan
	COP = "COP" // Colombian Peso
	CRC = "CRC" // Costa Rican Colón
	CUC = "CUC" // Cuban Convertible Peso
	CUP = "CUP" // Cuban Peso
	CVE = "CVE" // Cape Verdean Escudo
	CZK = "CZK" // Czech Koruna
	DJF = "DJF" // Djiboutian Franc
	DKK = "DKK" // Danish Krone
	DOP = "DOP" // Dominican Peso
	DZD = "DZD" // Algerian Dinar
	EGP = "EGP" // Egyptian Pound
	ERN = "ERN" // Eritrean Nakfa
	ETB = "ETB" // Ethiopian Birr
	EUR = "EUR" // Euro
	FJD = "FJD" // Fijian Dollar
	FKP = "FKP" // Falkland Islands Pound
	GBP = "GBP" // British Pound Sterling
	GEL = "GEL" // Georgian Lari
	GGP = "GGP" // Guernsey Pound
	GHS = "GHS" // Ghanaian Cedi
	GIP = "GIP" // Gibraltar Pound
	GMD = "GMD" // Gambian Dalasi
	GNF = "GNF" // Guinean Franc
	GTQ = "GTQ" // Guatemalan Quetzal
	GYD = "GYD" // Guyanaese Dollar
	HKD = "HKD" // Hong Kong Dollar
	HNL = "HNL" // Honduran Lempira
	HRK = "HRK" // Croatian Kuna
	HTG = "HTG" // Haitian Gourde
	HUF = "HUF" // Hungarian Forint
	IDR = "IDR" // Indonesian Rupiah
	ILS = "ILS" // Israeli New Sheqel
	IMP = "IMP" // Isle of Man Pound
	INR = "INR" // Indian Rupee
	IQD = "IQD" // Iraqi Dinar
	IRR = "IRR" // Iranian Rial
	ISK = "ISK" // Icelandic Króna
	JEP = "JEP" // Jersey Pound
	JMD = "JMD" // Jamaican Dollar
	JOD = "JOD" // Jordanian Dinar
	JPY = "JPY" // Japanese Yen
	KES = "KES" // Kenyan Shilling
	KGS = "KGS" // Kyrgystani Som
	KHR = "KHR" // Cambodian Riel
	KID = "KID" // Kiribati Dollar
	KMF = "KMF" // Comorian Franc
	KRW = "KRW" // South Korean Won
	KWD = "KWD" // Kuwaiti Dinar
	KYD = "KYD" // Cayman Islands Dollar
	KZT = "KZT" // Kazakhstani Tenge
	LAK = "LAK" // Laotian Kip
	LBP = "LBP" // Lebanese Pound
	LKR = "LKR" // Sri Lankan Rupee
	LRD = "LRD" // Liberian Dollar
	LSL = "LSL" // Lesotho Loti
	LYD = "LYD" // Libyan Dinar
	MAD = "MAD" // Moroccan Dirham
	MDL = "MDL" // Moldovan Leu
	MGA = "MGA" // Malagasy Ariary
	MKD = "MKD" // Macedonian Denar
	MMK = "MMK" // Myanma Kyat
	MNT = "MNT" // Mongolian Tugrik
	MOP = "MOP" // Macanese Pataca
	MRO = "MRO" // Mauritanian Ouguiya
	MRU = "MRU" // Mauritanian Ouguiya
	MUR = "MUR" // Mauritian Rupee
	MVR = "MVR" // Maldivian Rufiyaa
	MWK = "MWK" // Malawian Kwacha
	MXN = "MXN" // Mexican Peso
	MYR = "MYR" // Malaysian Ringgit
	MZN = "MZN" // Mozambican Metical
	NAD = "NAD" // Namibian Dollar
	NGN = "NGN" // Nigerian Naira
	NIO = "NIO" // Nicaraguan Córdoba
	NOK = "NOK" // Norwegian Krone
	NPR = "NPR" // Nepalese Rupee
	NZD = "NZD" // New Zealand Dollar
	OMR = "OMR" // Omani Rial
	PAB = "PAB" // Panamanian Balboa
	PEN = "PEN" // Peruvian Nuevo Sol
	PGK = "PGK" // Papua New Guinean Kina
	PHP = "PHP" // Philippine Peso
	PKR = "PKR" // Pakistani Rupee
	PLN = "PLN" // Polish Zloty
	PYG = "PYG" // Paraguayan Guarani
	QAR = "QAR" // Qatari Rial
	RON = "RON" // Romanian Leu
	RSD = "RSD" // Serbian Dinar
	RUB = "RUB" // Russian Ruble
	RWF = "RWF" // Rwandan Franc
	SAR = "SAR" // Saudi Riyal
	SBD = "SBD" // Solomon Islands Dollar
	SCR = "SCR" // Seychellois Rupee
	SDG = "SDG" // Sudanese Pound
	SEK = "SEK" // Swedish Krona
	SGD = "SGD" // Singapore Dollar
	SHP = "SHP" // Saint Helena Pound
	SLL = "SLL" // Sierra Leonean Leone
	SOS = "SOS" // Somali Shilling
	SRD = "SRD" // Surinamese Dollar
	SSP = "SSP" // South Sudanese Pound
	STD = "STD" // São Tomé and Príncipe Dobra
	STN = "STN" // São Tomé and Príncipe Dobra
	SVC = "SVC" // Salvadoran Colón
	SYP = "SYP" // Syrian Pound
	SZL = "SZL" // Swazi Lilangeni
	THB = "THB" // Thai Baht
	TJS = "TJS" // Tajikistani Somoni
	TMT = "TMT" // Turkmenistani Manat
	TND = "TND" // Tunisian Dinar
	TOP = "TOP" // Tongan Pa'anga
	TRY = "TRY" // Turkish Lira
	TTD = "TTD" // Trinidad and Tobago Dollar
	TWD = "TWD" // New Taiwan Dollar
	TZS = "TZS" // Tanzanian Shilling
	UAH = "UAH" // Ukrainian Hryvnia
	UGX = "UGX" // Ugandan Shilling
	USD = "USD" // United States Dollar
	UYU = "UYU" // Uruguayan Peso
	UZS = "UZS" // Uzbekistan Som
	VEF = "VEF" // Venezuelan Bolívar Fuerte
	VES = "VES" // Venezuelan Bolívar Soberano
	VND = "VND" // Vietnamese Dong
	VUV = "VUV" // Vanuatu Vatu
	WST = "WST" // Samoan Tala
	XAF = "XAF" // CFA Franc BEAC
	XAG = "XAG" // Silver Ounce
	XAU = "XAU" // Gold Ounce
	XCD = "XCD" // East Caribbean Dollar
	XDR = "XDR" // Special Drawing Rights
	XOF = "XOF" // CFA Franc BCEAO
	XPD = "XPD" // Palladium Ounce
	XPF = "XPF" // CFP Franc
	XPT = "XPT" // Platinum Ounce
	YER = "YER" // Yemeni Rial
	ZAR = "ZAR" // South African Rand
	ZMW = "ZMW" // Zambian Kwacha
	ZWL = "ZWL" // Zimbabwean Dollar
)
