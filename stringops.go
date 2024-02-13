package datautils

import "time"

// string utilities
func TruncateTextWithEllipsis(text string, max_len int) string {
	if len(text) > max_len {
		return text[:max_len] + "..."
	}
	return text
}

func DateToString(time_val float64) string {
	timeValue := time.Unix(int64(time_val), 0)
	return timeValue.Format("2006-01-02")
}
