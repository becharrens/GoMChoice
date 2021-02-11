package callbacks

import (
	"GoMChoice/validinconsistentlabel/results"
)

type A_Env interface {
	M1_To_B() int
	M1_To_B_2()
	Done() results.A_Result
	M2_To_B() int
	M1_To_B_3() int
	M1_To_B_4()
}

type AState struct {
}

func (A *AState) M1_To_B() int {
	return 1
}

func (A *AState) M1_To_B_2() {
}

func (A *AState) Done() results.A_Result {
	return results.A_Result{}
}

func (A *AState) M2_To_B() int {
	return 2
}

func (A *AState) M1_To_B_3() int {
	return 3
}

func (A *AState) M1_To_B_4() {
}
