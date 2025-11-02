package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/StefanNSTanev/stocks_CLI/models"
)

// printWeeklyTable formats and prints up to 5 weekly entries in a table.
func printWeeklyTable(weeklyData models.CompanyWeeklyPriceData, symbol string) {
	fmt.Printf("Weekly Data for %s:\n\n", symbol)

	if len(weeklyData.WeeklyTimeSeries) == 0 {
		fmt.Println("no weekly data")
		return
	}

	// Collect and sort keys so output is deterministic (most recent first).
	keys := make([]string, 0, len(weeklyData.WeeklyTimeSeries))
	for k := range weeklyData.WeeklyTimeSeries {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	limit := 5
	if len(keys) < limit {
		limit = len(keys)
	}

	// Print table header
	fmt.Printf("%-12s %10s %10s %10s %10s %12s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Printf("%-12s %10s %10s %10s %10s %12s\n", strings.Repeat("-", 12), strings.Repeat("-", 10), strings.Repeat("-", 10), strings.Repeat("-", 10), strings.Repeat("-", 10), strings.Repeat("-", 12))

	// Print up to first 5 items (most recent after sorting)
	for i := 0; i < limit; i++ {
		k := keys[i]
		v := weeklyData.WeeklyTimeSeries[k]
		fmt.Printf("%-12s %10.2f %10.2f %10.2f %10.2f %12d\n",
			k, v.Open, v.High, v.Low, v.Close, v.Volume)
	}
}

func printDailyTable(weeklyData models.CompanyDailyPriceData, symbol string) {
	fmt.Printf("Daily Data for %s:\n\n", symbol)

	if len(weeklyData.WeeklyTimeSeries) == 0 {
		fmt.Println("no daily data")
		return
	}

	// Collect and sort keys so output is deterministic (most recent first).
	keys := make([]string, 0, len(weeklyData.WeeklyTimeSeries))
	for k := range weeklyData.WeeklyTimeSeries {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	limit := 5
	if len(keys) < limit {
		limit = len(keys)
	}

	// ANSI color codes
	green := "\x1b[32m"
	red := "\x1b[31m"
	reset := "\x1b[0m"

	// Column widths (must match header formatting)
	dateW := 12
	floatW := 10
	volW := 12

	// Print table header
	fmt.Printf("%-12s %10s %10s %10s %10s %12s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Printf("%-12s %10s %10s %10s %10s %12s\n",
		strings.Repeat("-", dateW),
		strings.Repeat("-", floatW),
		strings.Repeat("-", floatW),
		strings.Repeat("-", floatW),
		strings.Repeat("-", floatW),
		strings.Repeat("-", volW),
	)

	// helpers that keep padding outside of color codes so alignment remains correct
	colorFieldFloat := func(cur, prev float64, hasPrev bool) string {
		s := fmt.Sprintf("%.2f", cur)
		padding := floatW - len(s)
		if padding < 0 {
			padding = 0
		}
		pad := strings.Repeat(" ", padding)
		if !hasPrev {
			return pad + s
		}
		switch {
		case cur > prev:
			return pad + green + s + reset
		case cur < prev:
			return pad + red + s + reset
		default:
			return pad + s
		}
	}

	colorFieldInt := func(cur, prev int64, hasPrev bool) string {
		s := fmt.Sprintf("%d", cur)
		padding := volW - len(s)
		if padding < 0 {
			padding = 0
		}
		pad := strings.Repeat(" ", padding)
		if !hasPrev {
			return pad + s
		}
		switch {
		case cur > prev:
			return pad + green + s + reset
		case cur < prev:
			return pad + red + s + reset
		default:
			return pad + s
		}
	}

	// Print up to first `limit` items (most recent after sorting).
	// Compare each entry with the previous (older) entry at i+1.
	for i := 0; i < limit; i++ {
		k := keys[i]
		v := weeklyData.WeeklyTimeSeries[k]

		var prev models.PriceData
		hasPrev := false
		if i+1 < len(keys) {
			prev = weeklyData.WeeklyTimeSeries[keys[i+1]]
			hasPrev = true
		}

		openStr := colorFieldFloat(v.Open, prev.Open, hasPrev)
		highStr := colorFieldFloat(v.High, prev.High, hasPrev)
		lowStr := colorFieldFloat(v.Low, prev.Low, hasPrev)
		closeStr := colorFieldFloat(v.Close, prev.Close, hasPrev)
		volStr := colorFieldInt(v.Volume, prev.Volume, hasPrev)

		// print with fixed date column and pre-padded value strings so columns align
		fmt.Printf("%-12s %s %s %s %s %s\n",
			k, openStr, highStr, lowStr, closeStr, volStr)
	}
}
