package main

import "sort"

/*
 * Complete the 'minDiff' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func minDiff(arr []int32) int32 {
	// Write your code here
	a := convertToIntArr(arr)
	sort.Ints(a)
	a32 := convertToInt32Arr(a)
	return sum(a32)
}

func sum(arr []int32) int32 {
	var sum int32
	for i := 0; i+1 < len(arr); i++ {
		sum += diff(arr[i], arr[i+1])
	}
	return sum
}

func diff(a, b int32) int32 {
	if a < b {
		return b - a
	}
	return a - b
}

func convertToIntArr(arr []int32) []int {
	i := make([]int, 0)
	for _, v := range arr {
		i = append(i, int(v))
	}
	return i
}

func convertToInt32Arr(arr []int) []int32 {
	i := make([]int32, 0)
	for _, v := range arr {
		i = append(i, int32(v))
	}
	return i
}

// Original code from hackerrank

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var arr []int32

// 	for i := 0; i < int(arrCount); i++ {
// 		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		arrItem := int32(arrItemTemp)
// 		arr = append(arr, arrItem)
// 	}

// 	result := minDiff(arr)

// 	fmt.Fprintf(writer, "%d\n", result)

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
