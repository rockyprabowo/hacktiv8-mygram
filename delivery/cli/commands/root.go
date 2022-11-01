package commands

import (
	_ "embed"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

//go:embed ascii_banner.txt
var ASCIIBanner string
var executableName, _ = os.Executable()

var RootCmd = &cobra.Command{
	Use:   filepath.Base(executableName),
	Short: "MyGram",
	Long: ASCIIBanner + "\n" +
		"MyGram - Backend CLI",
}
