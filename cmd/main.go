package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/giovane-aG/ai-assistant/internal/client"
	"github.com/giovane-aG/ai-assistant/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup flags so that we can run this example with different question inputs for the gemini model
	question := flag.String("q", "Explain how AI works in a few words", "The question to ask the model")
	deleteFile := flag.String("d", "no", "Defines if the audio file should be deleted after playing")
	flag.Parse()

	ctx := context.Background()
	geminiClient, err := client.NewGeminiClient(ctx, os.Getenv("GENAI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	generateContentService := service.NewGenerateContentService(geminiClient)

	result, err := generateContentService.Generate(ctx, "gemini-2.0-flash", *question)
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	log.Println("Gemini response: ", result.Text())

	// call the ElevenLabs API to generate speech from the text
	elevenLabsClient := client.NewElevenLabsClient(os.Getenv("ELEVEN_LABS_API_KEY"))

	err = elevenLabsClient.GenerateSpeech(result.Text())
	if err != nil {
		log.Fatalf("Error generating speech: %v", err)
	}

	playAudioError := playAudio("output.mp3")
	if playAudioError != nil {
		log.Fatalf("Error playing audio: %v", playAudioError)
	}

	log.Println("Audio finished playing")

	handleDeleteFile(*deleteFile)

	log.Println("Exiting...")

}

func handleDeleteFile(deleteFile string) {
	if deleteFile == "yes" {
		log.Println("Deleting audio file...")
		err := os.Remove("output.mp3")
		if err != nil {
			log.Fatalf("Error deleting audio file: %v", err)
		}
	} else {
		log.Println("Audio file not deleted")
	}
}

func playAudio(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening audio file: %w", err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("error decoding audio file: %w", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}
