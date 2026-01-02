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
It may require administrator privileges to write to '/usr/local/bin'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting setup for dev-pal...")

		destDir := "/usr/local/bin"
		destName := "dev-pal"
		destPath := filepath.Join(destDir, destName)

		// 1. Get current executable path
		sourcePath, err := os.Executable()
		if err != nil {
			fmt.Printf("❌ Error: Could not find path of current executable: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ℹ️  Current binary path: %s\n", sourcePath)

		// 2. Check for write permissions
		if err := checkWritePerms(destDir); err != nil {
			fmt.Printf("❌ Error: Write permission denied for %s.\n", destDir)
			fmt.Println("Please try running this command again with sudo:")
			// Attempt to find the original command used to run the program
			originalCommand := "dev-pal"
			if len(os.Args) > 0 {
				originalCommand = os.Args[0]
			}
			fmt.Printf("sudo %s setup\n", originalCommand)
			os.Exit(1)
		}
		fmt.Printf("✅ Write permissions verified for %s\n", destDir)

		// 3. Copy the binary
		fmt.Printf("Installing dev-pal to %s...\n", destPath)
		if err := copyFile(sourcePath, destPath); err != nil {
			fmt.Printf("❌ Error: Failed to copy binary: %v\n", err)
			os.Exit(1)
		}

		// 4. Make the destination binary executable
		if err := os.Chmod(destPath, 0755); err != nil {
			fmt.Printf("❌ Error: Failed to set executable permissions on %s: %v\n", destPath, err)
			// Attempt to clean up
			os.Remove(destPath)
			os.Exit(1)
		}

		fmt.Println("\n🎉 dev-pal setup is complete!")
		fmt.Println("You can now use the 'dev-pal' command from any directory.")
	},
}

// checkWritePerms tries to create a temporary file in a directory to check for write permissions.
func checkWritePerms(path string) error {
	tmpFile := filepath.Join(path, "dev-pal.tmp")
	if err := os.WriteFile(tmpFile, []byte("test"), 0666); err != nil {
		return err
	}
	return os.Remove(tmpFile)
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
