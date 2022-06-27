package main

import "fmt"

func splash(more bool) {
	fmt.Println("___    __________________   ______________ \n__ |  / /_  __ \\__  ____/   ___  __ \\__  / \n__ | / /_  / / /_  __/________  / / /_  /  \n__ |/ / / /_/ /_  /___/_____/  /_/ /_  /___\n_____/  \\____/ /_____/      /_____/ /_____/.go")
	fmt.Println("\nyoutube-dl / yt-dlp wrapper for voe.net HLS streams")
	if more {
		fmt.Println("---------------------------------------------------------------------")
		fmt.Println("Usage: voe-dl <voe.sx stream url> [youtube-dl args]")
	}
	fmt.Println()
}
