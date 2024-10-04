package simplegetter

import (
	"context"
	"testing"
)

func TestGetter(t *testing.T) {
	c := Client{
		Ctx: context.Background(),
		Src: "https://github.com/onnx/models/raw/refs/heads/main/validated/vision/style_transfer/fast_neural_style/model/udnie-8.onnx",
		Dst: "./udnie-8.onnx",
	}

	err := c.Get()
	if err != nil {
		t.Fatal(err)
	}
}
