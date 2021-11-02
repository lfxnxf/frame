package main

func main() {
	a := make([]int, 0)

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	b := a[1:5]
	println(b)
}
