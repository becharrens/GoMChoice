package roles

import (
   "GoMChoice/pubsub07/channels"
   "GoMChoice/pubsub07/callbacks"
   "GoMChoice/pubsub07/messages"
   "GoMChoice/pubsub07/results"
)

func A1(roleChannels channels.A1_Chan, env callbacks.A1_Env) results.A1_Result {
   env.Sub_To_B()
   roleChannels.B_label <- messages.Sub

t0:
   for {
      <-roleChannels.B_label
      env.Pub_From_B()

      continue t0
   }
}