package main

func main() {
	a := make(map[*int]int)
	a[nil] = 1
	println(a[nil])
}
func test1(a map[int]int) {
	for i := 2; i <= 1000000; i++ {
		a[i] = i
	}
}
