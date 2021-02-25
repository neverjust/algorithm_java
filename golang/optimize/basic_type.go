package optimize

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

/*
	[]bytes和string的相互转换
	原理 底层都是byte数组
	go语言原生转换是 内存分配+数据拷贝
	如下的强转换只需要将指针进行转换 不涉及内存分配
	风险
	1. 修改一个string强转换后得到的[]byte会fatal error，因为只做了指针转换，这个值还是不可写的
	2. 通过强转[]byte得到的string，如果原始[]byte被修改，string的值也对应改变
	结论 ：只读场景下放心用强转换
 */
func stringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}


/*
	字符串拼接
	Benchmark_bytesBuffer    805302      435 ns/op    3456 B/op     1 allocs/op
	Benchmark_stringPlus     1204712     972 ns/op    7168 B/op     3 allocs/op
	Benchmark_fmtSprintf     753339     1688 ns/op    7996 B/op    11 allocs/op
	Benchmark_stringsJoin    2737509     440 ns/op    3456 B/op     1 allocs/op
 */
func bytesBuffer(arr []string) string {
	var b = bytes.Buffer{}
	length := 0
	for _, v := range arr {
		length += len(v)
	}
	b.Grow(length)
	for _, v := range arr {
		b.WriteString(v)
	}
	return b.String()
}
func fmtSprintf(arr []string) string {
	var r string
	for _, v := range arr {
		r = fmt.Sprintf("%s%s", r, v)
	}
	return r
}
func stringPlus(arr []string) string {
	var r string
	for _, v := range arr {
		r += v
	}
	return r
}
func stringsJoin(arr []string) string {
	return strings.Join(arr, "")
}


/*
	map和slice
	1. 尽量提前预估容量，减少内存分配次数
	2. 尽量规避使用append，因为需要值拷贝，且涉及到重新申请内存，可能会发生逃逸
	3. 如果无法预估，一般场景下可以考虑申请足够大的空间，并在场景允许的情况下优先考虑复用slice
	Benchmark_UseCap1     980874    1194 ns/op          0 B/op     0 allocs/op
	Benchmark_UseCap2    2278182     543 ns/op          0 B/op     0 allocs/op
	Benchmark_NoCap       119982    8824 ns/op      58616 B/op    14 allocs/op
 */
func useCap1() {
	arr := make([]int, 0, 2048)
	for i := 0; i < 2048; i++ {
		arr = append(arr, i)
	}
}
func useCap2() {
	arr := make([]int, 2048)
	for i := 0; i < 2048; i++ {
		arr[i] = i
	}
}
//slice扩容的主要代码，常规场景下的扩容逻辑为cap<1024时每次翻倍，cap>1024时每次增长25%，此处也可以对应上benchmark中noCap()分配在了堆上，并经过了14次扩容
func noCap() {
	var arr []int
	for i := 0; i < 2048; i++ {
		arr = append(arr, i)
	}
}


