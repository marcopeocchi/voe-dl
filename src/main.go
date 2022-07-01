package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	cliUrl := flag.String("u", "", "Video url")
	driver := flag.String("d", "yt-dlp", "Which driver to use (yt-dlp or youtube-dl)")
	youtubeDLArgs := flag.String("a", "", "youtube-dl/yt-dlp additional arguments")

	flag.Parse()

	youtubeDLArgsList := strings.Split(*youtubeDLArgs, " ")

	if *cliUrl == "" {
		splash(true)
		os.Exit(-1)
	}

	url, title, err := getHLSIndexUrl(*cliUrl)
	if err != nil {
		panic(err)
	}

	if *driver == "yt-dlp" {
		spawnYoutubeDL(url, title, *driver, youtubeDLArgsList)
	} else if *driver == "youtube-dl" {
		spawnYoutubeDL(url, title, *driver, youtubeDLArgsList)
	} else {
		fmt.Printf("\033[1;31m%s\033[0m\n", "You must specify which driver to use")
	}
}
