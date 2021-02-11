package main

import (
	"GoMChoice/validinconsistentlabel"
	"GoMChoice/validinconsistentlabel/callbacks"
	"GoMChoice/validinconsistentlabel/results"
)

type ValidInconsistentLabelInitEnv struct {
}

func (v *ValidInconsistentLabelInitEnv) New_A_Env() callbacks.A_Env {
	return &callbacks.AState{}
}

func (v *ValidInconsistentLabelInitEnv) New_B_Env() callbacks.B_Env {
	return &callbacks.BState{}
}

type ValidInconsistentLabelResultEnv struct {
}

func (v *ValidInconsistentLabelResultEnv) A_Result(result results.A_Result) {
}

func (v *ValidInconsistentLabelResultEnv) B_Result(result results.B_Result) {
}

func NewValidInconsistentLabelInitEnv() *ValidInconsistentLabelInitEnv {
	return &ValidInconsistentLabelInitEnv{}
}

func NewValidInconsistentLabelResultEnv() *ValidInconsistentLabelResultEnv {
	return &ValidInconsistentLabelResultEnv{}
}

func RunValidInconsistentLabel() {
	initEnv := NewValidInconsistentLabelInitEnv()
	resultEnv := NewValidInconsistentLabelResultEnv()
	validinconsistentlabel.Validinconsistentlabel(initEnv, resultEnv)
}
