package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/reactivex/rxgo/v2"
	"golang.org/x/exp/slices"
)

func getHLSIndexUrl(url string) (string, string, error) {
	var (
		title      string
		streamUrl  string
		streamSlug string
		hlsIndex   = ".urlset/master.m3u8"
	)

	res, err := http.Get(url)

	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)
	scan := bufio.NewScanner(reader)

	rStreamSlug, _ := regexp.Compile(",(.*?),")
	rVoeStream, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)`)
	rHTMLTag, _ := regexp.Compile(`<[^>]+>`)

	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, "\"hls\"") {
			streamSlug = rStreamSlug.FindString(line)
			streamUrl = rVoeStream.FindString(line)
		}
		if strings.Contains(line, "<title>") {
			title = rHTMLTag.ReplaceAllString(line, "")
			title = strings.TrimSpace(strings.ReplaceAll(title, "Watch", ""))
		}
	}

	return streamUrl + streamSlug + hlsIndex, title, err
}

func spawnYoutubeDL(streamUrl string, title string, ytdlp bool, cliParams []string) {
	params := append([]string{streamUrl, "--newline"}, cliParams...)

	if !slices.Contains(params, "-o") {
		params = append(params, "-o", title+".%(ext)s")
	}

	var driver string

	if ytdlp {
		driver = "yt-dlp"
	} else {
		driver = "youtube-dl"
	}

	cmd := exec.Command(driver, params...)
	r, _ := cmd.StdoutPipe()
	scan := bufio.NewScanner(r)

	ch := make(chan rxgo.Item)

	go func() {
		for scan.Scan() {
			item := rxgo.Of(scan.Text())
			ch <- item
		}

		ch <- rxgo.Of("\n")
		close(ch)
	}()

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	ob := rxgo.FromChannel(ch, rxgo.WithObservationStrategy(rxgo.Lazy)).
		Debounce(rxgo.WithDuration(125 * time.Millisecond))

	for item := range ob.Observe() {
		fmt.Print("\033[G\033[K")
		fmt.Println(item.V)
		fmt.Print("\033[A")
	}

	fmt.Println()
	cmd.Wait()
}
