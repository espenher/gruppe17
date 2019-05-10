package wit

import (
	"io/ioutil"
	"log"
	"os"

	witai "github.com/wit-ai/wit-go"
)

func Convert(filePath string) string {
	token, _ := ioutil.ReadFile("wit-credentials.creds.txt")
	client := witai.NewClient(string(token))

	file, _ := os.Open(filePath + ".wav")

	msg, err := client.Speech(&witai.MessageRequest{
		Speech: &witai.Speech{
			File:        file,
			ContentType: "audio/wav",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return msg.Text
}
