package main

import (
	"fmt"
)

type VideoEntry struct {
	ID               string `json:"id"`
	TargetResolution string `json:"resolution"`
	SourceLocation   string `json:"source_file"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
	Progress         int    `json:"transcode_progress"`
}

func SeedData() {
	_, err := QueryDB(
		`INSERT INTO video (
			target_resolution, 
			source_location,
			transcode_progress
		) VALUES ('4k', '"/mnt/videos/input/video.avi"', '0')`)
	if err != nil {
		panic(1)
	}
	vids := GetVideosFromDB()
	fmt.Println(vids)
}
