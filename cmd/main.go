package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	client "github.com/giovane-aG/ai-assistant/internal/client"
	"github.com/giovane-aG/ai-assistant/internal/service"
	"github.com/joho/godotenv"
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

	// Create a channel to handle signal interrupt
	signalChan := setupSignalHandler()

	err = playAudio(signalChan)
	if err != nil {
		log.Fatalf("Error playing audio: %v", err)
	}

	close(signalChan)

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

func handleInterrupt(cmd *exec.Cmd, signalChan chan os.Signal) {
	<-signalChan
	log.Println("Received interrupt signal, cleaning up...")
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatalf("Error killing process: %v", err)
	}
	os.Exit(0)
}

func setupSignalHandler() chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	return signalChan
}

func playAudio(signalChan chan os.Signal) error {
	cmd := exec.Command("afplay", "output.mp3")
	go handleInterrupt(cmd, signalChan)

	err := cmd.Start()
	if err != nil {
		return err
	}

	log.Println("Playing audio...")
	return cmd.Wait()
}
