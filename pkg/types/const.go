//go:build windows

package types

const (
	DefaultCollectors            = "cpu,cpu_info,cs,logical_disk,memory,net,os,remote_fx,process,service,tcp"
	DefaultCollectorsPlaceholder = "[defaults]"
	Namespace                    = "windows"
)
