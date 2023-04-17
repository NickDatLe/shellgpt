package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/atotto/clipboard"
	openai "github.com/sashabaranov/go-openai"
)

func getKey(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(
			`ERROR: Please make sure there is a file named key.txt in the same directory as this program and that your OpenAI API key is in it
		Error details: %v`, err)
	}
	defer file.Close()

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	return content
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Please provide your natural language command and I will convert it to bash script.")
		fmt.Println("Example: bashgpt \"list all files in current directory\"")
		return
	}

	prompt := args[1]

	myOS := runtime.GOOS

	systemPrompt := fmt.Sprintf(`You are an expert command-line script writer for OS %s.
	You take human commands for processing files and folders in natural language and convert it to
	bash single liners to run. Your output is only the command and if you don't know then write "NULL"

	%s
	`, myOS, prompt)

	client := openai.NewClient(getKey("key.txt"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	response := resp.Choices[0].Message.Content
	clipboard.WriteAll(response)
	fmt.Println("(Copied to clipboard)")
	fmt.Println(response)
}
