package main

import "time"

func nanoseconds(n int) time.Duration {
	// 1
	return time.Duration(n) * time.Nanosecond
}

func microseconds(n int) time.Duration {
	// 1000 * Nanosecond
	return time.Duration(n) * time.Microsecond
}

func milliseconds(n int) time.Duration {
	// 1000 * Microsecond
	return time.Duration(n) * time.Millisecond
}

func seconds(n int) time.Duration {
	// 1000 * Millisecond
	return time.Duration(n) * time.Second
}

func minutes(n int) time.Duration {
	// 60 * Second
	return time.Duration(n) * time.Minute
}

func hours(n int) time.Duration {
	// 60 * Minute
	return time.Duration(n) * time.Hour
}
