package roles

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/results"
   "GoMChoice/controller/channels"
)

func Output1(roleChannels channels.Output1_Chan, env callbacks.Output1_Env) results.Output1_Result {
   <-roleChannels.Routing_label
   env.Packet_From_Routing()

t0:
   for {
      <-roleChannels.Routing_label
      env.Packet_From_Routing_2()

      continue t0
   }
}