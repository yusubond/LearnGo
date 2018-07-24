package sample

func Add(x int, y int) (z int, err error) {
	return x + y, nil
}

func Del(x int, y int) (z int, err error) {
	return  x - y, nil
}