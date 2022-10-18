package convert

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StringToStringWrapper(val string) *wrappers.StringValue {
	if val == "" {
		return nil
	}
	return &wrappers.StringValue{Value: val}
}

func StringToPtr(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

func StringWrapperToPtr(val *wrappers.StringValue) *string {
	if val == nil || val.GetValue() == "" {
		return nil
	}
	return &val.Value
}

func PtrToStringWrapper(val *string) *wrappers.StringValue {
	if val == nil {
		return nil
	}
	return &wrappers.StringValue{Value: *val}
}

func BoolWrapperToPtr(val *wrappers.BoolValue) *bool {
	if val == nil {
		return nil
	}
	return &val.Value
}

func PtrToBoolWrapper(val *bool) *wrappers.BoolValue {
	if val == nil {
		return nil
	}
	return &wrappers.BoolValue{Value: *val}
}

func BoolToBoolWrapper(val bool) *wrappers.BoolValue {
	return &wrappers.BoolValue{Value: val}
}

func Int32ToInt32Wrapper(val int32) *wrappers.Int32Value {
	return &wrappers.Int32Value{Value: val}
}

func UInt32ToInt32Wrapper(val uint32) *wrappers.UInt32Value {
	return &wrappers.UInt32Value{Value: val}
}

func Int64ToInt64Wrapper(val int64) *wrappers.Int64Value {
	return &wrappers.Int64Value{Value: val}
}

func PtrToInt64Wrapper(val *int64) *wrappers.Int64Value {
	if val == nil {
		return nil
	}
	return &wrappers.Int64Value{Value: *val}
}

func Int64WrapperToPtr(val *wrappers.Int64Value) *int64 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func PtrToBool(val *bool) bool {
	if val == nil {
		return false
	}
	return *val
}

func PtrToString(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func Float64WrapperToPtr(val *wrappers.DoubleValue) *float64 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func PtrToDoubleWrapper(val *float64) *wrappers.DoubleValue {
	if val == nil {
		return nil
	}
	return &wrappers.DoubleValue{Value: *val}
}

func Float32ToFloatWrapper(val float32) *wrappers.FloatValue {
	return &wrappers.FloatValue{Value: val}
}

func Float64ToDoubleWrapper(val float64) *wrappers.DoubleValue {
	return &wrappers.DoubleValue{Value: val}
}

func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	return timestamppb.New(t)
}

func PtrToTimestamp(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func TimestampToTime(t *timestamp.Timestamp) time.Time {
	return t.AsTime()
}

func TimestampToPtr(t *timestamp.Timestamp) *time.Time {
	if t == nil || t.CheckValid() != nil {
		return nil
	}
	res := t.AsTime()
	return &res
}
