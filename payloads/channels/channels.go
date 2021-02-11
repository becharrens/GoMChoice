package channels

import (
   "GoMChoice/payloads/messages"
)

type A_Chan struct {
    B_label chan messages.Payloads_Label
    B_int chan int
    B_string chan string
}

type B_Chan struct {
    A_label chan messages.Payloads_Label
    A_string chan string
    A_int chan int
}