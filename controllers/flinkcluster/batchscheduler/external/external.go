package external

import (
	"github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
	schedulerinterface "github.com/spotify/flink-on-k8s-operator/controllers/flinkcluster/batchscheduler/interface"
	"github.com/spotify/flink-on-k8s-operator/controllers/flinkcluster/model"
	ctrl "sigs.k8s.io/controller-runtime"
)

// volcano scheduler implements the BatchScheduler interface.
type ExternalBatchScheduler struct {
}

func (e ExternalBatchScheduler) Name() string {
	return "external"
}

func (e ExternalBatchScheduler) Schedule(cluster *v1beta1.FlinkCluster, desired *model.DesiredClusterState) error {
	ctrl.Log.Info("External Scheduler scheduling")
	if desired == nil || desired.TmStatefulSet == nil {
		ctrl.Log.Info("Spec template is nil")
		return nil
	}

	desired.TmStatefulSet.Spec.Template.Spec.SchedulerName = "my-scheduler"
	desired.TmStatefulSet.Spec.Template.Spec.PriorityClassName = "high-priority"

	return nil
}

// Create volcano BatchScheduler
func New() (schedulerinterface.BatchScheduler, error) {
	ctrl.Log.Info("External Scheduler Initialized")
	return &ExternalBatchScheduler{}, nil
}
