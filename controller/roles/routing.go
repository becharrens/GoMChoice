package roles

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/results"
   "GoMChoice/controller/channels"
   "GoMChoice/controller/messages"
)

func Routing(roleChannels channels.Routing_Chan, env callbacks.Routing_Env) results.Routing_Result {
   <-roleChannels.Processing_label
   env.Packet_From_Processing()

t0:
   for {
      select {
      case roleChannels.Output1_label <- messages.Packet:
         env.Packet_To_Output1()

         <-roleChannels.Processing_label
         env.Packet_From_Processing_2()

         continue t0
      case roleChannels.Output2_label <- messages.Packet:
         env.Packet_To_Output2()

         <-roleChannels.Processing_label
         env.Packet_From_Processing_3()

         continue t0
      }
   }
}