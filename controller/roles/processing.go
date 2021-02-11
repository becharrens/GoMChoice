package roles

import (
	"GoMChoice/controller/callbacks"
	"GoMChoice/controller/channels"
	"GoMChoice/controller/messages"
	"GoMChoice/controller/results"
)

func Processing(roleChannels channels.Processing_Chan, env callbacks.Processing_Env) results.Processing_Result {
	select {
	case label_from_input1 := <-roleChannels.Input1_label:
		switch label_from_input1 {
		case messages.Packet:
			env.Packet_From_Input1()

		t0:
			for {
				env.Packet_To_Routing()
				roleChannels.Routing_label <- messages.Packet

				select {
				case label_from_input1_2 := <-roleChannels.Input1_label:
					switch label_from_input1_2 {
					case messages.Packet:
						env.Packet_From_Input1_2()
						// time.Sleep(1 * time.Microsecond)
						continue t0
					default:
						panic("Invalid choice was made")
					}
				case label_from_input2 := <-roleChannels.Input2_label:
					switch label_from_input2 {
					case messages.Packet:
						env.Packet_From_Input2()
						// time.Sleep(1 * time.Microsecond)

						continue t0
					default:
						panic("Invalid choice was made")
					}
				}
			}
		default:
			panic("Invalid choice was made")
		}
	case label_from_input2_2 := <-roleChannels.Input2_label:
		switch label_from_input2_2 {
		case messages.Packet:
			env.Packet_From_Input2_2()

		t0_2:
			for {
				env.Packet_To_Routing_2()
				roleChannels.Routing_label <- messages.Packet

				select {
				case label_from_input1_3 := <-roleChannels.Input1_label:
					switch label_from_input1_3 {
					case messages.Packet:
						env.Packet_From_Input1_3()
						// time.Sleep(1 * time.Microsecond)

						continue t0_2
					default:
						panic("Invalid choice was made")
					}
				case label_from_input2_3 := <-roleChannels.Input2_label:
					switch label_from_input2_3 {
					case messages.Packet:
						env.Packet_From_Input2_3()
						// time.Sleep(1 * time.Microsecond)

						continue t0_2
					default:
						panic("Invalid choice was made")
					}
				}
			}
		default:
			panic("Invalid choice was made")
		}
	}
}
