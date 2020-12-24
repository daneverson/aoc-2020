package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type operation int

const (
	nop operation = iota
	acc
	jmp
)

type instruction struct {
	op       operation
	value    int
	executed bool
}

func (i *instruction) execute() (int, int) {
	i.executed = true
	switch i.op {
	case acc:
		return i.value, 1
	case jmp:
		return 0, i.value
	}
	// Treat all unknown instructions as "nop" instructions
	return 0, 1
}

type program struct {
	instructions []instruction
}

func (p *program) run() (int, error) {
	v, i := 0, 0
	for i <= len(p.instructions)-1 {
		// If we've already executed this instruction this program is no good.
		if p.instructions[i].executed {
			return 0, errors.New("loop detected")
		}
		p.instructions[i].executed = true

		// Otherwise, execute the instruction as usual
		switch p.instructions[i].op {
		case nop:
			i++
		case acc:
			v += p.instructions[i].value
			i++
		case jmp:
			i += p.instructions[i].value
		}
	}

	// If the instruction index ran past, see how far. If only one, then (probably) it
	// terminated normally.
	// TODO: may need to check this, it's possible a bad 'jmp' happened to just go one past end
	if i == len(p.instructions) {
		return v, nil
	}

	return 0, errors.New("index past end")
}

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	origInstructions := parseInstructionSet(strings.Split(string(contents), "\n"))

	for i, inst := range origInstructions {
		p := program{}
		switch inst.op {
		case jmp:
			p.instructions = make([]instruction, len(origInstructions))
			copy(p.instructions, origInstructions)
			p.instructions[i].op = nop
		case nop:
			p.instructions = make([]instruction, len(origInstructions))
			copy(p.instructions, origInstructions)
			p.instructions[i].op = jmp
		default:
			// we didn't flip an operation, so this isn't a variant to try
			continue
		}

		// Run the program with the altered instruction and check output
		accumulator, err := p.run()
		if err == nil {
			fmt.Printf("PASS - changed instruction %d - (%d, %v)\n", i, accumulator, err)
			return
		}
		fmt.Printf("FAIL - changed instruction %d - (%d, %v)\n", i, accumulator, err)
	}
}

func parseInstructionSet(ss []string) []instruction {
	instructions := make([]instruction, len(ss))
	for i, s := range ss {
		split := strings.Split(s, " ")
		op := nop
		switch split[0] {
		case "acc":
			op = acc
		case "jmp":
			op = jmp
		}
		val, _ := strconv.ParseInt(split[1], 10, 64)
		instructions[i] = instruction{
			op:    op,
			value: int(val),
		}
	}
	return instructions
}
