// Package conversion contains several conversion utilities between ONNX and GoMLX.
//
// TODO: For now mostly it includes ONNX -> GoMLX. Add the reverse direction translations.
package conversion

import (
	"github.com/gomlx/gopjrt/dtypes"
	"github.com/pkg/errors"
	"onnx-gomlx/internal/protos"
)

// DTypeOnnxToGomlx converts an ONNX data type to a gomlx data type.
func DTypeOnnxToGomlx(onnxDType protos.TensorProto_DataType) (dtypes.DType, error) {
	switch onnxDType {
	case protos.TensorProto_FLOAT:
		return dtypes.Float32, nil
	case protos.TensorProto_DOUBLE:
		return dtypes.Float64, nil
	case protos.TensorProto_INT32:
		return dtypes.Int32, nil
	case protos.TensorProto_INT64:
		return dtypes.Int64, nil
	case protos.TensorProto_UINT8:
		return dtypes.Uint8, nil
	case protos.TensorProto_INT8:
		return dtypes.Int8, nil
	case protos.TensorProto_INT16:
		return dtypes.Int16, nil
	case protos.TensorProto_UINT16:
		return dtypes.Uint16, nil
	case protos.TensorProto_UINT32:
		return dtypes.Uint32, nil
	case protos.TensorProto_UINT64:
		return dtypes.Uint64, nil
	case protos.TensorProto_BOOL:
		return dtypes.Bool, nil
	case protos.TensorProto_COMPLEX64:
		return dtypes.Complex64, nil
	case protos.TensorProto_COMPLEX128:
		return dtypes.Complex128, nil
	default:
		return dtypes.InvalidDType, errors.Errorf("unsupported/unknown ONNX data type %v", onnxDType)
	}
}
