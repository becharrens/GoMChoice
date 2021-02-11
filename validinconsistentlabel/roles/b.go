package roles

import (
   "GoMChoice/validinconsistentlabel/messages"
   "GoMChoice/validinconsistentlabel/results"
   "GoMChoice/validinconsistentlabel/channels"
   "GoMChoice/validinconsistentlabel/callbacks"
)

func B(roleChannels channels.B_Chan, env callbacks.B_Env) results.B_Result {
   select {
   case label_from_a := <-roleChannels.A_label:
      switch label_from_a {
      case messages.M2:
         a := <-roleChannels.A_int
         env.M2_From_A(a)

         <-roleChannels.A_label
         a_2 := <-roleChannels.A_int
         env.M1_From_A(a_2)

         <-roleChannels.A_label
         env.M1_From_A_2()

         return env.Done()
      case messages.M1:
         a_3 := <-roleChannels.A_int
         env.M1_From_A_3(a_3)

         <-roleChannels.A_label
         env.M1_From_A_4()

         return env.Done()
      default:
         panic("Invalid choice was made")
      }
   }
}