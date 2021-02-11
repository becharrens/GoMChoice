package callbacks

import "fmt"

type Processing_Env interface {
    Packet_From_Input1() 
    Packet_To_Routing() 
    Packet_From_Input1_2() 
    Packet_From_Input2() 
    Packet_From_Input2_2() 
    Packet_To_Routing_2() 
    Packet_From_Input1_3() 
    Packet_From_Input2_3() 
}

type ProcessingState struct {

}

func (p *ProcessingState) Packet_From_Input1() {
    fmt.Println("processing: packet from Input1")
}

func (p *ProcessingState) Packet_To_Routing() {
    fmt.Println("processing: packet to Routing")
}

func (p *ProcessingState) Packet_From_Input1_2() {
    fmt.Println("processing: packet from Input1")
}

func (p *ProcessingState) Packet_From_Input2() {
    fmt.Println("processing: packet from Inpu2")
}

func (p *ProcessingState) Packet_From_Input2_2() {
    fmt.Println("processing: packet from Input2")
}

func (p *ProcessingState) Packet_To_Routing_2() {
    fmt.Println("processing: packet to Routing")
}

func (p *ProcessingState) Packet_From_Input1_3() {
    fmt.Println("processing: packet from Input1")
}

func (p *ProcessingState) Packet_From_Input2_3() {
    fmt.Println("processing: packet from Input2")
}
