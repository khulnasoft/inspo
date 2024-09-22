package runtime

import (
	"github.com/spf13/viper"

	"github.com/khulnasoft/inspo/inspo"
)

type Options struct {
	Ci           bool
	Image        string
	Source       inspo.ImageSource
	IgnoreErrors bool
	ExportFile   string
	CiConfig     *viper.Viper
	BuildArgs    []string
}
