package callbacks

import (
   "GoMChoice/payloads/results"
)

type A_Env interface {
    L1_To_B() int
    Done() results.A_Result
    L2_From_B(p_0 int, y int) 
    L3_From_B(a string, c int) 
}