package main

func main() {
	mmap := make(map[int]int)
	mmap[1] = 1
	i := 2
	for _, v := range mmap {
		mmap[i] = i
		i++
		println(v)
	}
}
