package main

import (
	"assignment3/job"
	"assignment3/routers"
)

func main() {
	job.StartScheduler()
	routers.StartServer()
}
