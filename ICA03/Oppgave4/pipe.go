// https://www.socketloop.com/references/golang-io-pipe-function-example
package main



import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ExampleEncode(str []byte) ([]byte) {
	dst := make([]byte, hex.EncodedLen(len(str)))
	hex.Encode(dst, str)

	return dst
}

func ExampleDecode(bstr []byte) {
	dst := make([]byte, hex.DecodedLen(len(bstr)))
	n, err := hex.Decode(dst, bstr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])

	// Output:
	// Hello Gopher!
}

func main() {

	// read and write with pipe
	pReader, pWriter := io.Pipe()

	// use base64 encoder
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)

	gWriter := gzip.NewWriter(b64Writer)

	// write text to be gzipped and push to pipe
	go func() {
		fmt.Println("Start writing")
		n, err := gWriter.Write([]byte("These words will be compressed and pushed into pipe"))

		fmt.Printf("len = %d\n", n)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		gWriter.Close()
		b64Writer.Close()
		pWriter.Close()
		fmt.Printf("Written %d bytes \n", n)
	}()

	// start reading from the pipe(reverse of the above process)

	// use base64 decoder to wrap pipe Reader
	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)

	// read gzipped text and decompressed the text
	gReader, err := gzip.NewReader(b64Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Look at the final output at the other side of the pipe

	// print out the text
	text, err := ioutil.ReadAll(gReader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", text)
	hex_utf8 := ExampleEncode(text)
	fmt.Printf("%s\n", hex_utf8)

	sEnc := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(sEnc)


}
