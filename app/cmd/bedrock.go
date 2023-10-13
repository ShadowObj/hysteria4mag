package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	fakeLogList = [25]string{
		"[%s:%-3d INFO] Starting Server",
		"[%s:%-3d INFO] Version: 1.20.30.02",
		"[%s:%-3d INFO] Session ID: 6382494d-4f3b-42cb-9f17-0069d6f8e3ed",
		"[%s:%-3d INFO] Build ID: 17844395",
		"[%s:%-3d INFO] Branch: r/20_u3",
		"[%s:%-3d INFO] Commit ID: b228d54f3db2d513a6a0a6527d7eb16b2f883916",
		"[%s:%-3d INFO] Configuration: Publish",
		"[%s:%-3d INFO] Level Name: Bedrock level",
		"[%s:%-3d INFO] No CDN config file found for dedicated server",
		"[%s:%-3d INFO] Game mode: 0 Survival",
		"[%s:%-3d INFO] Difficulty: 1 EASY",
		"[%s:%-3d INFO] Content logging to console is enabled.",
		"[%s:%-3d INFO] Opening level 'worlds/Bedrock level/db'",
		strings.ReplaceAll("[%s:%-3d INFO] IPv4 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
		"[%s:%-3d INFO] IPv6 not supported",
		"[%s:%-3d INFO] Server started.",
		"[%s:%-3d INFO] ================ TELEMETRY MESSAGE ===================",
		"[%s:%-3d INFO] Server Telemetry is currently not enabled.",
		"[%s:%-3d INFO] Enabling this telemetry helps us improve the game.",
		"[%s:%-3d INFO]",
		"[%s:%-3d INFO] To enable this feature, add the line 'emit-server-telemetry=true'",
		"[%s:%-3d INFO] to the server.properties file in the handheld/src-server directory",
		"[%s:%-3d INFO] ======================================================",
		"[%s:%-3d INFO] IPv4 supported, port: 19132 used for LAN discovery",
		"[%s:%-3d INFO] IPv6 not supported ",
	}
	fakeInput string
)

func fakeBedrockServer() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	go http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	fmt.Println("NO LOG FILE! - setting up server logging...")
	for n := 0; n < len(fakeLogList); n++ {
		fmt.Printf(fakeLogList[n]+"\n", time.Now().Format("2006-01-02 15:03:04"), time.Now().Nanosecond())
	}
	for {
		fmt.Scanln(&fakeInput)
		if fakeInput == "stop" {
			os.Exit(0)
		} else {
			fmt.Printf(
				"[%s:%-3d ERROR] Unknown command: %s. Please check that the command exists and that you have permission to use it.",
				time.Now().Format("2006-01-02 15:03:04"),
				time.Now().Nanosecond(),
				fakeInput,
			)
		}
	}
}
