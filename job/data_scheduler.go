package job

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("Starting scheduler...")
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(ReadAndUpdateData)
	<-s.Start()
}

func StartScheduler() {
	go task()
}
