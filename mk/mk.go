package mk

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type mk struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'mk' locale
func New() locales.Translator {
	return &mk{
		locale:                 "mk",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 5, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "ден.", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "јан.", "фев.", "мар.", "апр.", "мај", "јун.", "јул.", "авг.", "септ.", "окт.", "ноем.", "дек."},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануари", "февруари", "март", "април", "мај", "јуни", "јули", "август", "септември", "октомври", "ноември", "декември"},
		daysAbbreviated:        []string{"нед.", "пон.", "вт.", "сре.", "чет.", "пет.", "саб."},
		daysNarrow:             []string{"н", "п", "в", "с", "ч", "п", "с"},
		daysShort:              []string{"нед.", "пон.", "вто.", "сре.", "чет.", "пет.", "саб."},
		daysWide:               []string{"недела", "понеделник", "вторник", "среда", "четврток", "петок", "сабота"},
		periodsAbbreviated:     []string{"претпл.", "попл."},
		periodsNarrow:          []string{"претпл.", "попл."},
		periodsWide:            []string{"претпладне", "попладне"},
		erasAbbreviated:        []string{"пр.н.е.", "н.е."},
		erasNarrow:             []string{"пр.н.е.", "н.е."},
		erasWide:               []string{"пред нашата ера", "од нашата ера"},
		timezones:              map[string]string{"ACDT": "Летно сметање на времето во Централна Австралија", "ACST": "Стандардно време во Централна Австралија", "ACWDT": "Летно сметање на времето во Централна и Западна Австралија", "ACWST": "Стандардно време во Централна и Западна Австралија", "ADT": "Атлантско летно сметање на времето", "AEDT": "Летно сметање на времето во Источна Австралија", "AEST": "Стандардно време во Источна Австралија", "AKDT": "Летно сметање на времето во Алјаска", "AKST": "Стандардно време во Алјаска", "ARST": "Летно сметање на времето во Аргентина", "ART": "Стандардно време во Аргентина", "AST": "Атлантско стандардно време", "AWDT": "Летно сметање на времето во Западна Австралија", "AWST": "Стандардно време во Западна Австралија", "BOT": "Време во Боливија", "BT": "Време во Бутан", "CAT": "Средноафриканско време", "CDT": "Централно летно сметање на времето", "CHADT": "Летно сметање на времето во Чатам", "CHAST": "Стандардно време во Чатам", "CLST": "Летно сметање на времето во Чиле", "CLT": "Стандардно време во Чиле", "COST": "Летно сметање на времето во Колумбија", "COT": "Стандардно време во Колумбија", "CST": "Централно стандардно време во Северна Америка", "ChST": "Време во Чаморо", "EAT": "Источноафриканско време", "ECT": "Време во Еквадор", "EDT": "Источно летно сметање на времето", "EST": "Источно стандардно време", "GFT": "Време во Француска Гвајана", "GMT": "Средно време по Гринич", "GST": "Време на Мексиканскиот Залив", "GYT": "Време во Гвајана", "HADT": "Летно сметање на времето во Хаваи - Алеутски острови", "HAST": "Стандардно време во Хаваи - Алеутски острови", "HAT": "Летно сметање на времето на Њуфаундленд", "HECU": "Летно сметање на времето во Куба", "HEEG": "Летно сметање на времето во Источен Гренланд", "HENOMX": "Летно сметање на времето во северозападно Мексико", "HEOG": "Летно сметање на времето во Западен Гренланд", "HEPM": "Летно сметање на времето на Сент Пјер и Микелан", "HEPMX": "Летно пацифичко време во Мексико", "HKST": "Летно време во Хонг Конг", "HKT": "Стандардно време во Хонг Конг", "HNCU": "Стандардно време во Куба", "HNEG": "Стандардно време во Источен Гренланд", "HNNOMX": "Стандардно време во северозападно Мексико", "HNOG": "Стандардно време во Западен Гренланд", "HNPM": "Стандардно време на Сент Пјер и Микелан", "HNPMX": "Стандардно пацифичко време во Мексико", "HNT": "Стандардно време на Њуфаундленд", "IST": "Време во Индија", "JDT": "Летно сметање на времето во Јапонија", "JST": "Стандардно време во Јапонија", "LHDT": "Летно сметање на времето во Лорд Хау", "LHST": "Стандардно време во Лорд Хау", "MDT": "Планинско летно сметање на времето", "MESZ": "Средноевропско летно време", "MEZ": "Средноевропско стандардно време", "MST": "Планинско стандардно време", "MYT": "Време во Малезија", "NZDT": "Летно сметање на времето во Нов Зеланд", "NZST": "Стандардно време во Нов Зеланд", "OESZ": "Источноевропско летно време", "OEZ": "Источноевропско стандардно време", "PDT": "Пацифичко летно сметање на времето", "PST": "Пацифичко стандардно време", "SAST": "Време во Јужноафриканска Република", "SGT": "Време во Сингапур", "SRT": "Време во Суринам", "TMST": "Летно време во Туркменистан", "TMT": "Стандардно време во Туркменистан", "UYST": "Летно сметање на времето во Уругвај", "UYT": "Стандардно време во Уругвај", "VET": "Време во Венецуела", "WARST": "Летно сметање на времето во западна Аргентина", "WART": "Стандардно време во западна Аргентина", "WAST": "Западноафриканско летно сметање на времето", "WAT": "Западноафриканско стандардно време", "WESZ": "Западноевропско летно време", "WEZ": "Западноевропско стандардно време", "WIB": "Време во Западна Индонезија", "WIT": "Време во Источна Индонезија", "WITA": "Време во Централна Индонезија", "∅∅∅": "Летно време на Азорските Острови"},
	}
}

// Locale returns the current translators string locale
func (mk *mk) Locale() string {
	return mk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mk'
func (mk *mk) PluralsCardinal() []locales.PluralRule {
	return mk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mk'
func (mk *mk) PluralsOrdinal() []locales.PluralRule {
	return mk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mk'
func (mk *mk) PluralsRange() []locales.PluralRule {
	return mk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mk'
func (mk *mk) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mk'
func (mk *mk) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if iMod10 == 2 && iMod100 != 12 {
		return locales.PluralRuleTwo
	} else if (iMod10 == 7 || iMod10 == 8) && (iMod100 != 17 && iMod100 != 18) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mk'
func (mk *mk) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mk *mk) MonthAbbreviated(month time.Month) string {
	return mk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mk *mk) MonthsAbbreviated() []string {
	return mk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mk *mk) MonthNarrow(month time.Month) string {
	return mk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mk *mk) MonthsNarrow() []string {
	return mk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mk *mk) MonthWide(month time.Month) string {
	return mk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mk *mk) MonthsWide() []string {
	return mk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mk *mk) WeekdayAbbreviated(weekday time.Weekday) string {
	return mk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mk *mk) WeekdaysAbbreviated() []string {
	return mk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mk *mk) WeekdayNarrow(weekday time.Weekday) string {
	return mk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mk *mk) WeekdaysNarrow() []string {
	return mk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mk *mk) WeekdayShort(weekday time.Weekday) string {
	return mk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mk *mk) WeekdaysShort() []string {
	return mk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mk *mk) WeekdayWide(weekday time.Weekday) string {
	return mk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mk *mk) WeekdaysWide() []string {
	return mk.daysWide
}

// Decimal returns the decimal point of number
func (mk *mk) Decimal() string {
	return mk.decimal
}

// Group returns the group of number
func (mk *mk) Group() string {
	return mk.group
}

// Group returns the minus sign of number
func (mk *mk) Minus() string {
	return mk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mk' and handles both Whole and Real numbers based on 'v'
func (mk *mk) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mk' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mk *mk) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mk.percentSuffix...)

	b = append(b, mk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mk'
func (mk *mk) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, mk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mk'
// in accounting notation.
func (mk *mk) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mk.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, mk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, mk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, mk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'mk'
func (mk *mk) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'mk'
func (mk *mk) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mk'
func (mk *mk) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mk'
func (mk *mk) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mk'
func (mk *mk) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mk'
func (mk *mk) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mk'
func (mk *mk) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mk'
func (mk *mk) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
