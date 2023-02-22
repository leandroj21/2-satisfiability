package src

func Abs(num int) int {
	if num < 0 {
		num = -1 * num
	}
	return num
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
