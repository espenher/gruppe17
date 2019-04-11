package unicode

import "fmt"

const nor = "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72"
const is = "\u0022\x6e\x6f\x72\u00f0\x75\x72\x20\x6f\x67\x20\x73\x75\u00f0\x75\x72\u0022"
const jp = "\u0022\xe5\x8c\x97\xE3\x81\xA8\xE5\x8D\x97\u0022"

// Kode for Oppgave 4a
func Translate(nor string, language string) string {
	if nor == "nord og sør" {
		if language == "jp" {
			return jp
		}
		if language == "is" {
			return is
		}
	}
	return ""
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
  fmt.Println("\xf0\x9F\x98\x80")
  fmt.Println("\xf0\x9F\x98\x97")
  // Demonstrerer at deler av et tegn representert med flere bytes
  // kan ikke tolkes innenfor koden (unicode)
  fmt.Println("\xf0\x9F\x98")
  fmt.Println("\xf0\x9F")
  fmt.Println("\xf0")
}



