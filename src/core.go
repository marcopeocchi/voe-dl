package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func getHLSIndexUrl(url string) (string, error) {
	var (
		streamUrl  string
		streamSlug string
		hlsIndex   = ".urlset/master.m3u8"
	)

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	flattened := strings.ReplaceAll(string(body), "\n", "")
	rSourcesObject, _ := regexp.Compile("{(.*?)}")
	rStreamSlug, _ := regexp.Compile(",(.*?),")
	rVoeStream, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)`)

	for _, line := range rSourcesObject.FindAllString(flattened, -1) {
		if strings.Contains(line, "\"hls\"") {
			streamSlug = rStreamSlug.FindString(line)
			streamUrl = rVoeStream.FindString(line)
		}
	}

	return streamUrl + streamSlug + hlsIndex, err
}

func spawnYoutubeDL(streamUrl string, cliParams []string) {
	params := append([]string{streamUrl, "--newline"}, cliParams...)

	cmd := exec.Command("yt-dlp", params...)
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

	cmd.Wait()
}
