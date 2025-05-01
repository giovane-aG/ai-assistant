# AI Assistant

A Go-based AI assistant that combines the power of Google's Gemini AI model with ElevenLabs' text-to-speech capabilities to create an interactive AI experience.

## Features

- Integration with Google's Gemini AI model for intelligent responses
- Text-to-speech conversion using ElevenLabs API
- Command-line interface for easy interaction
- Configurable question input
- Optional audio file cleanup

## Prerequisites

- Go 1.23 or later
- ElevenLabs API key
- Google Gemini API key
- macOS (for audio playback using `afplay`)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/giovane-aG/ai-assistant.git
cd ai-assistant
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the project root with your API keys:
```
GENAI_API_KEY=your_gemini_api_key
ELEVEN_LABS_API_KEY=your_elevenlabs_api_key
```

## Usage

Run the program with default settings:
```bash
go run cmd/main.go
```

### Command-line Options

- `-q`: Specify a custom question (default: "Explain how AI works in a few words")
- `-d`: Delete the audio file after playing (options: "yes" or "no", default: "no")

Example:
```bash
go run cmd/main.go -q "What is the meaning of life?" -d yes
```

## Project Structure

```
.
├── cmd/
│   └── main.go          # Main application entry point
├── internal/
│   └── client/          # API client implementations
├── go.mod               # Go module definition
├── go.sum              # Go dependencies checksum
└── .env                # Environment variables (not tracked in git)
```

## Dependencies

- github.com/joho/godotenv v1.5.1
- google.golang.org/genai v1.1.0

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Giovane Almeida 