package cmd

import (
	"fmt"
	"github.com/khulnasoft/inspo/inspo"
	"os"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/inspo/runtime"
)

// doAnalyzeCmd takes a docker image tag, digest, or id and displays the
// image analysis to the screen
func doAnalyzeCmd(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		printVersionFlag, err := cmd.PersistentFlags().GetBool("version")
		if err == nil && printVersionFlag {
			printVersion(cmd, args)
			return
		}

		fmt.Println("No image argument given")
		os.Exit(1)
	}

	userImage := args[0]
	if userImage == "" {
		fmt.Println("No image argument given")
		os.Exit(1)
	}

	initLogging()

	isCi, ciConfig, err := configureCi()

	if err != nil {
		fmt.Printf("ci configuration error: %v\n", err)
		os.Exit(1)
	}

	var sourceType inspo.ImageSource
	var imageStr string

	sourceType, imageStr = inspo.DeriveImageSource(userImage)

	if sourceType == inspo.SourceUnknown {
		sourceStr, err := cmd.PersistentFlags().GetString("source")
		if err != nil {
			fmt.Printf("unable to determine image source: %v\n", err)
			os.Exit(1)
		}

		sourceType = inspo.ParseImageSource(sourceStr)
		imageStr = userImage
	}

	runtime.Run(runtime.Options{
		Ci:         isCi,
		Source:     sourceType,
		Image:      imageStr,
		ExportFile: exportFile,
		CiConfig:   ciConfig,
	})
}