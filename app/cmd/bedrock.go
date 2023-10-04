package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	fakeLogList = [10]string{
		"[%s:%-3d INFO] Starting Server",
		"[%s:%-3d INFO] Version 1.17.0.03",
		"[%s:%-3d INFO] Session ID ccp114514",
		"[%s:%-3d INFO] Level Name: Bedrock level",
		"[%s:%-3d INFO] Game mode: 0 Survival",
		"[%s:%-3d INFO] Difficulty: 1 EASY",
		"[%s:%-3d INFO] opening worlds/Bedrock level/db",
		strings.ReplaceAll("[%s:%-3d INFO] IPv4 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
		strings.ReplaceAll("[%s:%-3d INFO] IPv6 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
		"[%s:%-3d INFO] Server started.",
	}
	fakeInput string
)

func fakeBedrockServer() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	go http.ListenAndServe(os.Getenv("SERVER_PORT"), nil)
	for n := 0; n < len(fakeLogList); n++ {
		fmt.Printf(fakeLogList[n]+"\n", time.Now().Format("2006-01-02 15:03:04"), time.Now().Nanosecond())
	}
	for {
		fmt.Scanln(&fakeInput)
		if fakeInput == "stop" {
			os.Exit(0)
		} else {
			fmt.Println("Invaild Command.")
		}
	}
}
