package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVideos(c *gin.Context) {
	var found_videos []VideoEntry
	result, err := QueryDB("SELECT * from video")
	if err != nil {
		panic(1)
	}
	defer result.Close()
	for result.Next() {
		var inner_video VideoEntry
		if err := result.Scan(
			&inner_video.ID,
			&inner_video.TargetResolution,
			&inner_video.SourceLocation,
			&inner_video.Progress,
			&inner_video.Created,
			&inner_video.Updated,
		); err != nil {
			fmt.Println(err)
			panic(1)
		}
		found_videos = append(found_videos, inner_video)
	}
	c.IndentedJSON(http.StatusOK, found_videos)
}
