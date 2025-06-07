package main

import (
	"bufio"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Hey, I started!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s\n", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received raw content: %s", string(content))
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Couldn't parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("didn't receive a good file")
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
