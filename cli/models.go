package cli

import "github.com/robfig/cron/v3"

func Forever(job cron.Job) {
	go func() {
		for {
			job.Run()
		}
	}()
}

func Once(job cron.Job, sync bool) {
	if sync {
		job.Run()
	} else {
		go job.Run()
	}
}