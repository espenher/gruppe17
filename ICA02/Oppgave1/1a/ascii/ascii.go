package ascii

import ( 
		"fmt"
)

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// Oppgave 1b
// Implementer en funksjon som eksportere const ascii


// Funksjon tar en "string literal" med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
//
// Eksempel (utskriften kommer fra en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...


func IterateOverASCIIStringLiteral(s string) {
	// Kode for Oppgave 1a
	for i := 0; i < len(s); i++ {
			fmt.Printf("%X %+q %b \n", s[i], s[i], s[i])
	}
}

// Unix-like operating systems are known to use it as erase control character, i.e. to delete the previous character in the line mode. 

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen GreetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1c


func GreetingASCII() string {
	fmt.Printf("Hello :)")
	return "Hello :)" 
} 
