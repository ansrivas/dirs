//go:build windows

package dirs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Helper to check if a path is non-empty and the error is nil
func checkPath(t *testing.T, name string, path string, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("%s() returned an error: %v", name, err)
	}
	if path == "" {
		t.Errorf("%s() returned an empty path", name)
	}
	// Basic sanity check: ensure paths are absolute
	if !filepath.IsAbs(path) {
		t.Errorf("%s() returned a non-absolute path: %s", name, path)
	}
}

// Helper to check if a path is empty and the error is nil
func checkEmptyPath(t *testing.T, name string, path string, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("%s() returned an error: %v", name, err)
	}
	if path != "" {
		t.Errorf("%s() returned a non-empty path (%s), expected empty", name, path)
	}
}

func TestWindowsDirs(t *testing.T) {
	d := NewDirs()
	userProfile := os.Getenv("USERPROFILE")
	appData := os.Getenv("APPDATA")
	localAppData := os.Getenv("LOCALAPPDATA")
	public := os.Getenv("PUBLIC")

	if userProfile == "" || appData == "" || localAppData == "" || public == "" {
		t.Skip("Skipping some checks because standard environment variables (USERPROFILE, APPDATA, LOCALAPPDATA, PUBLIC) are not set.")
	}

	t.Run("HomeDir", func(t *testing.T) {
		path, err := d.HomeDir()
		checkPath(t, "HomeDir", path, err)
		if !strings.EqualFold(path, userProfile) { // Case-insensitive comparison for paths
			t.Errorf("HomeDir expected '%s', got '%s'", userProfile, path)
		}
	})

	t.Run("CacheDir", func(t *testing.T) {
		path, err := d.CacheDir()
		checkPath(t, "CacheDir", path, err)
		if !strings.EqualFold(path, localAppData) {
			t.Errorf("CacheDir expected '%s', got '%s'", localAppData, path)
		}
	})

	t.Run("ConfigDir", func(t *testing.T) {
		path, err := d.ConfigDir()
		checkPath(t, "ConfigDir", path, err)
		if !strings.EqualFold(path, appData) {
			t.Errorf("ConfigDir expected '%s', got '%s'", appData, path)
		}
	})

	t.Run("DataDir", func(t *testing.T) {
		path, err := d.DataDir()
		checkPath(t, "DataDir", path, err)
		if !strings.EqualFold(path, appData) {
			t.Errorf("DataDir expected '%s', got '%s'", appData, path)
		}
	})

	t.Run("DataLocalDir", func(t *testing.T) {
		path, err := d.DataLocalDir()
		checkPath(t, "DataLocalDir", path, err)
		if !strings.EqualFold(path, localAppData) {
			t.Errorf("DataLocalDir expected '%s', got '%s'", localAppData, path)
		}
	})

	t.Run("PreferenceDir", func(t *testing.T) {
		path, err := d.PreferenceDir()
		checkPath(t, "PreferenceDir", path, err)
		if !strings.EqualFold(path, appData) {
			t.Errorf("PreferenceDir expected '%s', got '%s'", appData, path)
		}
	})

	// Dirs expected to be empty on Windows
	t.Run("ExecutableDir", func(t *testing.T) {
		path, err := d.ExecutableDir()
		checkEmptyPath(t, "ExecutableDir", path, err)
	})
	t.Run("RuntimeDir", func(t *testing.T) {
		path, err := d.RuntimeDir()
		checkEmptyPath(t, "RuntimeDir", path, err)
	})
	t.Run("StateDir", func(t *testing.T) {
		path, err := d.StateDir()
		checkEmptyPath(t, "StateDir", path, err)
	})
	t.Run("FontDir", func(t *testing.T) {
		path, err := d.FontDir()
		checkEmptyPath(t, "FontDir", path, err)
	})

	// User Dirs - check they are under USERPROFILE or PUBLIC
	userDirs := map[string]func() (string, error){
		"AudioDir":    d.AudioDir,
		"DesktopDir":  d.DesktopDir,
		"DocumentDir": d.DocumentDir,
		"DownloadDir": d.DownloadDir,
		"PictureDir":  d.PictureDir,
		"VideoDir":    d.VideoDir,
	}
	for name, getter := range userDirs {
		t.Run(name, func(t *testing.T) {
			path, err := getter()
			checkPath(t, name, path, err)
			if !strings.HasPrefix(strings.ToLower(path), strings.ToLower(userProfile)) {
				t.Errorf("%s path '%s' does not start with USERPROFILE '%s'", name, path, userProfile)
			}
		})
	}

	// Special cases
	t.Run("PublicDir", func(t *testing.T) {
		path, err := d.PublicDir()
		checkPath(t, "PublicDir", path, err)
		if !strings.EqualFold(path, public) {
			t.Errorf("PublicDir expected '%s', got '%s'", public, path)
		}
	})

	t.Run("TemplateDir", func(t *testing.T) {
		path, err := d.TemplateDir()
		checkPath(t, "TemplateDir", path, err)
		expectedSuffix := `Microsoft\Windows\Templates`
		if !strings.HasSuffix(strings.ToLower(path), strings.ToLower(expectedSuffix)) {
			t.Errorf("TemplateDir path '%s' does not end with expected suffix '%s'", path, expectedSuffix)
		}
		if !strings.HasPrefix(strings.ToLower(path), strings.ToLower(appData)) {
			t.Errorf("TemplateDir path '%s' does not start with APPDATA '%s'", path, appData)
		}
	})
}
