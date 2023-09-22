package coalminePlugin

import (
	"cts/plugin"
	"cts/utils/http"
	"fmt"
)

type CoalmineParam struct {
	Company string
	Biz     string
	Channel string
	Pdt     string
}

type CoalmineExecutor struct {
	FuncList   []func(params plugin.IPluginParams)
	FuncParams []plugin.IPluginParams
}

func (e *CoalmineExecutor) AddFunc(f func(params plugin.IPluginParams), params plugin.IPluginParams) {
	fmt.Println("append", f, params, e.FuncList)
	e.FuncList = append(e.FuncList, f)
	e.FuncParams = append(e.FuncParams, params)
}

func (e *CoalmineExecutor) RunFunc() {
	for i, f := range e.FuncList {
		fmt.Println("run func: ", i, f)
		f(e.FuncParams[i])
	}
}

func CoalmineFunc(p plugin.IPluginParams) {
	cc := http.CoalmineClient{}
	cp, ok := p.(CoalmineParam)
	fmt.Println(cp, ok)
	executor := CoalmineExecutor{
		FuncList:   make([]func(params plugin.IPluginParams), 0),
		FuncParams: make([]plugin.IPluginParams, 0),
	}
	fmt.Println("executor add step")
	executor.AddFunc(cc.Step1, cp)
	executor.AddFunc(cc.Step2, cp)
	executor.AddFunc(cc.Step3, cp)
	executor.RunFunc()
}
