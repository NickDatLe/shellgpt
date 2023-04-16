# ShellGPT

This is command line tool written in Go that takes in natural language commands and translates it to command line instructions for the host operating system. The translated command is printed to your shell and also copied to your clipboard.

## Examples:

with go runtime 
```bash
go run shellgpt.go "find any file in the /var/logs folder, and output the first 10 lines of each file"
```

output:
```bash
(Copied to clipboard)
find /var/logs -type f -exec head -n 10 {} \;
```

## Requirements

- You'll need Go on your system, go to: https://go.dev/doc/install
- You'll need an OpenAI API key

## Installation

Copy and paste your OpenAI API key into the "key.txt" file.

Build the app:
```bash
go build shellgpt.go
```

Now you can run shellgpt without the go runtime:

```bash
shellgpt "list all of the current files in the folder temp and remove any that starts with 'tmp'"
(Copied to clipboard)
ls temp | grep '^tmp' | xargs -I {} rm temp/{}
```

