package roles

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/results"
   "GoMChoice/controller/channels"
   "GoMChoice/controller/messages"
)

func Input1(roleChannels channels.Input1_Chan, env callbacks.Input1_Env) results.Input1_Result {
   env.Packet_To_Processing()
   roleChannels.Processing_label <- messages.Packet

t0:
   for {
      env.Packet_To_Processing_2()
      roleChannels.Processing_label <- messages.Packet

      continue t0
   }
}