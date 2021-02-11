package callbacks

import (
   "GoMChoice/payloads/results"
)

type B_Env interface {
    L3_To_A() (string, int)
    Done() results.B_Result
    L1_From_A(x int) 
    L2_To_A() (int, int)
}