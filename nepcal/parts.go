package nepcal

// Month represents a B.S. month much like time.Month represents a Gregorian month.
type Month int

// List of B.S. Months.
const (
	Baisakh Month = 1 + iota
	Jestha
	Ashar
	Shrawan
	Bhadra
	Ashoj
	Kartik
	Mangshir
	Poush
	Magh
	Falgun
	Chaitra
)

// NumDays returns the total number of days in the month. Since B.S. dates are
// are not guaranteed to have the same number of days in a month every year,
// this method takes in the 'year' value as a parameter.
//
// Note that the 'year' value should be in the supported B.S. year
// range (2000 - 2090) which can be checked using the 'IsInRangeYear' method.
// This method will return an ErrOutOfBounds if it is not in that range.
func (m Month) NumDays(year int) (int, error) {
	if !IsInRangeYear(year) {
		return -1, ErrOutOfBounds
	}

	return m.numDaysUnchecked(year), nil
}

// An unchecked variant of NumDays. This is only for private uses through the
// Time struct when it is certain that the date in question is in range.
func (m Month) numDaysUnchecked(yy int) int {
	// Invariant: int(m) - 1 is between 0 and 11.
	return bsDaysInMonthsByYear[yy][int(m)-1]
}

// Name returns valid UTF-8 encoded human readable names for this month.
func (m Month) Name() string {
	names := map[Month]string{
		Baisakh:  "बैशाख",
		Jestha:   "जेठ",
		Ashar:    "असार",
		Shrawan:  "साउन",
		Bhadra:   "भदौ",
		Ashoj:    "असोज",
		Kartik:   "कार्तिक",
		Mangshir: "मंसिर",
		Poush:    "पौष",
		Magh:     "माघ",
		Falgun:   "फागुन",
		Chaitra:  "चैत",
	}

	// Invariant: the month always exists in the map.
	v, _ := names[m]

	return v
}

// String implements the Stringer interface for Month.
func (m Month) String() string {
	return m.Name()
}

// Weekday represents a B.S. weekday much like time.Weekday represents a Gregorian weekday.
// This is actually equivalent to time.Weekday's enumerations, but we avoid wrapping
// that type; a little copying is better.
type Weekday int

// Enumerations of the Weekday.
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Name returns valid UTF-8 encoded human readable names for this weekday.
func (w Weekday) Name() string {
	names := map[Weekday]string{
		Sunday:    "आइतबार",
		Monday:    "सोमबार",
		Tuesday:   "मंगलबार",
		Wednesday: "बुधबार",
		Thursday:  "बिहिबार",
		Friday:    "शुक्रबार",
		Saturday:  "शनिबार",
	}

	// Invariant: the day always exists in the map.
	v, _ := names[w]

	return v
}

// String implements the Stringer interface for Month.
func (w Weekday) String() string {
	return w.Name()
}
