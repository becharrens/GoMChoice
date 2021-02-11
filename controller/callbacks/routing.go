package callbacks

import "fmt"

type Routing_Env interface {
	Packet_From_Processing()
	Packet_To_Output1()
	Packet_From_Processing_2()
	Packet_To_Output2()
	Packet_From_Processing_3()
}

type RoutingState struct {
}

func (r *RoutingState) Packet_From_Processing() {
	fmt.Println("routing: packet from Processing")
}

func (r *RoutingState) Packet_To_Output1() {
	fmt.Println("routing: packet to Output1")
}

func (r *RoutingState) Packet_From_Processing_2() {
	fmt.Println("routing: packet from Processing")
}

func (r *RoutingState) Packet_To_Output2() {
	fmt.Println("routing: packet to Output2")
}

func (r *RoutingState) Packet_From_Processing_3() {
	fmt.Println("routing: packet from Processing")
}
