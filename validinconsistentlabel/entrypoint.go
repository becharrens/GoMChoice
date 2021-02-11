package validinconsistentlabel

import (
   "GoMChoice/validinconsistentlabel/messages"
   "GoMChoice/validinconsistentlabel/roles"
   "GoMChoice/validinconsistentlabel/results"
   "GoMChoice/validinconsistentlabel/channels"
   "GoMChoice/validinconsistentlabel/callbacks"
   "sync"
)

type Init_Validinconsistentlabel_Env interface {
    New_A_Env() callbacks.A_Env
    New_B_Env() callbacks.B_Env
}

type Validinconsistentlabel_Result_Env interface {
    A_Result(result results.A_Result) 
    B_Result(result results.B_Result) 
}


func Start_A(resultEnv Validinconsistentlabel_Result_Env, wg *sync.WaitGroup, roleChannels channels.A_Chan, env callbacks.A_Env) {
    defer wg.Done()
    result := roles.A(roleChannels, env)
    resultEnv.A_Result(result)
}

func Start_B(resultEnv Validinconsistentlabel_Result_Env, wg *sync.WaitGroup, roleChannels channels.B_Chan, env callbacks.B_Env) {
    defer wg.Done()
    result := roles.B(roleChannels, env)
    resultEnv.B_Result(result)
}


func Validinconsistentlabel(initEnv Init_Validinconsistentlabel_Env, resultEnv Validinconsistentlabel_Result_Env) {
    a_b_label_chan := make(chan messages.Validinconsistentlabel_Label)
    a_b_int_chan := make(chan int)
    
    a_chan := channels.A_Chan{
        B_label: a_b_label_chan,
        B_int: a_b_int_chan,
    }
    b_chan := channels.B_Chan{
        A_label: a_b_label_chan,
        A_int: a_b_int_chan,
    }

    var wg sync.WaitGroup

    wg.Add(2)
    a_env := initEnv.New_A_Env()
    b_env := initEnv.New_B_Env()
    go Start_A(resultEnv, &wg, a_chan, a_env)
    go Start_B(resultEnv, &wg, b_chan, b_env)

    wg.Wait()
}