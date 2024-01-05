package models
import (
    "errors"
)
// 
type JobExecution_status int

const (
    COMPLETED_JOBEXECUTION_STATUS JobExecution_status = iota
    STARTING_JOBEXECUTION_STATUS
    STARTED_JOBEXECUTION_STATUS
    STOPPING_JOBEXECUTION_STATUS
    STOPPED_JOBEXECUTION_STATUS
    FAILED_JOBEXECUTION_STATUS
    ABANDONED_JOBEXECUTION_STATUS
    UNKNOWN_JOBEXECUTION_STATUS
)

func (i JobExecution_status) String() string {
    return []string{"COMPLETED", "STARTING", "STARTED", "STOPPING", "STOPPED", "FAILED", "ABANDONED", "UNKNOWN"}[i]
}
func ParseJobExecution_status(v string) (any, error) {
    result := COMPLETED_JOBEXECUTION_STATUS
    switch v {
        case "COMPLETED":
            result = COMPLETED_JOBEXECUTION_STATUS
        case "STARTING":
            result = STARTING_JOBEXECUTION_STATUS
        case "STARTED":
            result = STARTED_JOBEXECUTION_STATUS
        case "STOPPING":
            result = STOPPING_JOBEXECUTION_STATUS
        case "STOPPED":
            result = STOPPED_JOBEXECUTION_STATUS
        case "FAILED":
            result = FAILED_JOBEXECUTION_STATUS
        case "ABANDONED":
            result = ABANDONED_JOBEXECUTION_STATUS
        case "UNKNOWN":
            result = UNKNOWN_JOBEXECUTION_STATUS
        default:
            return 0, errors.New("Unknown JobExecution_status value: " + v)
    }
    return &result, nil
}
func SerializeJobExecution_status(values []JobExecution_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i JobExecution_status) isMultiValue() bool {
    return false
}
