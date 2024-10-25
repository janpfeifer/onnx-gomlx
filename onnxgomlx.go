package onnxgomlx

import (
	. "github.com/gomlx/gomlx/graph"
	"github.com/gomlx/gomlx/ml/context"
	"github.com/pkg/errors"
)

// Model represents a parsed ONNX file.
type Model struct {
	nodes []string
}

// ReadFile parses an ONNX model file into an internal representation that can be used to build a GoMLX graph.
// Notice any large constant is converted to variables.
func ReadFile(filePath string) (*Model, error) {
	return nil, errors.New("Not implemented")
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
