package intcode

// Program contains all the Code to be exectuted and state.
type Program struct {
	Code   []int
	Cur    int
	Input  int
	Output int
}

//Instruction 1
func (p *Program) add() {
	p.Code[p.paramLocation(3)] = p.Code[p.paramLocation(1)] + p.Code[p.paramLocation(2)]
	p.Cur += 4
}

//Instruction 2
func (p *Program) multiply() {
	p.Code[p.paramLocation(3)] = p.Code[p.paramLocation(1)] * p.Code[p.paramLocation(2)]
	p.Cur += 4
}

//Instruction 3
func (p *Program) write() {
	p.Code[p.paramLocation(1)] = p.Input
	p.Cur += 2
}

//Instruction 4
func (p *Program) read() {
	p.Output = p.Code[p.paramLocation(1)]
	p.Cur += 2
}

//Instruction 5
func (p *Program) jumpIfTrue() {
	if p.Code[p.paramLocation(1)] != 0 {
		p.Cur = p.Code[p.paramLocation(2)]
	} else {
		p.Cur += 3
	}
}

//Instruction 6
func (p *Program) jumpIfFalse() {
	if p.Code[p.paramLocation(1)] == 0 {
		p.Cur = p.Code[p.paramLocation(2)]
	} else {
		p.Cur += 3
	}
}

//Instruction 7
func (p *Program) lessThan() {
	if p.Code[p.paramLocation(1)] < p.Code[p.paramLocation(2)] {
		p.Code[p.paramLocation(3)] = 1
	} else {
		p.Code[p.paramLocation(3)] = 0
	}
	p.Cur += 4
}

//Instruction 8
func (p *Program) equal() {
	if p.Code[p.paramLocation(1)] == p.Code[p.paramLocation(2)] {
		p.Code[p.paramLocation(3)] = 1
	} else {
		p.Code[p.paramLocation(3)] = 0
	}
	p.Cur += 4
}

type inst struct {
	opCode int
	params []int
}

func (p *Program) intCode() {
	op := p.Code[p.Cur] % 100
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

// Run executes the code in the struct in place.
func (p *Program) Run() {
	for p.Cur < len(p.Code) {
		if p.Code[p.Cur] != 99 {
			p.intCode()
		} else {
			return
		}
	}
	return
}

func (p *Program) paramLocation(offset int) int {

	// Set location to memory address stored in Cur + offset assuming positional mode
	location := p.Code[p.Cur+offset]

	// If we're in immediate mode, the parameter is simply stored at the Cur + offset
	if (p.Code[p.Cur] / pow(10, offset+1) % 10) == 1 {
		location = p.Cur + offset
	}

	return location
}

func pow(a, b int) int {
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
