package listen

// Reverse erwartet eine Liste und dreht
// die Reihenfolge der Elemente um.
// Liefert die umgedrehte Liste zurück.
// Die Funktion muss rekursiv sein.
func Reverse(list []int) []int {
	// TODO
	if len(list) == 0 {
		return []int{}
	}
	return append(Reverse(list[1:]), list[0])
}

// DuplicateElements erwartet eine Liste und dupliziert
// jedes Element darin.
// Liefert die Ergebnisliste zurück.
// Die Funktion muss rekursiv sein.
func DuplicateElements(list []int) []int {
	// TODO
	if len(list) == 0 {
		return []int{}
	}

	return append([]int{list[0], list[0]}, DuplicateElements(list[1:])...)
}
