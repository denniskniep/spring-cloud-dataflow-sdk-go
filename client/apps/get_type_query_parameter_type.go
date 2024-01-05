package apps
import (
    "errors"
)
// 
type GetTypeQueryParameterType int

const (
    APP_GETTYPEQUERYPARAMETERTYPE GetTypeQueryParameterType = iota
    SOURCE_GETTYPEQUERYPARAMETERTYPE
    PROCESSOR_GETTYPEQUERYPARAMETERTYPE
    SINK_GETTYPEQUERYPARAMETERTYPE
    TASK_GETTYPEQUERYPARAMETERTYPE
)

func (i GetTypeQueryParameterType) String() string {
    return []string{"app", "source", "processor", "sink", "task"}[i]
}
func ParseGetTypeQueryParameterType(v string) (any, error) {
    result := APP_GETTYPEQUERYPARAMETERTYPE
    switch v {
        case "app":
            result = APP_GETTYPEQUERYPARAMETERTYPE
        case "source":
            result = SOURCE_GETTYPEQUERYPARAMETERTYPE
        case "processor":
            result = PROCESSOR_GETTYPEQUERYPARAMETERTYPE
        case "sink":
            result = SINK_GETTYPEQUERYPARAMETERTYPE
        case "task":
            result = TASK_GETTYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetTypeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetTypeQueryParameterType(values []GetTypeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTypeQueryParameterType) isMultiValue() bool {
    return false
}
