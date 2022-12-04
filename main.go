// Remove Symbol and Debug info at compile
// go build -ldflags "-s -w"

// Run on linux and mac os x machines
// sudo ./get-thumbnail-Youtube

package main

import (
	"errors"
	"io"
	"log"
	"os"
	"regexp"
)

func GetYTThumbnailFromId(id string) (string, error) {
	t := NewYTThumbnail()
	t.VideoID = id
	t.link = "https://www.youtube.com/watch?v=" + t.VideoID
	err := createFolder(t.thumbnailsDir)
	if err != nil {
		return "", err
	}

	// check if file already exists
	filename := t.setYTThumbnailName()
	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// file does not exist
	} else {
		// file exists
		return filename, nil
	}

	// Check if video id is valid by regex
	if !IsValidYTVideoID(t.VideoID) {
		return "", errors.New("Invalid video id")
	}
	// // Walk walks thumbnails dir and save file names
	// err = filepath.Walk(t.thumbnailsDir, t.walkFunc)
	// if err != nil {
	// 	return "", err
	// }

	readyFile, errCreate := createFile(filename)
	if errCreate != nil {
		return "", errCreate
	}

	errWrite := writeFile(readyFile, t.getURLResponse())
	if errWrite != nil {
		return "", errWrite
	}
	return t.setYTThumbnailName(), errWrite
}

func main() {

	// take first argument from command line and save it to variable
	VideoID := os.Args[1]
	// check if video id is valid by regex
	if !IsValidYTVideoID(VideoID) {
		log.Println("Invalid video id")
		os.Exit(1)
	}

	filename, err := GetYTThumbnailFromId(VideoID)

	if err != nil {
		os.Exit(0)
	}

	log.Println(filename)

}

func GetYTThumbnailImageFromId(id string) ([]byte, error) {

	t := NewYTThumbnail()

	t.VideoID = id
	// set video url from video id
	t.link = "https://www.youtube.com/watch?v=" + t.VideoID

	// Check if video id is valid by regex
	if !IsValidYTVideoID(t.VideoID) {
		return []byte{}, errors.New("Invalid video id")
	}

	err := createFolder(t.thumbnailsDir)
	if err != nil {
		return nil, err
	}

	readyFile, errCreate := createFile(t.setYTThumbnailName())

	if errCreate != nil {
		return nil, errCreate
	}

	errWrite := writeFile(readyFile, t.getURLResponse())

	if errWrite != nil {
		return nil, errWrite
	}
	// convert file to []byte
	return io.ReadAll(readyFile)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func IsValidYTVideoID(id string) bool {
	// regex for video id
	regExVideoID := `^[a-zA-Z0-9_-]{11}$`
	// check if video id is valid
	return regexp.MustCompile(regExVideoID).MatchString(id)
}
