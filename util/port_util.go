package util

import (
	"strconv"
	"fmt"
	"strings"
)

func GetBytePeerSecond(sum float64) string {
	fmt.Println(sum)
	if sum >= 1048576 {
		return strconv.FormatFloat(sum/1048576, 'f', 2, 64) + " MB/s"
	}
	if sum >= 1024 {
		return strconv.FormatFloat(sum/1024, 'f', 2, 64) + " KB/s"
	}
	if sum > 0 {
		return strconv.FormatFloat(sum, 'f', -1, 64) + " B/s"
	} else
	{
		return ""
	}
}

func GetBytes(sumByte float64) string {
	if sumByte >= 1048576*1024 {
		return strconv.FormatFloat(sumByte/(1048576*1024), 'f', 2, 64) + " GB"
	}
	if sumByte >= 1048576 {
		return strconv.FormatFloat(sumByte/1048576, 'f', 2, 64) + " MB"
	}
	if sumByte >= 1024 {
		return strconv.FormatFloat(sumByte/1024, 'f', 2, 64) + " KB"
	}
	if sumByte >= 0 {
		return strconv.FormatFloat(sumByte, 'f', -1, 64) + " B"
	} else
	{
		return ""
	}
}

func GetPort(addr string) string {
	index := strings.Index(addr, ":")
	return addr[index+1:]
}

func CheckParam(aa ... string) bool {
	for _, value := range aa {
		if len(value) == 0 {
			return false
		}
	}
	return true
}

func GetByteFromString(value string) uint64 {
	if strings.HasSuffix(value, "GB") {
		v, _ := strconv.ParseUint(value[:len(value)-2], 10, 64)
		return v * 1048576 * 1024
	}
	if strings.HasSuffix(value, "MB") {
		v, _ := strconv.ParseUint(value[:len(value)-2], 10, 64)
		return v * 1048576
	}
	if strings.HasSuffix(value, "KB") {
		v, _ := strconv.ParseUint(value[:len(value)-2], 10, 64)
		return v * 1024
	}
	return 0
}
