package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere 
// Returnerer en slice av type []byte
// 
func AllocateVar(b int) []byte {
	// Kode for Oppgave 5a

	var slice []byte
	slice = make([]byte, b)

	return slice;
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	// Kode for Oppgave 5a
	byte_Slice := make([]byte, b, b) // len(byte_Slice)=b, cap(byte_Slice)=b
	return byte_Slice;
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	var newSlice []byte = AllocateMake(uidx)
	var slice_copy []byte = CopySlice(slc)
	copy (slice_copy[lidx:uidx], newSlice)
	

	return newSlice; 
}

// CopySlice ???
func CopySlice(slc []byte) []byte{
	slice_copy := make([]byte, len(slc))
	copy(slc, slice_copy)
	return slice_copy;
}