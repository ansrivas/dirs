//go:build windows

package dirs

import (
	"os"
	"path/filepath"
)

type windowsDirs struct{}

func NewDirs() Dirs {
	return &windowsDirs{}
}

func (d *windowsDirs) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (d *windowsDirs) CacheDir() (string, error) {
	return os.Getenv("LOCALAPPDATA"), nil
}

func (d *windowsDirs) ConfigDir() (string, error) {
	return os.Getenv("APPDATA"), nil
}

func (d *windowsDirs) DataDir() (string, error) {
	return os.Getenv("APPDATA"), nil
}

func (d *windowsDirs) DataLocalDir() (string, error) {
	return os.Getenv("LOCALAPPDATA"), nil
}

func (d *windowsDirs) ExecutableDir() (string, error) {
	return "", nil // Not standard on Windows
}

func (d *windowsDirs) PreferenceDir() (string, error) {
	return os.Getenv("APPDATA"), nil
}

func (d *windowsDirs) RuntimeDir() (string, error) {
	return "", nil // Not standard on Windows
}

func (d *windowsDirs) StateDir() (string, error) {
	return "", nil // Not standard on Windows
}

func (d *windowsDirs) AudioDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Music"), nil
}

func (d *windowsDirs) DesktopDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Desktop"), nil
}

func (d *windowsDirs) DocumentDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Documents"), nil
}

func (d *windowsDirs) DownloadDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Downloads"), nil
}

func (d *windowsDirs) FontDir() (string, error) {
	return "", nil // Not standard on Windows
}

func (d *windowsDirs) PictureDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Pictures"), nil
}

func (d *windowsDirs) PublicDir() (string, error) {
	return filepath.Join(os.Getenv("PUBLIC")), nil
}

func (d *windowsDirs) TemplateDir() (string, error) {
	return filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Templates"), nil
}

func (d *windowsDirs) VideoDir() (string, error) {
	return filepath.Join(os.Getenv("USERPROFILE"), "Videos"), nil
}
