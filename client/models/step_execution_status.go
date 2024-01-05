package models
import (
    "errors"
)
// 
type StepExecution_status int

const (
    COMPLETED_STEPEXECUTION_STATUS StepExecution_status = iota
    STARTING_STEPEXECUTION_STATUS
    STARTED_STEPEXECUTION_STATUS
    STOPPING_STEPEXECUTION_STATUS
    STOPPED_STEPEXECUTION_STATUS
    FAILED_STEPEXECUTION_STATUS
    ABANDONED_STEPEXECUTION_STATUS
    UNKNOWN_STEPEXECUTION_STATUS
)

func (i StepExecution_status) String() string {
    return []string{"COMPLETED", "STARTING", "STARTED", "STOPPING", "STOPPED", "FAILED", "ABANDONED", "UNKNOWN"}[i]
}
func ParseStepExecution_status(v string) (any, error) {
    result := COMPLETED_STEPEXECUTION_STATUS
    switch v {
        case "COMPLETED":
            result = COMPLETED_STEPEXECUTION_STATUS
        case "STARTING":
            result = STARTING_STEPEXECUTION_STATUS
        case "STARTED":
            result = STARTED_STEPEXECUTION_STATUS
        case "STOPPING":
            result = STOPPING_STEPEXECUTION_STATUS
        case "STOPPED":
            result = STOPPED_STEPEXECUTION_STATUS
        case "FAILED":
            result = FAILED_STEPEXECUTION_STATUS
        case "ABANDONED":
            result = ABANDONED_STEPEXECUTION_STATUS
        case "UNKNOWN":
            result = UNKNOWN_STEPEXECUTION_STATUS
        default:
            return 0, errors.New("Unknown StepExecution_status value: " + v)
    }
    return &result, nil
}
func SerializeStepExecution_status(values []StepExecution_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StepExecution_status) isMultiValue() bool {
    return false
}
