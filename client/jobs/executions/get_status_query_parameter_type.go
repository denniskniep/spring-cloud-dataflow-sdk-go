package executions
import (
    "errors"
)
// 
type GetStatusQueryParameterType int

const (
    COMPLETED_GETSTATUSQUERYPARAMETERTYPE GetStatusQueryParameterType = iota
    STARTING_GETSTATUSQUERYPARAMETERTYPE
    STARTED_GETSTATUSQUERYPARAMETERTYPE
    STOPPING_GETSTATUSQUERYPARAMETERTYPE
    STOPPED_GETSTATUSQUERYPARAMETERTYPE
    FAILED_GETSTATUSQUERYPARAMETERTYPE
    ABANDONED_GETSTATUSQUERYPARAMETERTYPE
    UNKNOWN_GETSTATUSQUERYPARAMETERTYPE
)

func (i GetStatusQueryParameterType) String() string {
    return []string{"COMPLETED", "STARTING", "STARTED", "STOPPING", "STOPPED", "FAILED", "ABANDONED", "UNKNOWN"}[i]
}
func ParseGetStatusQueryParameterType(v string) (any, error) {
    result := COMPLETED_GETSTATUSQUERYPARAMETERTYPE
    switch v {
        case "COMPLETED":
            result = COMPLETED_GETSTATUSQUERYPARAMETERTYPE
        case "STARTING":
            result = STARTING_GETSTATUSQUERYPARAMETERTYPE
        case "STARTED":
            result = STARTED_GETSTATUSQUERYPARAMETERTYPE
        case "STOPPING":
            result = STOPPING_GETSTATUSQUERYPARAMETERTYPE
        case "STOPPED":
            result = STOPPED_GETSTATUSQUERYPARAMETERTYPE
        case "FAILED":
            result = FAILED_GETSTATUSQUERYPARAMETERTYPE
        case "ABANDONED":
            result = ABANDONED_GETSTATUSQUERYPARAMETERTYPE
        case "UNKNOWN":
            result = UNKNOWN_GETSTATUSQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetStatusQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetStatusQueryParameterType(values []GetStatusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStatusQueryParameterType) isMultiValue() bool {
    return false
}
