//go:build linux

package dirs

import (
	"os"
	"path/filepath"
)

type linuxDirs struct{}

func NewDirs() Dirs {
	return &linuxDirs{}
}

func (d *linuxDirs) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (d *linuxDirs) CacheDir() (string, error) {
	if dir := os.Getenv("XDG_CACHE_HOME"); dir != "" {
		return dir, nil
	}
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache"), nil
}

func (d *linuxDirs) ConfigDir() (string, error) {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return dir, nil
	}
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config"), nil
}

func (d *linuxDirs) DataDir() (string, error) {
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return dir, nil
	}
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share"), nil
}

func (d *linuxDirs) DataLocalDir() (string, error) {
	// XDG Base Directory Spec suggests local data is the same as data.
	return d.DataDir()
}

func (d *linuxDirs) ExecutableDir() (string, error) {
	if dir := os.Getenv("XDG_BIN_HOME"); dir != "" {
		return dir, nil
	}
	// Fallback based on XDG spec recommendation: $HOME/.local/bin
	home, err := d.HomeDir()
	if err != nil {
		// Cannot determine fallback without home dir
		return "", err
	}
	return filepath.Join(home, ".local", "bin"), nil
}

func (d *linuxDirs) PreferenceDir() (string, error) {
	// Preferences are typically stored in the config directory on Linux.
	return d.ConfigDir()
}

func (d *linuxDirs) RuntimeDir() (string, error) {
	dir := os.Getenv("XDG_RUNTIME_DIR")
	// Runtime dir might not be set or available, return empty string if so,
	// as per the spec (it's optional). Error is not appropriate here.
	return dir, nil
}

func (d *linuxDirs) StateDir() (string, error) {
	if dir := os.Getenv("XDG_STATE_HOME"); dir != "" {
		return dir, nil
	}
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "state"), nil
}

// getUserDir checks the XDG environment variable and falls back to a default path.
func (d *linuxDirs) getUserDir(envVar, defaultSubPath string) (string, error) {
	if dir := os.Getenv(envVar); dir != "" {
		return dir, nil
	}
	home, err := d.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, defaultSubPath), nil
}

func (d *linuxDirs) AudioDir() (string, error) {
	return d.getUserDir("XDG_MUSIC_DIR", "Music")
}

func (d *linuxDirs) DesktopDir() (string, error) {
	return d.getUserDir("XDG_DESKTOP_DIR", "Desktop")
}

func (d *linuxDirs) DocumentDir() (string, error) {
	return d.getUserDir("XDG_DOCUMENTS_DIR", "Documents")
}

func (d *linuxDirs) DownloadDir() (string, error) {
	return d.getUserDir("XDG_DOWNLOAD_DIR", "Downloads")
}

func (d *linuxDirs) FontDir() (string, error) {
	// Fonts are typically placed in the data directory according to XDG spec.
	dataDir, err := d.DataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "fonts"), nil
}

func (d *linuxDirs) PictureDir() (string, error) {
	return d.getUserDir("XDG_PICTURES_DIR", "Pictures")
}

func (d *linuxDirs) PublicDir() (string, error) {
	return d.getUserDir("XDG_PUBLICSHARE_DIR", "Public")
}

func (d *linuxDirs) TemplateDir() (string, error) {
	return d.getUserDir("XDG_TEMPLATES_DIR", "Templates")
}

func (d *linuxDirs) VideoDir() (string, error) {
	return d.getUserDir("XDG_VIDEOS_DIR", "Videos")
}
