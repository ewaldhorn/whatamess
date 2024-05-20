package main

import (
	"fmt"
	"strings"
)

const TRUE string = "-"
const FALSE string = "_"

type Instruction struct {
	name   string
	gate   string
	first  string
	second string
}

func main() {
	var inputCount int
	fmt.Scan(&inputCount)

	var outputCount int
	fmt.Scan(&outputCount)

	// get inputs
	inputs := map[string][]byte{}
	for i := 0; i < inputCount; i++ {
		var inputName, inputSignal string
		fmt.Scan(&inputName, &inputSignal)
		inputs[inputName] = []byte(inputSignal)
	}

	// get instructions
	instructions := []Instruction{}
	for i := 0; i < outputCount; i++ {
		var outputName, _type, inputName1, inputName2 string
		fmt.Scan(&outputName, &_type, &inputName1, &inputName2)
		instructions = append(instructions, Instruction{name: outputName, gate: _type, first: inputName1, second: inputName2})
	}

	// perform instructions
	for i := 0; i < outputCount; i++ {
		outputArray := []string{}

		for pos := 0; pos < len(inputs[instructions[i].first]); pos++ {
			left := string(inputs[string(instructions[i].first)][pos])
			right := string(inputs[string(instructions[i].second)][pos])

			switch instructions[i].gate {
			case "AND":
				outputArray = append(outputArray, doAND(left, right))
			case "OR":
				outputArray = append(outputArray, doOR(left, right))
			case "XOR":
				outputArray = append(outputArray, doXOR(left, right))
			case "NAND":
				outputArray = append(outputArray, doNAND(left, right))
			case "NOR":
				outputArray = append(outputArray, doNOR(left, right))
			case "NXOR":
				outputArray = append(outputArray, doNXOR(left, right))
			}
		}

		fmt.Printf("%s %s\n", instructions[i].name, strings.Join(outputArray, ""))
	}

}

func doAND(left string, right string) string {
	if left == TRUE && right == TRUE {
		return TRUE
	} else {
		return FALSE
	}
}

func doOR(left string, right string) string {
	if left == TRUE || right == TRUE {
		return TRUE
	} else {
		return FALSE
	}
}

func doXOR(left string, right string) string {
	if (left == TRUE && right == FALSE) || (left == FALSE && right == TRUE) {
		return TRUE
	} else {
		return FALSE
	}
}

func doNAND(left string, right string) string {
	if doAND(left, right) == TRUE {
		return FALSE
	} else {
		return TRUE
	}
}

func doNOR(left string, right string) string {
	if doOR(left, right) == TRUE {
		return FALSE
	} else {
		return TRUE
	}
}

func doNXOR(left string, right string) string {
	if doXOR(left, right) == TRUE {
		return FALSE
	} else {
		return TRUE
	}
}
