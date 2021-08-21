package befunge

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

type field struct {
	pointer     pointer
	field       [][]byte
	stack       []int
	scaleModify int
}
type pointer struct {
	x      int
	y      int
	vector uint8
}

func (f *field) changePointerCell(newValue byte) {
	f.field[f.pointer.y][f.pointer.x] = newValue
}

func (f *field) changeCode(code string) {
	splitted := bytes.Split([]byte(code), []byte("\n"))

	yMax := len(splitted) // Определение размеров поля
	xMax := len(splitted[0])
	for _, bs := range splitted {
		if len(bs) > xMax {
			xMax = len(bs)
		}
	}

	f.field = make([][]uint8, yMax+1) // Пересоздание поля с нужными размерами
	for i := range f.field {
		f.field[i] = make([]uint8, xMax+1)
	}
	var i int
	for y, bs := range splitted {
		for x, b := range bs {
			i++
			f.field[y][x] = b
		}
	}
}

func (f *field) step() {
	switch f.pointer.vector {
	case 1:
		if f.pointer.y > 0 {
			f.pointer.y -= 1
		} else {
			f.pointer.y = len(f.field) - 1
		}
	case 2:
		if f.pointer.x < len(f.field[0])-1 {
			f.pointer.x += 1
		} else {
			f.pointer.x = 0
		}
	case 3:
		if f.pointer.y < len(f.field)-1 {
			f.pointer.y += 1
		} else {
			f.pointer.y = 0
		}
	case 4:
		if f.pointer.x > 0 {
			f.pointer.x -= 1
		} else {
			f.pointer.x = len(f.field[0]) - 1
		}
	}
}
func (f *field) do() bool {
	defer f.step()
	value := int(f.field[f.pointer.y][f.pointer.x])

	switch value {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		i, err := strconv.Atoi(string([]rune{rune(value)}))
		if err == nil {
			f.addStack(i)
		}
	case '+', '-', '*', '/', '%':
		for len(f.stack) < 2 {
			f.addStack(0)
		}
		a := f.popStack()
		b := f.popStack()
		switch value {
		case '+':
			f.addStack(a + b)
		case '-':
			f.addStack(b - a)
		case '*':
			f.addStack(a * b)
		case '/':
			if b == 0 {
				f.addStack(0)
			} else {
				f.addStack(a / b)
			}
		case '%':
			if b == 0 {
				f.addStack(0)
			} else {
				f.addStack(a % b)
			}
		}
	case '!':
		for len(f.stack) < 1 {
			f.addStack(0)
		}
		a := f.popStack()
		if a == 0 {
			f.addStack(1)
		} else {
			f.addStack(0)
		}
	case '`':
		for len(f.stack) < 2 {
			f.addStack(0)
		}
		a := f.popStack()
		b := f.popStack()
		if b > a {
			f.addStack(1)
		} else {
			f.addStack(0)
		}
	case '?':
		rand.Seed(time.Now().UnixNano())
		f.pointer.vector = uint8(rand.Intn(4) + 1)
	case ':':
		if len(f.stack) == 0 {
			f.addStack(0)
		}
		a := f.popStack()
		f.addStack(a)
		f.addStack(a)
	case '\\':
		if len(f.stack) < 1 {
			f.addStack(0)
		}
		if len(f.stack) < 2 {
			f.addStack(0)
			break
		}
		a := f.popStack()
		b := f.popStack()
		f.addStack(a)
		f.addStack(b)
	case '@':
		return true
	case '#':
		f.step()
	case '_':
		if f.popStack() == 0 {
			f.pointer.vector = 2
		} else {
			f.pointer.vector = 4
		}
	case '|':
		if f.popStack() == 0 {
			f.pointer.vector = 3
		} else {
			f.pointer.vector = 1
		}
	case '$':
		f.popStack()
	case '>':
		f.pointer.vector = 2
	case '<':
		f.pointer.vector = 4
	case '^':
		f.pointer.vector = 1
	case 'v':
		f.pointer.vector = 3
	}

	return false
}

func (f *field) popStack() (last int) {
	if len(f.stack) == 0 {
		return 0
	}
	last = f.stack[len(f.stack)-1]
	f.stack = f.stack[:len(f.stack)-1]
	return
}
func (f *field) addStack(add int) {
	f.stack = append(f.stack, add)
}
