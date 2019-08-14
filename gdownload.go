package main

/////////////////////////////////////////////////////////////////
// Require installing the following packages
// - youtube-dl
// - ffmpeg
/////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("")
	fmt.Println("gdownload")
	fmt.Println("  Download mp3 from currently playing youtube video on Google Chrome")
	fmt.Println("  by giovanneafonso@gmail.com")
	fmt.Println("")

	outputPath := "/home/opah/Música"
	videoID := getChromePlayingYoutubeID()
	log.Printf("videoID: %v", videoID)
	download(videoID, outputPath)
}

// Receives a url and outputs a file with the same name
func download(videoID, outputPath string) {
	filePath := fmt.Sprintf("%s/%s", outputPath, "%(title)s-%(id)s.%(ext)s")
	cmdString := fmt.Sprintf(
		"youtube-dl --extract-audio --audio-format mp3 -o \"%s\" https://youtube.com/watch?v=%s",
		filePath,
		videoID)

	cmd := exec.Command("bash", "-c", cmdString)
	out, err := cmd.Output()

	if err != nil {
		log.Fatalf("Unable to execute youtube-dl command: \n-Cmd: %s\n-Error: %v", cmdString, err)
	} else {
		log.Fatalf("Out: %s", out)
	}

	fmt.Print(string(out))
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
