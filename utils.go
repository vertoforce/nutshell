package nutshell

func clampLimit(l int) int {
	if l > 100 {
		l = 100
	}
	if l <= 0 {
		l = 1
	}
	return l
}
