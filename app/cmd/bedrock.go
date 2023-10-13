package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	fakeLogList = [26]string{
		"NO LOG FILE! - setting up server logging..",
		"[2023-10-13 00:17:21:097 INFO] Starting Server",
		"[2023-10-13 00:17:21:097 INFO] Version: 1.20.30.02",
		"[2023-10-13 00:17:21:097 INFO] Session ID: 6382494d-4f3b-42cb-9f17-0069d6f8e3ed",
		"[2023-10-13 00:17:21:097 INFO] Build ID: 17844395",
		"[2023-10-13 00:17:21:097 INFO] Branch: r/20_u3",
		"[2023-10-13 00:17:21:097 INFO] Commit ID: b228d54f3db2d513a6a0a6527d7eb16b2f883916",
		"[2023-10-13 00:17:21:097 INFO] Configuration: Publish",
		"[2023-10-13 00:17:21:098 INFO] Level Name: Bedrock level",
		"[2023-10-13 00:17:21:101 INFO] No CDN config file found for dedicated server",
		"[2023-10-13 00:17:21:101 INFO] Game mode: 0 Survival",
		"[2023-10-13 00:17:21:101 INFO] Difficulty: 1 EASY",
		"[2023-10-13 00:17:21:103 INFO] Content logging to console is enabled.",
		"[2023-10-13 00:17:21:666 INFO] Opening level 'worlds/Bedrock level/db'",
		strings.ReplaceAll("[%s:%-3d INFO] IPv4 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
		"[2023-10-13 00:17:23:932 INFO] IPv6 not supported",
		"[2023-10-13 00:17:23:942 INFO] Server started.",
		"[2023-10-13 00:17:23:942 INFO] ================ TELEMETRY MESSAGE ===================",
		"[2023-10-13 00:17:23:942 INFO] Server Telemetry is currently not enabled.",
		"[2023-10-13 00:17:23:942 INFO] Enabling this telemetry helps us improve the game.",
		"[2023-10-13 00:17:23:942 INFO]",
		"[2023-10-13 00:17:23:942 INFO] To enable this feature, add the line 'emit-server-telemetry=true'",
		"[2023-10-13 00:17:23:942 INFO] to the server.properties file in the handheld/src-server directory",
		"[2023-10-13 00:17:23:942 INFO] ======================================================",
		"[2023-10-13 00:17:23:972 INFO] IPv4 supported, port: 19132 used for LAN discovery",
		"[2023-10-13 00:17:23:972 INFO] IPv6 not supported ",
	}
	//	fakeLogList = [11]string{
	//		"NO LOG FILE! - setting up server logging...",
	//		"[%s:%-3d INFO] Starting Server",
	//		"[%s:%-3d INFO] Version 1.20.30.02",
	//		"[%s:%-3d INFO] Session ID ccp114514",
	//		"[%s:%-3d INFO] Level Name: Bedrock level",
	//		"[%s:%-3d INFO] Game mode: 0 Survival",
	//		"[%s:%-3d INFO] Difficulty: 1 EASY",
	//		"[%s:%-3d INFO] opening worlds/Bedrock level/db",
	//		strings.ReplaceAll("[%s:%-3d INFO] IPv4 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
	//		strings.ReplaceAll("[%s:%-3d INFO] IPv6 supported, port: ${SERVER_PORT}, used for game play", "${SERVER_PORT}", os.Getenv("SERVER_PORT")),
	//		"[%s:%-3d INFO] Server started.",
	//	}
	fakeInput string
)

func fakeBedrockServer() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	go http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	for n := 0; n < len(fakeLogList); n++ {
		fmt.Printf(fakeLogList[n]+"\n", time.Now().Format("2006-01-02 15:03:04"), time.Now().Nanosecond())
	}
	//	for {
	//		fmt.Scanln(&fakeInput)
	//		if fakeInput == "stop" {
	//			os.Exit(0)
	//		} else {
	//			fmt.Printf(
	//				"[%s:%-3d ERROR] Unknown command: %s. Please check that the command exists and that you have permission to use it.",
	//				time.Now().Format("2006-01-02 15:03:04"),
	//				time.Now().Nanosecond(),
	//				fakeInput,
	//			)
	//		}
	//	}
}
