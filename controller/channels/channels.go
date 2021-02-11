package channels

import (
   "GoMChoice/controller/messages"
)

type Processing_Chan struct {
    Input1_label chan messages.Controller_Label
    Routing_label chan messages.Controller_Label
    Input2_label chan messages.Controller_Label
}

type Output1_Chan struct {
    Routing_label chan messages.Controller_Label
}

type Routing_Chan struct {
    Processing_label chan messages.Controller_Label
    Output1_label chan messages.Controller_Label
    Output2_label chan messages.Controller_Label
}

type Output2_Chan struct {
    Routing_label chan messages.Controller_Label
}

type Input1_Chan struct {
    Processing_label chan messages.Controller_Label
}

type Input2_Chan struct {
    Processing_label chan messages.Controller_Label
}