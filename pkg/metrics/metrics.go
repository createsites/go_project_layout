package metrics

import (
	"time"
)

func Init() {}

func Count(status, path string) {}

func Measure(startTime time.Time, status, path string) {}
