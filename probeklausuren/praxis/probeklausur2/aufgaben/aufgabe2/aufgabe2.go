package aufgabe2

import "slices"

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	result := ""
	// TODO
	ziffern := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := range s {
		if !slices.Contains(ziffern, s[i:i+1]) {
			result = result + s[i:i+1]
		}
	}
	return result
}
