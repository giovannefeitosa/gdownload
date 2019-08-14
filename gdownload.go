package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/rylio/ytdl"
)

func main() {
	outputPath := "/home/opah/Música"
	videoID := getChromePlayingYoutubeID()
	log.Printf("videoID: %v", videoID)
	download("https://www.youtube.com/watch?v="+videoID, outputPath)
}

// Receives a url and outputs a file with the same name
func download(videoUrl, outputPath string) {
	vid, err := ytdl.GetVideoInfo(videoUrl)
	if err != nil {
		log.Fatalf("Failed to get video info: %v", err)
		return
	}
	filePath := fmt.Sprintf("%s/%s.mp4", outputPath, escapeVideoTitle(vid.Title))
	log.Printf("filePath: %v", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
		return
	}
	defer file.Close()
	vid.Download(vid.Formats[0], file)
}

// Returns the youtube video ID currently playing on chrome
func getChromePlayingYoutubeID() string {
	cmdString := "strings ~/.config/google-chrome/Default/Current\\ Session | grep /watch?v= | tail -1"
	out, err := exec.Command("bash", "-c", cmdString).Output()
	if err != nil {
		log.Fatalf("Can't get Chrome's Current Session: %v", err)
	}

	// The url is not always well formed, but it has the ID always into querystring `v`
	urlLine := string(out)

	// Extract video Id
	s := strings.Split(urlLine, "v=")

	if len(s) < 2 {
		log.Fatalf("Can't find any video id opened on Chrome... %v", urlLine)
	}

	videoID := strings.Split(s[1], "&")[0]

	if len(videoID) < 5 {
		log.Fatalf("Error extracting videoID from url: %v", urlLine)
	}

	return videoID
}

func escapeVideoTitle(videoTitle string) string {
	videoTitle = strings.Replace(videoTitle, " ", "-", -1)
	return videoTitle
}
