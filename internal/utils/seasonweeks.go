// Package utils provides season week configuration data.
// This file contains date ranges for each week of the football season used
// for scheduling and data extraction operations.
package utils

type weekRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

var SEASON_WEEKS = []weekRange{
	{Start: "2026-08-25", End: "2026-08-31"},
	{Start: "2026-09-01", End: "2026-09-07"},
	{Start: "2026-09-08", End: "2026-09-14"},
	{Start: "2026-09-15", End: "2026-09-21"},
	{Start: "2026-09-22", End: "2026-09-28"},
	{Start: "2026-09-29", End: "2026-10-05"},
	{Start: "2026-10-06", End: "2026-10-12"},
	{Start: "2026-10-13", End: "2026-10-19"},
	{Start: "2026-10-20", End: "2026-10-26"},
	{Start: "2026-10-27", End: "2026-11-02"},
	{Start: "2026-11-03", End: "2026-11-09"},
	{Start: "2026-11-10", End: "2026-11-16"},
	{Start: "2026-11-17", End: "2026-11-23"},
	{Start: "2026-11-24", End: "2026-11-30"},
	{Start: "2026-12-01", End: "2026-12-07"},
	{Start: "2026-12-08", End: "2026-12-14"},
	{Start: "2026-12-15", End: "2026-12-21"},
	{Start: "2026-12-22", End: "2026-12-28"},
	{Start: "2026-12-29", End: "2027-01-04"},
	{Start: "2027-01-05", End: "2027-01-11"},
	{Start: "2027-01-12", End: "2027-01-18"},
	{Start: "2027-01-19", End: "2027-01-25"},
}
