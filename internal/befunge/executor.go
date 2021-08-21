package befunge

func ExecuteCode(code string, stack []int) int {
	f := field{}
	f.changeCode(code)
	f.stack = stack

	for i := 0; i < 64 && !f.do(); i++ {
	}

	return f.popStack()
}
