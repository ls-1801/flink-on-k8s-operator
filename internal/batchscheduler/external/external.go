package external

import (
	schedulerinterface "github.com/spotify/flink-on-k8s-operator/internal/batchscheduler/types"
	"github.com/spotify/flink-on-k8s-operator/internal/model"
	ctrl "sigs.k8s.io/controller-runtime"
)

// external scheduler implements the BatchScheduler interface.
type ExternalBatchScheduler struct {
}

func (e ExternalBatchScheduler) Name() string {
	return "external"
}

func (e ExternalBatchScheduler) Schedule(options schedulerinterface.SchedulerOptions, desired *model.DesiredClusterState) error {
	ctrl.Log.Info("External Scheduler scheduling")
	if desired == nil || desired.TmStatefulSet == nil {
		ctrl.Log.Info("Spec template is nil")
		return nil
	}

	desired.TmStatefulSet.Spec.Template.Spec.SchedulerName = "external-scheduler"
	desired.TmStatefulSet.Spec.Template.Spec.PriorityClassName = "high-priority"

	return nil
}

// Create external BatchScheduler
func New() (schedulerinterface.BatchScheduler, error) {
	ctrl.Log.Info("External Scheduler Initialized")
	return &ExternalBatchScheduler{}, nil
}
