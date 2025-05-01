package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	client "github.com/giovane-aG/ai-assistant/internal/client"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup flags so that we can run this example with different question inputs for the gemini model
	question := flag.String("q", "Explain how AI works in a few words", "The question to ask the model")
	deleteFile := flag.String("d", "no", "Defines if the audio file should be deleted after playing")
	flag.Parse()

	ctx := context.Background()

	genAIClient, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GENAI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatal(err)
	}

	result, err := genAIClient.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(*question),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gemini response: ", result.Text())

	// call the ElevenLabs API to generate speech from the text
	elevenLabsClient := client.NewElevenLabsClient(os.Getenv("ELEVEN_LABS_API_KEY"))

	err = elevenLabsClient.GenerateSpeech(result.Text())
	if err != nil {
		log.Fatal(err)
	}

	// play the audio file
	cmd := exec.Command("afplay", "output.mp3")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Playing audio...")
	// wait for the audio to finish playing
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Audio finished playing")

	// clean up the audio file
	if *deleteFile == "yes" {
		fmt.Println("Deleting audio file...")
		err = os.Remove("output.mp3")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Audio file not deleted")
	}

	// wait for a while before exiting
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
}
