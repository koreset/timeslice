package services

import "time"

func stringDate(t time.Time) string {
	return t.Format("2006-01-02")
}

