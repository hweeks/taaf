package main

import (
	"fmt"
)

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
