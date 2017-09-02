package util

import "strconv"

func GetBytePeerSecond(sum int64) string {

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

func GetBytes(sumByte float64) string {
	if sumByte > 1048576*1024 {
		return strconv.FormatFloat(sumByte/(1048576*1024), byte('f'),-1,64) + " GB"
	}
	if sumByte > 1048576 {
		return strconv.FormatFloat(sumByte/1048576, byte('f'),2,64) + " MB"
	}
	if sumByte > 1024 {
		return strconv.FormatFloat(sumByte/1024, byte('f'),-1,64) + " KB"
	}
	if sumByte > 0 {
		return strconv.FormatFloat(sumByte, byte('f'),-1,64) + " B"
	} else
	{
		return ""
	}
}
