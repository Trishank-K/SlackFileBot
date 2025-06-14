# Slack File Bot

This project is a Go application that uploads files to a specified Slack channel using a bot token.

## Prerequisites

- Go (version 1.24.2 or higher)
- A Slack Bot Token with the necessary permissions (files:write)
- The ID of the Slack channel where files will be uploaded

## Installation

1. **Clone the repository (if applicable):**
   ```bash
   git clone  https://github.com/Trishank-K/SlackFileBot
   cd SlackFileBot
   ```
2. **Install dependencies:**
   The project uses Go modules. Dependencies are listed in the `go.mod` file and will be downloaded automatically when you build or run the project.
   ```bash
   go mod tidy
   ```
3. **Set up environment variables:**
   Create a `.env` file in the root of the project directory with the following content:
   ```env
   SLACK_BOT_TOKEN="YOUR_SLACK_BOT_TOKEN"
   CHANNEL_ID="YOUR_SLACK_CHANNEL_ID"
   ```

   Replace `YOUR_SLACK_BOT_TOKEN` with your actual Slack bot token and `YOUR_SLACK_CHANNEL_ID` with the target channel ID.

## Usage

1. **Prepare files for upload:**
   Place the files you want to upload in the project directory. The current implementation is hardcoded to upload a file named `test.txt`. You can modify the `fileArr` variable in `main.go` to include other files.

   ```go
   // main.go
   // ...
   fileArr := []string{"test.txt"} // Modify this array to include your desired files
   // ...
   ```
2. **Run the application:**

   ```bash
   go run main.go
   ```

   The application will then upload the specified files to the configured Slack channel. You should see output in the console indicating the name of the uploaded file.

## Dependencies

This project relies on the following Go packages:

- `github.com/joho/godotenv` (v1.5.1): For loading environment variables from a `.env` file.
- `github.com/slack-go/slack` (v0.17.1): A Go client library for the Slack API.
- `github.com/gorilla/websocket` (v1.5.3): (Indirect dependency via `github.com/slack-go/slack`)

These dependencies are managed using Go modules and are defined in the `go.mod` file.
