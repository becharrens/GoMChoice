package roles

import (
   "GoMChoice/payloads/messages"
   "GoMChoice/payloads/results"
   "GoMChoice/payloads/callbacks"
   "GoMChoice/payloads/channels"
)

func A(roleChannels channels.A_Chan, env callbacks.A_Env) results.A_Result {
   select {
   case roleChannels.B_label <- messages.L1:
      x := env.L1_To_B()
      roleChannels.B_int <- x

      <-roleChannels.B_label
      p_0 := <-roleChannels.B_int
      y := <-roleChannels.B_int
      env.L2_From_B(p_0, y)

      return env.Done()
   case label_from_b := <-roleChannels.B_label:
      switch label_from_b {
      case messages.L3:
         a := <-roleChannels.B_string
         c := <-roleChannels.B_int
         env.L3_From_B(a, c)

         return env.Done()
      default:
         panic("Invalid choice was made")
      }
   }
}