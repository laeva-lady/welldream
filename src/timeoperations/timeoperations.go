package timeoperations

import "time"

func Add(time1, time2 string) string {
	// Parse the input times
	t1, err := time.Parse("+%Y:%m:%d", time1)
	if err != nil {
		return "00:00:00"
	}
	t2, err := time.Parse("+%Y:%m:%d", time2)
	if err != nil {
		return "00:00:00"
	}

	// Add the times
	result := t1.Add(t2.Sub(time.Time{}))

	// Format the result
	return result.Format("+%Y:%m:%d")
}
