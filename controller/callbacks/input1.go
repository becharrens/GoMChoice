package callbacks

type Input1_Env interface {
    Packet_To_Processing() 
    Packet_To_Processing_2() 
}

type Input1State struct {
}

func (i *Input1State) Packet_To_Processing() {
}

func (i *Input1State) Packet_To_Processing_2() {
}
