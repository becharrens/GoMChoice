package callbacks

import (
	"GoMChoice/validinconsistentlabel/results"
	"fmt"
)

type B_Env interface {
	M2_From_A(a int)
	M1_From_A(a int)
	M1_From_A_2()
	Done() results.B_Result
	M1_From_A_3(a int)
	M1_From_A_4()
}

type BState struct {
}

func (B *BState) M2_From_A(a int) {
	fmt.Println("b: m2 -", a)
}

func (B *BState) M1_From_A(a int) {
	fmt.Println("b: m1 -", a)
}

func (B *BState) M1_From_A_2() {
	fmt.Println("b: m1")
}

func (B *BState) Done() results.B_Result {
	return results.B_Result{}
}

func (B *BState) M1_From_A_3(a int) {
	fmt.Println("b: m1 -", a)
}

func (B *BState) M1_From_A_4() {
	fmt.Println("b: m1")
}
