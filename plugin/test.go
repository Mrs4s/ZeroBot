package plugin

import (
	"github.com/wdvxdr1123/ZeroBot"
)

func init() {
	a := testPlugin{}
	zero.RegisterPlugin(a) // 注册插件
}

type testPlugin struct{}

func (testPlugin) GetPluginInfo() zero.PluginInfo { // 返回插件信息
	return zero.PluginInfo{
		Author:     "wdvxdr1123",
		PluginName: "test",
		Version:    "0.1.0",
		Details:    "这是一个测试插件",
	}
}

func (testPlugin) Start() { // 插件主体
	zero.OnPrefix([]string{"复读", "echo", "fudu"}, zero.OnlyToMe).
		Got(
			"echo",
			"请输入复读内容",
			func(matcher *zero.Matcher, event zero.Event, state zero.State) zero.Response {
				zero.Send(event, matcher.State["echo"])
				return zero.SuccessResponse
			},
		)
}
