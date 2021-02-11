package channels

import (
   "GoMChoice/pubsub07/messages"
)

type A7_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type B_Chan struct {
    A6_label chan messages.Pubsub07_Label
    A1_label chan messages.Pubsub07_Label
    A2_label chan messages.Pubsub07_Label
    A3_label chan messages.Pubsub07_Label
    A4_label chan messages.Pubsub07_Label
    A5_label chan messages.Pubsub07_Label
    A7_label chan messages.Pubsub07_Label
}

type A4_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type A1_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type A5_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type A2_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type A3_Chan struct {
    B_label chan messages.Pubsub07_Label
}

type A6_Chan struct {
    B_label chan messages.Pubsub07_Label
}