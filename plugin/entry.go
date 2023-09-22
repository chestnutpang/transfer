package plugin

type IPluginParams interface{}

type FPlugin func(params *IPluginParams)

type IPluginExecutor interface {
	AddFunc(f FPlugin, params *IPluginParams)
	RunFunc()
}
