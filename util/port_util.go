package util

import "strconv"

func GetSpeed(sum int64) string {
	if sum > 1048576 {
		return strconv.FormatInt(sum/1048576, 10) + " MB/s"
	}
	if sum > 1024 {
		return strconv.FormatInt(sum/1024, 10) + " KB/s"
	}
	if sum > 0 {
		return strconv.FormatInt(sum, 10) + " B/s"
	} else
	{
		return ""
	}
}

func GetBytePeerSecond(sumByte int64) string {
	if sumByte > 1048576 {
		return strconv.FormatInt(sumByte/1048576, 10) + " MB"
	}
	if sumByte > 1024 {
		return strconv.FormatInt(sumByte/1024, 10) + " KB"
	}
	if sumByte > 0 {
		return strconv.FormatInt(sumByte, 10) + " B"
	} else
	{
		return ""
	}
}
