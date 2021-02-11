package roles

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/results"
   "GoMChoice/controller/channels"
)

func Output2(roleChannels channels.Output2_Chan, env callbacks.Output2_Env) results.Output2_Result {
   <-roleChannels.Routing_label
   env.Packet_From_Routing()

t0:
   for {
      <-roleChannels.Routing_label
      env.Packet_From_Routing_2()

      continue t0
   }
}