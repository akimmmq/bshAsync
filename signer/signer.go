package main

import (
	"fmt"
	"strconv"
)

// Write code here

//ExecutePipeline

func ExecutePipeline() {
	SingleHash(2)
	MultiHash(2)
	CombineResults(2)
}

//Hash calculating

func SingleHash(data int) string {
	stData := strconv.Itoa(data)
	for i := 0; i < data; i++ {
		crs32 := DataSignerCrc32(stData)
		crs32md5 := DataSignerCrc32(DataSignerMd5(stData))
		fmt.Println(stData + "SingleHash data: " + stData)
		fmt.Println(stData + "SingleHash md5(data): " + DataSignerMd5(stData))
		fmt.Println(stData + "SingleHash crc32(md5(data)): " + DataSignerCrc32(DataSignerMd5(stData)))
		fmt.Println(stData + "SingleHash crc32(data): " + DataSignerCrc32(stData))
		fmt.Println(stData + "SingleHash result: " + (crs32 + crs32md5))
		return crs32 + crs32md5
	}
	return ""
}

func MultiHash(data int) string {
	step1 := SingleHash(data)
	var strI string
	for i := 0; i < data; i++ {
		var result string
		for i := 0; i < 6; i++ {
			strI = strconv.Itoa(i)
			result += DataSignerCrc32(strI + step1)
			fmt.Println(step1 + " MultiHash: crc32(th+step1)) " +
				strI + " " + DataSignerCrc32(strI+step1))
		}
		fmt.Println(step1 + " MultiHash result :  " + result)
		return result
	}
	return ""
}

func CombineResults(data int) {
	var tmp []string
	var result string
	tmp = append(tmp, MultiHash(data))
	for i := range tmp {
		result += strconv.Itoa(i) + "_"
	}
	s1 := result
	if last := len(s1) - 1; last >= 0 && s1[last] == '_' {
		s1 = s1[:last]
	}
	fmt.Println("s1:", s1)
	fmt.Println("CombineResults: " + s1)

}
