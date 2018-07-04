package main

import (
	ri "github.com/dotSlashLu/nightswatch/agent/raven_interface"
	"time"
)

const pluginName = "libvirt"

func Report(ch chan *ri.PluginReport, errCh chan *error) {
	plugin := ri.NewPlugin(pluginName, ch, errCh)
	for {
		intervalReport(plugin)
	}
}

func intervalReport(plugin *ri.Plugin) {
	ticker := time.NewTicker(3 * time.Second)
	_ = <-ticker.C
	plugin.SingleReport(ri.ReportValInt, "hb", 1)
}
