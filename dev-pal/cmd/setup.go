package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Performs one-time setup for dev-pal.",
	Long: `This command installs the dev-pal binary to your system's PATH for global access.
It requires administrator privileges to write to '/usr/local/bin'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting setup for dev-pal...")

		destDir := "/usr/local/bin"
		destName := "dev-pal"
		destPath := filepath.Join(destDir, destName)

		sourcePath, err := os.Executable()
		if err != nil {
			fmt.Printf("❌ Error: Could not find path of current executable: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Installing dev-pal to %s...\n", destPath)
		if err := copyFile(sourcePath, destPath); err != nil {
			fmt.Printf("❌ Error: Failed to copy binary: %v\n", err)
			fmt.Println("   This usually requires administrator privileges. Please ensure the script is run correctly.")
			os.Exit(1)
		}

		if err := os.Chmod(destPath, 0755); err != nil {
			fmt.Printf("❌ Error: Failed to set executable permissions on %s: %v\n", destPath, err)
			os.Remove(destPath)
			os.Exit(1)
		}

		fmt.Println("\n🎉 dev-pal setup is complete!")
		fmt.Println("You can now use the 'dev-pal' command from any directory.")
	},
}

// copyFile copies a file from a source path to a destination path.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create a new file, truncating it if it exists
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure all data is written to stable storage
	return destFile.Sync()
}


func init() {
	rootCmd.AddCommand(setupCmd)
}
