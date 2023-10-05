package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

var Stack1 Stack
var Stack2 Stack
var Stack3 Stack
var Stack4 Stack
var Stack5 Stack
var Stack6 Stack
var Stack7 Stack
var Stack8 Stack
var Stack9 Stack
var moves []string
var numChanges string
var takeFrom string
var giveTo string


// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) PushTop(str string) {
	*s = append([]string{str},*s...) // Simply append the new value to the end of the stack
}

func (s *Stack) Push(str string) {
	*s = append(*s,str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) PopTop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		element := (*s)[0] // Index into the slice and obtain the element.
		*s = (*s)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}


// func setupStacks() (Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, []string) {
// 	// var stack Stack // create a stack variable of type Stack
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		y := scanner.Text()
// 		if strings.Contains(y, "[") {
// 			one := string(y[0:3])
// 			if !strings.Contains(one, " ") {
// 				Stack1.PushTop(one)
// 			}
// 			two := string(y[4:7])
// 			if !strings.Contains(two, " ") {
// 				Stack2.PushTop(two)
// 			}
// 			three := string(y[8:11])
// 			if !strings.Contains(three, " ") {
// 				Stack3.PushTop(three)
// 			}
// 			four := string(y[12:15])
// 			if !strings.Contains(four, " ") {
// 				Stack4.PushTop(four)
// 			}
// 			five := string(y[16:19])
// 			if !strings.Contains(five, " ") {
// 				Stack5.PushTop(five)
// 			}
// 			six := string(y[20:23])
// 			if !strings.Contains(six, " ") {
// 				Stack6.PushTop(six)
// 			}
// 			seven := string(y[24:27])
// 			if !strings.Contains(seven, " ") {
// 				Stack7.PushTop(seven)
// 			}
// 			eight := string(y[28:31])
// 			if !strings.Contains(eight, " ") {
// 				Stack8.PushTop(eight)
// 			}
// 			nine := string(y[32:35])
// 			if !strings.Contains(nine, " ") {
// 				Stack9.PushTop(nine)
// 			}
// 		} else if strings.Contains(y, "move") {
// 			moves = append(moves, y)
// 		}
// 	}
// 	return Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9, moves
// }

// func makeMoves(Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9 Stack, moves []string)(Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack) {
// 	var takenStack Stack
// 	var addToStack Stack
// 	for _, move := range moves {
// 		if len(move) == 19 {
// 			takeFrom = string(move[13])
// 		} else {
// 			takeFrom = string(move[12])
// 		}
// 		fmt.Println(takeFrom)
// 		switch takeFrom {
// 		case "1":
// 			takenStack = Stack1
// 		case "2":
// 			takenStack = Stack2
// 		case "3":
// 			takenStack = Stack3
// 		case "4":
// 			takenStack = Stack4
// 		case "5":
// 			takenStack = Stack5
// 		case "6":
// 			takenStack = Stack6
// 		case "7":
// 			takenStack = Stack7
// 		case "8":
// 			takenStack = Stack8
// 		case "9":
// 			takenStack = Stack9
// 		}
// 		if len(move) == 19 {
// 			giveTo = string(move[18])
// 		} else {
// 			giveTo= string(move[17])
// 		}

// 		switch giveTo {
// 		case "1":
// 			addToStack = Stack1
// 		case "2":
// 			addToStack = Stack2
// 		case "3":
// 			addToStack = Stack3
// 		case "4":
// 			addToStack = Stack4
// 		case "5":
// 			addToStack = Stack5
// 		case "6":
// 			addToStack = Stack6
// 		case "7":
// 			addToStack = Stack7
// 		case "8":
// 			addToStack = Stack8
// 		case "9":
// 			addToStack = Stack9
// 		}
// 		if len(move) == 19 {
// 			numChanges = string(move[5:7])
// 		} else {
// 			numChanges = string(move[5])
// 		}
// 		finalNum,err := strconv.Atoi(numChanges)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		for i := 0; i < finalNum; i++ {
// 			x, _ := takenStack.Pop()
// 			addToStack.Push(x)
// 		}

// 		switch takeFrom {
// 		case "1":
// 			Stack1 = takenStack
// 		case "2":
// 			Stack2 = takenStack
// 		case "3":
// 			Stack3 = takenStack
// 		case "4":
// 			Stack4 = takenStack
// 		case "5":
// 			Stack5 = takenStack
// 		case "6":
// 			Stack6 = takenStack
// 		case "7":
// 			Stack7 = takenStack
// 		case "8":
// 			Stack8 = takenStack
// 		case "9":
// 			Stack9 = takenStack
// 		}

// 		switch giveTo {
// 		case "1":
// 			Stack1 = addToStack
// 		case "2":
// 			Stack2 = addToStack
// 		case "3":
// 			Stack3 = addToStack
// 		case "4":
// 			Stack4 = addToStack
// 		case "5":
// 			Stack5 = addToStack
// 		case "6":
// 			Stack6 = addToStack
// 		case "7":
// 			Stack7 = addToStack
// 		case "8":
// 			Stack8 = addToStack
// 		case "9":
// 			Stack9 = addToStack
// 		}
// 	}
// 	return Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9
// }

func setupStacksPartTwo() (Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, []string) {
	// var stack Stack // create a stack variable of type Stack
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		y := scanner.Text()
		if strings.Contains(y, "[") {
			one := string(y[0:3])
			if !strings.Contains(one, " ") {
				Stack1.Push(one)
			}
			two := string(y[4:7])
			if !strings.Contains(two, " ") {
				Stack2.Push(two)
			}
			three := string(y[8:11])
			if !strings.Contains(three, " ") {
				Stack3.Push(three)
			}
			four := string(y[12:15])
			if !strings.Contains(four, " ") {
				Stack4.Push(four)
			}
			five := string(y[16:19])
			if !strings.Contains(five, " ") {
				Stack5.Push(five)
			}
			six := string(y[20:23])
			if !strings.Contains(six, " ") {
				Stack6.Push(six)
			}
			seven := string(y[24:27])
			if !strings.Contains(seven, " ") {
				Stack7.Push(seven)
			}
			eight := string(y[28:31])
			if !strings.Contains(eight, " ") {
				Stack8.Push(eight)
			}
			nine := string(y[32:35])
			if !strings.Contains(nine, " ") {
				Stack9.Push(nine)
			}
		} else if strings.Contains(y, "move") {
			moves = append(moves, y)
		}
	}
	return Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9, moves
}

func partTwo(Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9 Stack, moves []string)(Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack) {
	var takenStack Stack
	var addToStack Stack
	var revStack Stack
	// var testStack Stack
	for _, move := range moves {
		if len(move) == 19 {
			takeFrom = string(move[13])
		} else {
			takeFrom = string(move[12])
		}
		fmt.Println(takeFrom)
		switch takeFrom {
		case "1":
			takenStack = Stack1
		case "2":
			takenStack = Stack2
		case "3":
			takenStack = Stack3
		case "4":
			takenStack = Stack4
		case "5":
			takenStack = Stack5
		case "6":
			takenStack = Stack6
		case "7":
			takenStack = Stack7
		case "8":
			takenStack = Stack8
		case "9":
			takenStack = Stack9
		}
		if len(move) == 19 {
			giveTo = string(move[18])
		} else {
			giveTo= string(move[17])
		}

		switch giveTo {
		case "1":
			addToStack = Stack1
		case "2":
			addToStack = Stack2
		case "3":
			addToStack = Stack3
		case "4":
			addToStack = Stack4
		case "5":
			addToStack = Stack5
		case "6":
			addToStack = Stack6
		case "7":
			addToStack = Stack7
		case "8":
			addToStack = Stack8
		case "9":
			addToStack = Stack9
		}
		if len(move) == 19 {
			numChanges = string(move[5:7])
		} else {
			numChanges = string(move[5])
		}
		fmt.Println(numChanges)

		finalNum,err := strconv.Atoi(numChanges)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(finalNum)
		for i := 0; i < finalNum; i++ {
			x, _ := takenStack.PopTop()
			revStack.PushTop(x)
		}
		fmt.Println(revStack)

		for i := 0; i < finalNum; i++ {
			x, _ := revStack.PopTop()
			addToStack.PushTop(x)
		}
	

		switch takeFrom {
		case "1":
			Stack1 = takenStack
		case "2":
			Stack2 = takenStack
		case "3":
			Stack3 = takenStack
		case "4":
			Stack4 = takenStack
		case "5":
			Stack5 = takenStack
		case "6":
			Stack6 = takenStack
		case "7":
			Stack7 = takenStack
		case "8":
			Stack8 = takenStack
		case "9":
			Stack9 = takenStack
		}

		switch giveTo {
		case "1":
			Stack1 = addToStack
		case "2":
			Stack2 = addToStack
		case "3":
			Stack3 = addToStack
		case "4":
			Stack4 = addToStack
		case "5":
			Stack5 = addToStack
		case "6":
			Stack6 = addToStack
		case "7":
			Stack7 = addToStack
		case "8":
			Stack8 = addToStack
		case "9":
			Stack9 = addToStack
		}
	}
	return Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9
}

func main() {
	stack1,stack2,stack3,stack4,stack5,stack6,stack7,stack8,stack9,moves := setupStacksPartTwo()

	stack1New,stack2New,stack3New,stack4New,stack5New,stack6New,stack7New,stack8New,stack9New :=partTwo(stack1,stack2,stack3,stack4,stack5,stack6,stack7,stack8,stack9,moves)
	fmt.Printf("Stack 1: %v\n", stack1New)
	fmt.Printf("Stack 2: %v\n", stack2New)
	fmt.Printf("Stack 3: %v\n", stack3New)
	fmt.Printf("Stack 4: %v\n", stack4New)
	fmt.Printf("Stack 5: %v\n", stack5New)
	fmt.Printf("Stack 6: %v\n", stack6New)
	fmt.Printf("Stack 7: %v\n", stack7New)
	fmt.Printf("Stack 8: %v\n", stack8New)
	fmt.Printf("Stack 9: %v\n", stack9New)
}
