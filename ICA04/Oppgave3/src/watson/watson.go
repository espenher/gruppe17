package watson

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WatsonResponse struct {
	Results []SpeechReckognizeResult
}

type SpeechReckognizeResult struct {
	Alternatives []SpeechReckognitionResult
	Final        bool
}

type SpeechReckognitionResult struct {
	Confidence int
	Transcript string
}

func Convert(filePath string) string {
	url := "https://gateway-lon.watsonplatform.net/speech-to-text/api/v1/recognize"

	soundFile, err := ioutil.ReadFile(filePath + ".flac")

	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(soundFile))

	auth, _ := ioutil.ReadFile("watson-credentials.creds.txt")

	req.SetBasicAuth("apikey", string(auth))

	req.Header.Add("Content-Type", "audio/flac")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result WatsonResponse
	json.Unmarshal([]byte(body), &result)

	return result.Results[0].Alternatives[0].Transcript
}
