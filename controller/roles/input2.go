package roles

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/results"
   "GoMChoice/controller/channels"
   "GoMChoice/controller/messages"
)

func Input2(roleChannels channels.Input2_Chan, env callbacks.Input2_Env) results.Input2_Result {
   env.Packet_To_Processing()
   roleChannels.Processing_label <- messages.Packet

t0:
   for {
      env.Packet_To_Processing_2()
      roleChannels.Processing_label <- messages.Packet

      continue t0
   }
}