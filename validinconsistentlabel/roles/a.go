package roles

import (
   "GoMChoice/validinconsistentlabel/messages"
   "GoMChoice/validinconsistentlabel/results"
   "GoMChoice/validinconsistentlabel/channels"
   "GoMChoice/validinconsistentlabel/callbacks"
)

func A(roleChannels channels.A_Chan, env callbacks.A_Env) results.A_Result {
   select {
   case roleChannels.B_label <- messages.M1:
      a := env.M1_To_B()
      roleChannels.B_int <- a

      env.M1_To_B_2()
      roleChannels.B_label <- messages.M1

      return env.Done()
   case roleChannels.B_label <- messages.M2:
      a_2 := env.M2_To_B()
      roleChannels.B_int <- a_2

      a_3 := env.M1_To_B_3()
      roleChannels.B_label <- messages.M1
      roleChannels.B_int <- a_3

      env.M1_To_B_4()
      roleChannels.B_label <- messages.M1

      return env.Done()
   }
}