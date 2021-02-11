package controller

import (
   "GoMChoice/controller/callbacks"
   "GoMChoice/controller/channels"
   "GoMChoice/controller/roles"
   "GoMChoice/controller/results"
   "GoMChoice/controller/messages"
   "sync"
)

type Init_Controller_Env interface {
    New_Processing_Env() callbacks.Processing_Env
    New_Output1_Env() callbacks.Output1_Env
    New_Routing_Env() callbacks.Routing_Env
    New_Output2_Env() callbacks.Output2_Env
    New_Input1_Env() callbacks.Input1_Env
    New_Input2_Env() callbacks.Input2_Env
}

type Controller_Result_Env interface {
    Processing_Result(result results.Processing_Result) 
    Output1_Result(result results.Output1_Result) 
    Routing_Result(result results.Routing_Result) 
    Output2_Result(result results.Output2_Result) 
    Input1_Result(result results.Input1_Result) 
    Input2_Result(result results.Input2_Result) 
}


func Start_Processing(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Processing_Chan, env callbacks.Processing_Env) {
    defer wg.Done()
    result := roles.Processing(roleChannels, env)
    resultEnv.Processing_Result(result)
}

func Start_Output1(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Output1_Chan, env callbacks.Output1_Env) {
    defer wg.Done()
    result := roles.Output1(roleChannels, env)
    resultEnv.Output1_Result(result)
}

func Start_Routing(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Routing_Chan, env callbacks.Routing_Env) {
    defer wg.Done()
    result := roles.Routing(roleChannels, env)
    resultEnv.Routing_Result(result)
}

func Start_Output2(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Output2_Chan, env callbacks.Output2_Env) {
    defer wg.Done()
    result := roles.Output2(roleChannels, env)
    resultEnv.Output2_Result(result)
}

func Start_Input1(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Input1_Chan, env callbacks.Input1_Env) {
    defer wg.Done()
    result := roles.Input1(roleChannels, env)
    resultEnv.Input1_Result(result)
}

func Start_Input2(resultEnv Controller_Result_Env, wg *sync.WaitGroup, roleChannels channels.Input2_Chan, env callbacks.Input2_Env) {
    defer wg.Done()
    result := roles.Input2(roleChannels, env)
    resultEnv.Input2_Result(result)
}


func Controller(initEnv Init_Controller_Env, resultEnv Controller_Result_Env) {
    processing_input1_label_chan := make(chan messages.Controller_Label)
    processing_routing_label_chan := make(chan messages.Controller_Label)
    processing_input2_label_chan := make(chan messages.Controller_Label)
    output1_routing_label_chan := make(chan messages.Controller_Label)
    routing_output2_label_chan := make(chan messages.Controller_Label)
    
    processing_chan := channels.Processing_Chan{
        Input1_label: processing_input1_label_chan,
        Routing_label: processing_routing_label_chan,
        Input2_label: processing_input2_label_chan,
    }
    output1_chan := channels.Output1_Chan{
        Routing_label: output1_routing_label_chan,
    }
    routing_chan := channels.Routing_Chan{
        Processing_label: processing_routing_label_chan,
        Output1_label: output1_routing_label_chan,
        Output2_label: routing_output2_label_chan,
    }
    output2_chan := channels.Output2_Chan{
        Routing_label: routing_output2_label_chan,
    }
    input1_chan := channels.Input1_Chan{
        Processing_label: processing_input1_label_chan,
    }
    input2_chan := channels.Input2_Chan{
        Processing_label: processing_input2_label_chan,
    }

    var wg sync.WaitGroup

    wg.Add(6)
    processing_env := initEnv.New_Processing_Env()
    output1_env := initEnv.New_Output1_Env()
    routing_env := initEnv.New_Routing_Env()
    output2_env := initEnv.New_Output2_Env()
    input1_env := initEnv.New_Input1_Env()
    input2_env := initEnv.New_Input2_Env()
    go Start_Processing(resultEnv, &wg, processing_chan, processing_env)
    go Start_Output1(resultEnv, &wg, output1_chan, output1_env)
    go Start_Routing(resultEnv, &wg, routing_chan, routing_env)
    go Start_Output2(resultEnv, &wg, output2_chan, output2_env)
    go Start_Input1(resultEnv, &wg, input1_chan, input1_env)
    go Start_Input2(resultEnv, &wg, input2_chan, input2_env)

    wg.Wait()
}