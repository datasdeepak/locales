package fa_IR

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fa_IR struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'fa_IR' locale
func New() locales.Translator {
	return &fa_IR{
		locale:                 "fa_IR",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                "٫",
		group:                  "٬",
		minus:                  "‎−",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: "‎",
		currencyNegativePrefix: "‎",
		monthsAbbreviated:      []string{"", "ژانویهٔ", "فوریهٔ", "مارس", "آوریل", "مهٔ", "ژوئن", "ژوئیهٔ", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
		monthsNarrow:           []string{"", "ژ", "ف", "م", "آ", "م", "ژ", "ژ", "ا", "س", "ا", "ن", "د"},
		monthsWide:             []string{"", "ژانویهٔ", "فوریهٔ", "مارس", "آوریل", "مهٔ", "ژوئن", "ژوئیهٔ", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
		daysAbbreviated:        []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		daysNarrow:             []string{"ی", "د", "س", "چ", "پ", "ج", "ش"},
		daysShort:              []string{"۱ش", "۲ش", "۳ش", "۴ش", "۵ش", "ج", "ش"},
		daysWide:               []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		periodsAbbreviated:     []string{"ق.ظ.", "ب.ظ."},
		periodsNarrow:          []string{"ق", "ب"},
		periodsWide:            []string{"قبل\u200cازظهر", "بعدازظهر"},
		erasAbbreviated:        []string{"ق.م.", "م."},
		erasNarrow:             []string{"ق", "م"},
		erasWide:               []string{"قبل از میلاد", "میلادی"},
		timezones:              map[string]string{"ACDT": "وقت تابستانی مرکز استرالیا", "ACST": "وقت عادی مرکز استرالیا", "ACWDT": "وقت تابستانی مرکز-غرب استرالیا", "ACWST": "وقت عادی مرکز-غرب استرالیا", "ADT": "وقت تابستانی آتلانتیک", "AEDT": "وقت تابستانی شرق استرالیا", "AEST": "وقت عادی شرق استرالیا", "AKDT": "وقت تابستانی آلاسکا", "AKST": "وقت عادی آلاسکا", "ARST": "وقت تابستانی آرژانتین", "ART": "وقت عادی آرژانتین", "AST": "وقت عادی آتلانتیک", "AWDT": "وقت تابستانی غرب استرالیا", "AWST": "وقت عادی غرب استرالیا", "BOT": "وقت بولیوی", "BT": "وقت بوتان", "CAT": "وقت مرکز افریقا", "CDT": "وقت تابستانی مرکز امریکا", "CHADT": "وقت تابستانی چت\u200cهام", "CHAST": "وقت عادی چت\u200cهام", "CLST": "وقت تابستانی شیلی", "CLT": "وقت عادی شیلی", "COST": "وقت تابستانی کلمبیا", "COT": "وقت عادی کلمبیا", "CST": "وقت عادی مرکز امریکا", "ChST": "وقت عادی چامورو", "EAT": "وقت شرق افریقا", "ECT": "وقت اکوادور", "EDT": "وقت تابستانی شرق امریکا", "EST": "وقت عادی شرق امریکا", "GFT": "وقت گویان فرانسه", "GMT": "وقت گرینویچ", "GST": "وقت عادی خلیج فارس", "GYT": "وقت گویان", "HADT": "وقت تابستانی هاوایی‐الوشن", "HAST": "وقت عادی هاوایی‐الوشن", "HAT": "وقت تابستانی نیوفاندلند", "HECU": "وقت تابستانی کوبا", "HEEG": "وقت تابستانی شرق گرینلند", "HENOMX": "وقت تابستانی شمال غرب مکزیک", "HEOG": "وقت تابستانی غرب گرینلند", "HEPM": "وقت تابستانی سنت\u200cپیر و میکلون", "HEPMX": "وقت تابستانی شرق مکزیک", "HKST": "وقت تابستانی هنگ\u200cکنگ", "HKT": "وقت عادی هنگ\u200cکنگ", "HNCU": "وقت عادی کوبا", "HNEG": "وقت عادی شرق گرینلند", "HNNOMX": "وقت عادی شمال غرب مکزیک", "HNOG": "وقت عادی غرب گرینلند", "HNPM": "وقت عادی سنت\u200cپیر و میکلون", "HNPMX": "وقت عادی شرق مکزیک", "HNT": "وقت عادی نیوفاندلند", "IST": "وقت هند", "JDT": "وقت تابستانی ژاپن", "JST": "وقت عادی ژاپن", "LHDT": "وقت تابستانی لردهو", "LHST": "وقت عادی لردهو", "MDT": "وقت تابستانی کوهستانی امریکا", "MESZ": "وقت تابستانی مرکز اروپا", "MEZ": "وقت عادی مرکز اروپا", "MST": "وقت عادی کوهستانی امریکا", "MYT": "وقت مالزی", "NZDT": "وقت تابستانی زلاند نو", "NZST": "وقت عادی زلاند نو", "OESZ": "وقت تابستانی شرق اروپا", "OEZ": "وقت عادی شرق اروپا", "PDT": "وقت تابستانی غرب امریکا", "PST": "وقت عادی غرب امریکا", "SAST": "وقت عادی جنوب افریقا", "SGT": "وقت سنگاپور", "SRT": "وقت سورینام", "TMST": "وقت تابستانی ترکمنستان", "TMT": "وقت عادی ترکمنستان", "UYST": "وقت تابستانی اروگوئه", "UYT": "وقت عادی اروگوئه", "VET": "وقت ونزوئلا", "WARST": "وقت تابستانی غرب آرژانتین", "WART": "وقت عادی غرب آرژانتین", "WAST": "وقت تابستانی غرب افریقا", "WAT": "وقت عادی غرب افریقا", "WESZ": "وقت تابستانی غرب اروپا", "WEZ": "وقت عادی غرب اروپا", "WIB": "وقت غرب اندونزی", "WIT": "وقت شرق اندونزی", "WITA": "وقت مرکز اندونزی", "∅∅∅": "وقت تابستانی برازیلیا"},
	}
}

// Locale returns the current translators string locale
func (fa *fa_IR) Locale() string {
	return fa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fa_IR'
func (fa *fa_IR) PluralsCardinal() []locales.PluralRule {
	return fa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fa_IR'
func (fa *fa_IR) PluralsOrdinal() []locales.PluralRule {
	return fa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fa_IR'
func (fa *fa_IR) PluralsRange() []locales.PluralRule {
	return fa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fa_IR'
func (fa *fa_IR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fa_IR'
func (fa *fa_IR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fa_IR'
func (fa *fa_IR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := fa.CardinalPluralRule(num1, v1)
	end := fa.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fa *fa_IR) MonthAbbreviated(month time.Month) string {
	return fa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fa *fa_IR) MonthsAbbreviated() []string {
	return fa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fa *fa_IR) MonthNarrow(month time.Month) string {
	return fa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fa *fa_IR) MonthsNarrow() []string {
	return fa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fa *fa_IR) MonthWide(month time.Month) string {
	return fa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fa *fa_IR) MonthsWide() []string {
	return fa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fa *fa_IR) WeekdayAbbreviated(weekday time.Weekday) string {
	return fa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fa *fa_IR) WeekdaysAbbreviated() []string {
	return fa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fa *fa_IR) WeekdayNarrow(weekday time.Weekday) string {
	return fa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fa *fa_IR) WeekdaysNarrow() []string {
	return fa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fa *fa_IR) WeekdayShort(weekday time.Weekday) string {
	return fa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fa *fa_IR) WeekdaysShort() []string {
	return fa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fa *fa_IR) WeekdayWide(weekday time.Weekday) string {
	return fa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fa *fa_IR) WeekdaysWide() []string {
	return fa.daysWide
}

// Decimal returns the decimal point of number
func (fa *fa_IR) Decimal() string {
	return fa.decimal
}

// Group returns the group of number
func (fa *fa_IR) Group() string {
	return fa.group
}

// Group returns the minus sign of number
func (fa *fa_IR) Minus() string {
	return fa.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fa_IR' and handles both Whole and Real numbers based on 'v'
func (fa *fa_IR) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fa_IR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fa *fa_IR) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 10
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fa_IR'
func (fa *fa_IR) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(symbol) + 11 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fa.group) - 1; j >= 0; j-- {
					b = append(b, fa.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, fa.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fa_IR'
// in accounting notation.
func (fa *fa_IR) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fa.currencies[currency]
	l := len(s) + len(symbol) + 11 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fa.decimal) - 1; j >= 0; j-- {
				b = append(b, fa.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fa.group) - 1; j >= 0; j-- {
					b = append(b, fa.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(fa.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyNegativePrefix[j])
		}

		for j := len(fa.minus) - 1; j >= 0; j-- {
			b = append(b, fa.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(fa.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, fa.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fa_IR'
func (fa *fa_IR) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := fa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
