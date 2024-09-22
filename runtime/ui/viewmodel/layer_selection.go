package viewmodel

import (
	"github.com/khulnasoft/inspo/inspo/image"
)

type LayerSelection struct {
	Layer                                                      *image.Layer
	BottomTreeStart, BottomTreeStop, TopTreeStart, TopTreeStop int
}
