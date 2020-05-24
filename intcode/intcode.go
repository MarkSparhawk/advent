package main

type program struct {
	code   []int
	cur    int
	input  int
	output int
}

//Instruction 1
func (p *program) add() {
	p.code[p.paramLocation(3)] = p.code[p.paramLocation(1)] + p.code[p.paramLocation(2)]
	p.cur += 4
}

//Instruction 2
func (p *program) multiply() {
	p.code[p.paramLocation(3)] = p.code[p.paramLocation(1)] * p.code[p.paramLocation(2)]
	p.cur += 4
}

//Instruction 3
func (p *program) write() {
	p.code[p.paramLocation(1)] = p.input
	p.cur += 2
}

//Instruction 4
func (p *program) read() {
	p.output = p.code[p.paramLocation(1)]
	p.cur += 2
}

//Instruction 5
func (p *program) jumpIfTrue() {
	if p.code[p.paramLocation(1)] != 0 {
		p.cur = p.code[p.paramLocation(2)]
	} else {
		p.cur += 3
	}
}

//Instruction 6
func (p *program) jumpIfFalse() {
	if p.code[p.paramLocation(1)] == 0 {
		p.cur = p.code[p.paramLocation(2)]
	} else {
		p.cur += 3
	}
}

//Instruction 7
func (p *program) lessThan() {
	if p.code[p.paramLocation(1)] < p.code[p.paramLocation(2)] {
		p.code[p.paramLocation(3)] = 1
	} else {
		p.code[p.paramLocation(3)] = 0
	}
	p.cur += 4
}

//Instruction 8
func (p *program) equal() {
	if p.code[p.paramLocation(1)] == p.code[p.paramLocation(2)] {
		p.code[p.paramLocation(3)] = 1
	} else {
		p.code[p.paramLocation(3)] = 0
	}
	p.cur += 4
}

type inst struct {
	opcode int
	params []int
}

func (p *program) intcode() {
	op := p.code[p.cur] % 100
	switch op {
	case 1:
		p.add()
	case 2:
		p.multiply()
	case 3:
		p.write()
	case 4:
		p.read()
	case 5:
		p.jumpIfTrue()
	case 6:
		p.jumpIfFalse()
	case 7:
		p.lessThan()
	case 8:
		p.equal()
	case 99:
		return
	}
}

func (p *program) runProgram() {
	for p.cur < len(p.code) {
		if p.code[p.cur] != 99 {
			p.intcode()
		} else {
			return
		}
	}
	return
}

func (p *program) paramLocation(offset int) int {

	// Set location to memory address stored in cur + offset assuming positional mode
	location := p.code[p.cur+offset]

	// If we're in immediate mode, the parameter is simply stored at the cur + offset
	if (p.code[p.cur] / Pow(10, offset+1) % 10) == 1 {
		location = p.cur + offset
	}

	return location
}

func Pow(a, b int) int {
	//Pow() Integer power function a^b.
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
