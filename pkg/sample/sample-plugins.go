package sample

import (
	"context"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type SamplePlugin struct{}

var _ framework.FilterPlugin = &SamplePlugin{}

const (
	Name = "SamplePlugin"
)

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &SamplePlugin{}, nil
}

func (pl *SamplePlugin) Name() string {
	return Name
}

func (pl *SamplePlugin) Filter(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	if !strings.Contains(pod.Name, "sample") {
		return framework.NewStatus(framework.Error, "Pod name does not contain 'sample'")
	}

	return nil
}