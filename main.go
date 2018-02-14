package main

import "fmt"

func main() {
	gotoLabel()
	sameNameVariableAndLabel()
	fallthroughLabel()
	breakOuterLoopLabel()
	breakSwitchLabel()
	continueOuterLoopLabel()
}

func gotoLabel() {
	fmt.Println("goto label results:")
	fmt.Println(1)
	goto End
	fmt.Println(2) // won't be executed
End:
	fmt.Println(3)
}

func sameNameVariableAndLabel() {
	fmt.Println("identical var and label name results:")
	x := 1
	goto x
x:
	fmt.Println(x)
}

func fallthroughLabel() {
	fmt.Println("no fallthrough label results:")
	switch 1 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}

	fmt.Println("fallthrough label results:")
	switch 1 {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	}
}

func breakOuterLoopLabel() {
	fmt.Println("break outerloop label results:")
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			break OuterLoop
		}
	}
}

func breakSwitchLabel() {
	fmt.Println("break switch label results:")
SwitchStatement:
	switch 1 {
	case 1:
		fmt.Println(1)
		for i := 0; i < 10; i++ {
			break SwitchStatement
		}
		fmt.Println(2)
	}
	fmt.Println(3)
}

func continueOuterLoopLabel() {
	fmt.Println("continue outerloop label results:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			continue OuterLoop
		}
	}
}
