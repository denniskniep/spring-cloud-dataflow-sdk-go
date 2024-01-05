package models
import (
    "errors"
)
// 
type JobExecutionThinResource_status int

const (
    COMPLETED_JOBEXECUTIONTHINRESOURCE_STATUS JobExecutionThinResource_status = iota
    STARTING_JOBEXECUTIONTHINRESOURCE_STATUS
    STARTED_JOBEXECUTIONTHINRESOURCE_STATUS
    STOPPING_JOBEXECUTIONTHINRESOURCE_STATUS
    STOPPED_JOBEXECUTIONTHINRESOURCE_STATUS
    FAILED_JOBEXECUTIONTHINRESOURCE_STATUS
    ABANDONED_JOBEXECUTIONTHINRESOURCE_STATUS
    UNKNOWN_JOBEXECUTIONTHINRESOURCE_STATUS
)

func (i JobExecutionThinResource_status) String() string {
    return []string{"COMPLETED", "STARTING", "STARTED", "STOPPING", "STOPPED", "FAILED", "ABANDONED", "UNKNOWN"}[i]
}
func ParseJobExecutionThinResource_status(v string) (any, error) {
    result := COMPLETED_JOBEXECUTIONTHINRESOURCE_STATUS
    switch v {
        case "COMPLETED":
            result = COMPLETED_JOBEXECUTIONTHINRESOURCE_STATUS
        case "STARTING":
            result = STARTING_JOBEXECUTIONTHINRESOURCE_STATUS
        case "STARTED":
            result = STARTED_JOBEXECUTIONTHINRESOURCE_STATUS
        case "STOPPING":
            result = STOPPING_JOBEXECUTIONTHINRESOURCE_STATUS
        case "STOPPED":
            result = STOPPED_JOBEXECUTIONTHINRESOURCE_STATUS
        case "FAILED":
            result = FAILED_JOBEXECUTIONTHINRESOURCE_STATUS
        case "ABANDONED":
            result = ABANDONED_JOBEXECUTIONTHINRESOURCE_STATUS
        case "UNKNOWN":
            result = UNKNOWN_JOBEXECUTIONTHINRESOURCE_STATUS
        default:
            return 0, errors.New("Unknown JobExecutionThinResource_status value: " + v)
    }
    return &result, nil
}
func SerializeJobExecutionThinResource_status(values []JobExecutionThinResource_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i JobExecutionThinResource_status) isMultiValue() bool {
    return false
}
