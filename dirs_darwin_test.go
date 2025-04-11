//go:build darwin

package dirs

import (
	"path/filepath"
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

func TestDarwinDirs(t *testing.T) {
	d := NewDirs()
	home, err := d.HomeDir()
	if err != nil {
		t.Fatalf("Failed to get HomeDir for setup: %v", err)
	}

	t.Run("HomeDir", func(t *testing.T) {
		// Already fetched home above, just check basic validity
		checkPath(t, "HomeDir", home, err)
		if home == "" {
			t.Fatal("Home directory is empty, cannot proceed with relative tests")
		}
	})

	t.Run("CacheDir", func(t *testing.T) {
		path, err := d.CacheDir()
		checkPath(t, "CacheDir", path, err)
		expected := filepath.Join(home, "Library", "Caches")
		if path != expected {
			t.Errorf("CacheDir expected '%s', got '%s'", expected, path)
		}
	})

	t.Run("ConfigDir", func(t *testing.T) {
		path, err := d.ConfigDir()
		checkPath(t, "ConfigDir", path, err)
		expected := filepath.Join(home, "Library", "Application Support")
		if path != expected {
			t.Errorf("ConfigDir expected '%s', got '%s'", expected, path)
		}
	})

	t.Run("DataDir", func(t *testing.T) {
		path, err := d.DataDir()
		checkPath(t, "DataDir", path, err)
		expected := filepath.Join(home, "Library", "Application Support")
		if path != expected {
			t.Errorf("DataDir expected '%s', got '%s'", expected, path)
		}
	})

	t.Run("DataLocalDir", func(t *testing.T) {
		// Should be same as DataDir on macOS
		path, err := d.DataLocalDir()
		checkPath(t, "DataLocalDir", path, err)
		dataPath, _ := d.DataDir()
		if path != dataPath {
			t.Errorf("DataLocalDir (%s) should be the same as DataDir (%s) on macOS", path, dataPath)
		}
	})

	t.Run("PreferenceDir", func(t *testing.T) {
		path, err := d.PreferenceDir()
		checkPath(t, "PreferenceDir", path, err)
		expected := filepath.Join(home, "Library", "Preferences")
		if path != expected {
			t.Errorf("PreferenceDir expected '%s', got '%s'", expected, path)
		}
	})

	t.Run("FontDir", func(t *testing.T) {
		path, err := d.FontDir()
		checkPath(t, "FontDir", path, err)
		expected := filepath.Join(home, "Library", "Fonts")
		if path != expected {
			t.Errorf("FontDir expected '%s', got '%s'", expected, path)
		}
	})

	// Dirs expected to be empty on macOS
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
	t.Run("TemplateDir", func(t *testing.T) {
		path, err := d.TemplateDir()
		checkEmptyPath(t, "TemplateDir", path, err)
	})

	// User Dirs - check they are under HOME
	userDirs := map[string]struct {
		getter  func() (string, error)
		subPath string
	}{
		"AudioDir":    {d.AudioDir, "Music"},
		"DesktopDir":  {d.DesktopDir, "Desktop"},
		"DocumentDir": {d.DocumentDir, "Documents"},
		"DownloadDir": {d.DownloadDir, "Downloads"},
		"PictureDir":  {d.PictureDir, "Pictures"},
		"PublicDir":   {d.PublicDir, "Public"},
		"VideoDir":    {d.VideoDir, "Movies"}, // Note: Movies on macOS
	}
	for name, data := range userDirs {
		t.Run(name, func(t *testing.T) {
			path, err := data.getter()
			checkPath(t, name, path, err)
			expected := filepath.Join(home, data.subPath)
			if path != expected {
				t.Errorf("%s expected '%s', got '%s'", name, expected, path)
			}
		})
	}
}
