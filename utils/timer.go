package utils

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time, functionName string) string {
	elapsed := time.Since(start)
	totalTime := functionName + " took " + elapsed.String()

	fmt.Println(totalTime)

	return totalTime
}
