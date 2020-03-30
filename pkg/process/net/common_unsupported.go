// +build !linux,!windows

package net

import (
	model "github.com/DataDog/agent-payload/process"
	"github.com/DataDog/datadog-agent/pkg/ebpf"
	"github.com/DataDog/datadog-agent/pkg/process/config"
)

// RemoteSysProbeUtil is not supported
type RemoteSysProbeUtil struct{}

// SetSystemProbePath is not supported
func SetSystemProbePath(_ *config.AgentConfig) {
	// no-op
}

// GetRemoteSystemProbeUtil is not supported
func GetRemoteSystemProbeUtil() (*RemoteSysProbeUtil, error) {
	return &RemoteSysProbeUtil{}, nil
}

// GetConnections is not supported
func (r *RemoteSysProbeUtil) GetConnections(clientID string) (*model.Connections, error) {
	return nil, ebpf.ErrNotImplemented
}

// ShouldLogTracerUtilError is not supported
func ShouldLogTracerUtilError() bool {
	return false
}