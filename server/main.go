package main

import (
	serve "Sample_Douyin_Server/server/serve"
)

func main() {
	server := serve.CreateServer()
	server.Serve()
}
