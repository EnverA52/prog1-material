package Datentyp

type Time int

func FromSekunden(s int) Time {
	return Time(s)
}

func FromMinuten(m int) Time {
	return Time(m * 60)
}

func FromStunden(h int) Time {
	return Time(h * 60 * 60)
}

func (t Time) Sekunden() int {
	return int(t)
}

func (t Time) Minuten() int {
	return t.Sekunden() / 60
}
func (t Time) Stunden() int {
	return t.Sekunden() / 60 / 60
}
