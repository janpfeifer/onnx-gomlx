package onnxgomlx

import (
	. "github.com/gomlx/gomlx/graph"
	"github.com/gomlx/gomlx/ml/context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"onnx-gomlx/internal/protos"
	"os"
)

// Model represents a parsed ONNX file.
type Model struct {
	Proto protos.ModelProto
}

// ParseONNX parses an ONNX model into an internal representation that can be used to build a GoMLX graph.
func ParseONNX(contents []byte) (*Model, error) {
	m := &Model{}
	err := proto.Unmarshal(contents, &m.Proto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ONNX model proto")
	}
	return m, nil
}

// ReadONNXFile parses an ONNX model file into an internal representation that can be used to build a GoMLX graph.
// Notice any large constant is converted to variables.
func ReadONNXFile(filePath string) (*Model, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read ONNX model file in %s", filePath)
	}
	return ParseONNX(contents)
}

// Inputs returns a description of the inputs.
func (m *Model) Inputs() []string {
	return nil
}

// Outputs returns a description of the outputs.
func (m *Model) Outputs() []string {
	return nil
}

// Variables returns a description of the variables found in the model.
func (m *Model) Variables() []string {
	return nil
}

// BuildGraph that can be used both for inference and training.
// ctx can be set to nil if the model doesn't have any variables.
func (m *Model) BuildGraph(ctx *context.Context, inputs []*Node) (outputs []*Node) {
	return nil
}
