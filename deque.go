package main

import (
	"bufio"
	"fmt"
	"go/types"
	"math"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	Type string
	absoluteType types.Type
	maxSize int
	leftIndex int
	rightIndex int
	intSlice []int
	int64Slice []int64
	byteSlice []byte //stringもbyteで管理
	boolSlice []bool
	judgeMentSlice []bool
}

func NewDeque(maxSize int, Type string) *Deque {
	output := new(Deque)
	output.maxSize = maxSize
	output.Type = Type
	output.leftIndex = maxSize-2
	output.rightIndex = maxSize-1
	output.judgeMentSlice = make([]bool, maxSize*2)
	switch Type {
	case "byte":
		fallthrough
	case "string":
		output.byteSlice = make([]byte, maxSize*2)
	case "int64":
		output.int64Slice = make([]int64, maxSize*2)
	case "int":
		output.intSlice = make([]int, maxSize*2)
	case "bool":
		output.boolSlice = make([]bool, maxSize*2)
	default:
		panic("dequeの型が不明です")
	}
	return output
}

func (deque *Deque) push (value interface{}) {
	deque.rightIndex++
	if deque.rightIndex == deque.maxSize*2 {
		panic("dequeのindexが、最大indexを超えてます。maxSizeが小さい可能性があります。")
	}
	switch deque.Type {
	case "byte":
		fallthrough
	case "string":
		deque.byteSlice[deque.rightIndex] = value.(byte)
	case "int64":
		deque.int64Slice[deque.rightIndex] = value.(int64)
	case "int":
		deque.intSlice[deque.rightIndex] = value.(int)
	case "bool":
		deque.boolSlice[deque.rightIndex] = value.(bool)
	}
	deque.judgeMentSlice[deque.rightIndex] = true
}
func (deque *Deque) unshift (value interface{}) {
	deque.leftIndex--
	if deque.leftIndex < 0 {
		panic("dequeのindexが、最小indexを超えてます。maxSizeが小さい可能性があります。")
	}
	switch deque.Type {
	case "byte":
		fallthrough
	case "string":
		deque.byteSlice[deque.leftIndex] = value.(byte)
	case "int64":
		deque.int64Slice[deque.leftIndex] = value.(int64)
	case "int":
		deque.intSlice[deque.leftIndex] = value.(int)
	case "bool":
		deque.boolSlice[deque.leftIndex] = value.(bool)
	}
	deque.judgeMentSlice[deque.leftIndex] = true
}

func (deque *Deque) pop () interface{}{
	var output interface{}
	if !deque.judgeMentSlice[deque.rightIndex] {
		// 初期値を返すようにしとく
		switch deque.Type {
		case "byte":
			return byte(0)
		case "string":
			return ""
		case "int64":
			return int64(0)
		case "int":
			return int(0)
		case "bool":
			return false
		}
	}
	switch deque.Type {
	case "byte":
		output = deque.byteSlice[deque.rightIndex]
	case "string":
		output = string(deque.byteSlice[deque.rightIndex])
	case "int64":
		output = deque.int64Slice[deque.rightIndex]
	case "int":
		output = deque.intSlice[deque.rightIndex]
	case "bool":
		output = deque.boolSlice[deque.rightIndex]
	}
	deque.judgeMentSlice[deque.rightIndex] = false
	deque.rightIndex--
	if deque.rightIndex < deque.leftIndex {
		// leftとrightが入れ替わるようなことがあれば、位置をswap
		swap := deque.leftIndex
		deque.leftIndex = deque.rightIndex
		deque.rightIndex = swap
	}
	return output
}

func main() {
	scan_init()
	get := NewDeque(5, "int")
	fmt.Println(get)
	get.push(3)
	fmt.Println(get)
	fmt.Println(get.pop())
	fmt.Println(get)
	fmt.Print(get.pop())
	fmt.Println(get)
}

var sc = bufio.NewScanner(os.Stdin)

func scan_init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}
func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
func scanInts(n int) []int {
	take := make([]int, n)
	for i := 0; i < n; i++ {
		take[i] = scanInt()
	}
	return take
}
func scan() string {
	sc.Scan()
	return sc.Text()
}

var rdr = bufio.NewReaderSize(os.Stdin, 200000)

func readLine() string {
	buf := make([]byte, 0, 200000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
func readLineToNumber() int {
	S := readLine()
	number, _ := strconv.Atoi(S)
	return number
}
func readLineToSlice() []string {
	S := readLine()
	list := strings.Split(S, "")
	return list
}
func readLineToNumberSlice(memo map[string]int) []int {
	// err時は-1を代入
	S := readLine()
	intList := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if val, ok := memo[string(S[i])]; ok {
			intList[i] = val
		} else {
			intList[i] = -1
		}
	}
	return intList
}