package main

import (
	"fmt"
	"os"
)

/***
 *    ___    __________________   ______________
 *    __ |  / /_  __ \__  ____/   ___  __ \__  /
 *    __ | / /_  / / /_  __/________  / / /_  /
 *    __ |/ / / /_/ /_  /___/_____/  /_/ /_  /___
 *    _____/  \____/ /_____/      /_____/ /_____/
 *
 * 	  Golang wrapper around youtube-dl / yt-dlp for:
 * 	  voe.sx HLS streams
 */

func main() {
	if len(os.Args) < 3 {
		splash(true)
		os.Exit(0)
	}

	splash(false)

	cliUrl := os.Args[1]
	driver := os.Args[2]

	youtubeDLArgs := []string{}

	if len(os.Args) > 3 {
		youtubeDLArgs = os.Args[3:]
	}

	url, err := getHLSIndexUrl(cliUrl)
	if err != nil {
		panic(err)
	}

	if driver == "-d" {
		spawnYoutubeDL(url, false, youtubeDLArgs)
	} else if driver == "-p" {
		spawnYoutubeDL(url, true, youtubeDLArgs)
	} else if driver == "-s" {
		fmt.Println(url)
		return
	} else {
		fmt.Printf("\033[1;31m%s\033[0m\n", "You must specify which driver to use, -p for yt-dlp or -d for youtube-dl")
	}

}
