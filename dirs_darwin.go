//go:build darwin

package dirs

import (
	"os"
	"path/filepath"
)

type darwinDirs struct{}

func NewDirs() Dirs {
	return &darwinDirs{}
}

func (d *darwinDirs) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (d *darwinDirs) CacheDir() (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Caches"), nil
}

func (d *darwinDirs) ConfigDir() (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	// Corresponds to Application Support directory on macOS
	return filepath.Join(home, "Library", "Application Support"), nil
}

func (d *darwinDirs) DataDir() (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	// Corresponds to Application Support directory on macOS
	return filepath.Join(home, "Library", "Application Support"), nil
}

func (d *darwinDirs) DataLocalDir() (string, error) {
	// macOS doesn't typically distinguish between roaming and local data in the same way Windows does.
	// Use the standard Application Support directory.
	return d.DataDir()
}

func (d *darwinDirs) ExecutableDir() (string, error) {
	// No standard equivalent specified in the reference doc for macOS.
	return "", nil
}

func (d *darwinDirs) PreferenceDir() (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Preferences"), nil
}

func (d *darwinDirs) RuntimeDir() (string, error) {
	// No standard equivalent specified in the reference doc for macOS.
	return "", nil
}

func (d *darwinDirs) StateDir() (string, error) {
	// No standard equivalent specified in the reference doc for macOS.
	// Often, state might go into Application Support or Caches depending on volatility.
	// Returning empty as per the pattern for unspecified dirs.
	return "", nil
}

// Helper for standard user directories under $HOME
func (d *darwinDirs) getUserHomeSubDir(subPath string) (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, subPath), nil
}

func (d *darwinDirs) AudioDir() (string, error) {
	return d.getUserHomeSubDir("Music")
}

func (d *darwinDirs) DesktopDir() (string, error) {
	return d.getUserHomeSubDir("Desktop")
}

func (d *darwinDirs) DocumentDir() (string, error) {
	return d.getUserHomeSubDir("Documents")
}

func (d *darwinDirs) DownloadDir() (string, error) {
	return d.getUserHomeSubDir("Downloads")
}

func (d *darwinDirs) FontDir() (string, error) {
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Fonts"), nil
}

func (d *darwinDirs) PictureDir() (string, error) {
	return d.getUserHomeSubDir("Pictures")
}

func (d *darwinDirs) PublicDir() (string, error) {
	return d.getUserHomeSubDir("Public")
}

func (d *darwinDirs) TemplateDir() (string, error) {
	// No standard equivalent specified in the reference doc for macOS.
	return "", nil
}

func (d *darwinDirs) VideoDir() (string, error) {
	// Note: macOS standard is "Movies", not "Videos"
	return d.getUserHomeSubDir("Movies")
}
