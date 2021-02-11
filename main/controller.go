package main

import (
	"GoMChoice/controller"
	"GoMChoice/controller/callbacks"
	"GoMChoice/controller/results"
)

type ControllerInitEnv struct {
}

func (c *ControllerInitEnv) New_Processing_Env() callbacks.Processing_Env {
	return &callbacks.ProcessingState{}
}

func (c *ControllerInitEnv) New_Output1_Env() callbacks.Output1_Env {
	return &callbacks.Output1State{}
}

func (c *ControllerInitEnv) New_Routing_Env() callbacks.Routing_Env {
	return &callbacks.RoutingState{}
}

func (c *ControllerInitEnv) New_Output2_Env() callbacks.Output2_Env {
	return &callbacks.Output2State{}
}

func (c *ControllerInitEnv) New_Input1_Env() callbacks.Input1_Env {
	return &callbacks.Input1State{}
}

func (c *ControllerInitEnv) New_Input2_Env() callbacks.Input2_Env {
	return &callbacks.Input2State{}
}

type ControllerResultEnv struct {
}

func (c *ControllerResultEnv) Processing_Result(result results.Processing_Result) {
}

func (c *ControllerResultEnv) Output1_Result(result results.Output1_Result) {
}

func (c *ControllerResultEnv) Routing_Result(result results.Routing_Result) {
}

func (c *ControllerResultEnv) Output2_Result(result results.Output2_Result) {
}

func (c *ControllerResultEnv) Input1_Result(result results.Input1_Result) {
}

func (c *ControllerResultEnv) Input2_Result(result results.Input2_Result) {
}

func NewControllerInitEnv() *ControllerInitEnv {
	return &ControllerInitEnv{}
}

func NewControllerResultEnv() *ControllerResultEnv {
	return &ControllerResultEnv{}
}

func RunController() {
	initEnv := NewControllerInitEnv()
	resultEnv := NewControllerResultEnv()
	controller.Controller(initEnv, resultEnv)
}
