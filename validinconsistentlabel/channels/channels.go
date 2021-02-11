package channels

import (
   "GoMChoice/validinconsistentlabel/messages"
)

type A_Chan struct {
    B_label chan messages.Validinconsistentlabel_Label
    B_int chan int
}

type B_Chan struct {
    A_label chan messages.Validinconsistentlabel_Label
    A_int chan int
}