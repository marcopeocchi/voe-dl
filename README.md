# voe-dl

Download HLS streams from voe.sx using [youtube-dl](https://github.com/ytdl-org/youtube-dl) or [yt-dlp](https://github.com/yt-dlp/yt-dlp) as backends.

## Prerequisites

You must have youtube-dl or ytp-dlp in any folder of your **PATH** env variable.

```
___    __________________   ______________ 
__ |  / /_  __ \__  ____/   ___  __ \__  / 
__ | / /_  / / /_  __/________  / / /_  /  
__ |/ / / /_/ /_  /___/_____/  /_/ /_  /___
_____/  \____/ /_____/      /_____/ /_____/.go

youtube-dl / yt-dlp wrapper for voe.sx HLS streams
---------------------------------------------------------------------
Usage: voe-dl <voe.sx stream url> <driver> [youtube-dl args]
Drivers: -d -> youtube-dl
         -p -> yt-dlp
```

## Build from source
*Make* sure you have **go 1.18** and **make** installed.
```
$ make build
```
## Install
```
# make install
```

## youtube-dl / yt-dlp params/args

You can pass every compatible params after the driver flag.  
https://github.com/ytdl-org/youtube-dl#options for evidence.
