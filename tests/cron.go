package object

import (
	"fmt"
	"strings"
	"time"

	"github.com/mileusna/crontab"
)

var ctab *crontab.Crontab

func init() {
	ctab = crontab.New()
}

func schedulePost(postId string) {
	post := GetPost(postId)
	isBumped := post.bumpPost()
	if isBumped {
		fmt.Printf("Bump post: %s\n", post.Id)
	}
}

func parseDumpTime(bumpTime string) (string, string) {
	tokens := strings.Split(bumpTime, ":")
	return tokens[0], tokens[1]
}

func refreshCronTasks() bool {
	ctab.Clear()

	jobs := GetJobs()
	for _, job := range jobs {
		if job.State != "active" || job.BumpTime == "" {
			continue
		}

		hours, minutes := parseDumpTime(job.BumpTime)

		posts := GetPosts(job.Id)
		for _, post := range posts {
			if post.State != "sent" || post.Url == "" {
				continue
			}

			schedule := fmt.Sprintf("%s %s * * *", minutes, hours)
			//schedule := "* * * * *"
			err := ctab.AddJob(schedule, schedulePost, post.Id)
			if err != nil {
				panic(err)
			}
		}
	}

	return true
}

func timerRoutine() {
	for range time.Tick(time.Second * 3600) {
		//println("timer")
		refreshCronTasks()
	}
}

func InitTimer() {
	refreshCronTasks()

	go timerRoutine()
}
