package roles

import (
   "GoMChoice/pubsub07/channels"
   "GoMChoice/pubsub07/callbacks"
   "GoMChoice/pubsub07/messages"
   "GoMChoice/pubsub07/results"
)

func A6(roleChannels channels.A6_Chan, env callbacks.A6_Env) results.A6_Result {
   env.Sub_To_B()
   roleChannels.B_label <- messages.Sub

t0:
   for {
      <-roleChannels.B_label
      env.Pub_From_B()

      continue t0
   }
}