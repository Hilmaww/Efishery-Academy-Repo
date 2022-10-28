package main

type bandungdatar interface {
	keliling()
}

func panggilKeliling(bd bandungdatar) {
	bd.keliling()
}

type bulat struct {
	x: string
}
type lingkaran struct {
	y: string
}


func main() {
	blt := bulat{x: "Hello"}
	lkr := lingkaran{y: "Hi"}
}
