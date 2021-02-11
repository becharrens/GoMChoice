package roles

import (
   "GoMChoice/pubsub07/channels"
   "GoMChoice/pubsub07/callbacks"
   "GoMChoice/pubsub07/messages"
   "GoMChoice/pubsub07/results"
)

func A4(roleChannels channels.A4_Chan, env callbacks.A4_Env) results.A4_Result {
   env.Sub_To_B()
   roleChannels.B_label <- messages.Sub

t0:
   for {
      <-roleChannels.B_label
      env.Pub_From_B()

      continue t0
   }
}