// Package googlecloud uses Google Cloud Speech-to-text API to convert mono flac files to text
package googlecloud

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Convert(filePath string) string {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "google-credentials.creds.json")

	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filePath + ".flac")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
			ProfanityFilter: false,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	//Assume first result and result alternative is correct.
	return resp.Results[0].Alternatives[0].Transcript

	// Prints the results.
	/*for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}*/
}
