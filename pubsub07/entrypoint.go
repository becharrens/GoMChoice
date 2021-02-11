package pubsub07

import (
   "GoMChoice/pubsub07/results"
   "GoMChoice/pubsub07/callbacks"
   "GoMChoice/pubsub07/roles"
   "GoMChoice/pubsub07/channels"
   "GoMChoice/pubsub07/messages"
   "sync"
)

type Init_Pubsub07_Env interface {
    New_A7_Env() callbacks.A7_Env
    New_B_Env() callbacks.B_Env
    New_A4_Env() callbacks.A4_Env
    New_A1_Env() callbacks.A1_Env
    New_A5_Env() callbacks.A5_Env
    New_A2_Env() callbacks.A2_Env
    New_A3_Env() callbacks.A3_Env
    New_A6_Env() callbacks.A6_Env
}

type Pubsub07_Result_Env interface {
    A7_Result(result results.A7_Result) 
    B_Result(result results.B_Result) 
    A4_Result(result results.A4_Result) 
    A1_Result(result results.A1_Result) 
    A5_Result(result results.A5_Result) 
    A2_Result(result results.A2_Result) 
    A3_Result(result results.A3_Result) 
    A6_Result(result results.A6_Result) 
}


func Start_A7(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A7_Chan, env callbacks.A7_Env) {
    defer wg.Done()
    result := roles.A7(roleChannels, env)
    resultEnv.A7_Result(result)
}

func Start_B(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.B_Chan, env callbacks.B_Env) {
    defer wg.Done()
    result := roles.B(roleChannels, env)
    resultEnv.B_Result(result)
}

func Start_A4(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A4_Chan, env callbacks.A4_Env) {
    defer wg.Done()
    result := roles.A4(roleChannels, env)
    resultEnv.A4_Result(result)
}

func Start_A1(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A1_Chan, env callbacks.A1_Env) {
    defer wg.Done()
    result := roles.A1(roleChannels, env)
    resultEnv.A1_Result(result)
}

func Start_A5(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A5_Chan, env callbacks.A5_Env) {
    defer wg.Done()
    result := roles.A5(roleChannels, env)
    resultEnv.A5_Result(result)
}

func Start_A2(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A2_Chan, env callbacks.A2_Env) {
    defer wg.Done()
    result := roles.A2(roleChannels, env)
    resultEnv.A2_Result(result)
}

func Start_A3(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A3_Chan, env callbacks.A3_Env) {
    defer wg.Done()
    result := roles.A3(roleChannels, env)
    resultEnv.A3_Result(result)
}

func Start_A6(resultEnv Pubsub07_Result_Env, wg *sync.WaitGroup, roleChannels channels.A6_Chan, env callbacks.A6_Env) {
    defer wg.Done()
    result := roles.A6(roleChannels, env)
    resultEnv.A6_Result(result)
}


func Pubsub07(initEnv Init_Pubsub07_Env, resultEnv Pubsub07_Result_Env) {
    a7_b_label_chan := make(chan messages.Pubsub07_Label)
    b_a6_label_chan := make(chan messages.Pubsub07_Label)
    b_a1_label_chan := make(chan messages.Pubsub07_Label)
    b_a2_label_chan := make(chan messages.Pubsub07_Label)
    b_a3_label_chan := make(chan messages.Pubsub07_Label)
    b_a4_label_chan := make(chan messages.Pubsub07_Label)
    b_a5_label_chan := make(chan messages.Pubsub07_Label)
    
    a7_chan := channels.A7_Chan{
        B_label: a7_b_label_chan,
    }
    b_chan := channels.B_Chan{
        A6_label: b_a6_label_chan,
        A1_label: b_a1_label_chan,
        A2_label: b_a2_label_chan,
        A3_label: b_a3_label_chan,
        A4_label: b_a4_label_chan,
        A5_label: b_a5_label_chan,
        A7_label: a7_b_label_chan,
    }
    a4_chan := channels.A4_Chan{
        B_label: b_a4_label_chan,
    }
    a1_chan := channels.A1_Chan{
        B_label: b_a1_label_chan,
    }
    a5_chan := channels.A5_Chan{
        B_label: b_a5_label_chan,
    }
    a2_chan := channels.A2_Chan{
        B_label: b_a2_label_chan,
    }
    a3_chan := channels.A3_Chan{
        B_label: b_a3_label_chan,
    }
    a6_chan := channels.A6_Chan{
        B_label: b_a6_label_chan,
    }

    var wg sync.WaitGroup

    wg.Add(8)
    a7_env := initEnv.New_A7_Env()
    b_env := initEnv.New_B_Env()
    a4_env := initEnv.New_A4_Env()
    a1_env := initEnv.New_A1_Env()
    a5_env := initEnv.New_A5_Env()
    a2_env := initEnv.New_A2_Env()
    a3_env := initEnv.New_A3_Env()
    a6_env := initEnv.New_A6_Env()
    go Start_A7(resultEnv, &wg, a7_chan, a7_env)
    go Start_B(resultEnv, &wg, b_chan, b_env)
    go Start_A4(resultEnv, &wg, a4_chan, a4_env)
    go Start_A1(resultEnv, &wg, a1_chan, a1_env)
    go Start_A5(resultEnv, &wg, a5_chan, a5_env)
    go Start_A2(resultEnv, &wg, a2_chan, a2_env)
    go Start_A3(resultEnv, &wg, a3_chan, a3_env)
    go Start_A6(resultEnv, &wg, a6_chan, a6_env)

    wg.Wait()
}