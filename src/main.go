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
	if len(os.Args) < 2 {
		splash(true)
		os.Exit(0)
	}

	splash(false)

	cliUrl := os.Args[1]
	youtubeDLArgs := os.Args[2:]

	url, err := getHLSIndexUrl(cliUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[core] Fetched %s\n", url)
	spawnYoutubeDL(url, youtubeDLArgs)
}
