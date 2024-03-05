package utils

import (
	"fmt"
	"time"
)

func BytesToMB(bytes int) string {
	return fmt.Sprintf("%.2fMB", float64(bytes)/1024/1024)

}
func BytesToKilobytes(bytes int) string {
	return fmt.Sprintf("%.2fKB", float64(bytes)/1024)

}
func FormatBytes(bytes int) string {
	if bytes < 1024*1024 {
		return BytesToKilobytes(bytes)
	}
	return BytesToMB(bytes)
}

func PrintElapsedTime(startTime time.Time) string {
	elapsedTime := time.Since(startTime)
	elapsedSeconds := elapsedTime.Seconds()

	// timeEndMilliseconds := float64(elapsedTime) / float64(time.Millisecond)
	return fmt.Sprintf("%.1f s", elapsedSeconds)
}
