package roles

import (
   "GoMChoice/payloads/messages"
   "GoMChoice/payloads/results"
   "GoMChoice/payloads/callbacks"
   "GoMChoice/payloads/channels"
)

func B(roleChannels channels.B_Chan, env callbacks.B_Env) results.B_Result {
   select {
   case roleChannels.A_label <- messages.L3:
      a, c := env.L3_To_A()
      roleChannels.A_string <- a
      roleChannels.A_int <- c

      return env.Done()
   case label_from_a := <-roleChannels.A_label:
      switch label_from_a {
      case messages.L1:
         x := <-roleChannels.A_int
         env.L1_From_A(x)

         p_0, y := env.L2_To_A()
         roleChannels.A_label <- messages.L2
         roleChannels.A_int <- p_0
         roleChannels.A_int <- y

         return env.Done()
      default:
         panic("Invalid choice was made")
      }
   }
}