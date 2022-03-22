package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoEntry struct {
	ID               string `json:"id"`
	TargetResolution string `json:"resolution"`
	SourceLocation   string `json:"source_file"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
	Progress         int    `json:"transcode_progress"`
}

func GetVideosFromDB() []VideoEntry {
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
	return found_videos
}

func GetVideoFromDB(id_of_video int) VideoEntry {
	var found_video VideoEntry
	db := GimmeTheDB()
	err := db.QueryRow("SELECT * FROM video where video_id = ?", id_of_video).Scan(
		&found_video.ID,
		&found_video.TargetResolution,
		&found_video.SourceLocation,
		&found_video.Progress,
		&found_video.Created,
		&found_video.Updated)
	if err != nil {
		panic(1)
	}
	return found_video
}

func GetVideoByID(c *gin.Context) {
	id := c.Param("id")
	id_as_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	video := GetVideoFromDB(id_as_int)
	c.IndentedJSON(http.StatusOK, video)
}

func GetAllVideos(c *gin.Context) {
	found_videos := GetVideosFromDB()
	c.IndentedJSON(http.StatusOK, found_videos)
}
