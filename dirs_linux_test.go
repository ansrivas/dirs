//go:build linux

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
		// Allow empty paths for specific cases like RuntimeDir if env var is not set
		if name != "RuntimeDir" || os.Getenv("XDG_RUNTIME_DIR") != "" {
			t.Errorf("%s() returned an empty path", name)
		}
	}
	// Basic sanity check: ensure paths are absolute unless it's an intentionally empty result
	if path != "" && !filepath.IsAbs(path) {
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

func TestLinuxDirs(t *testing.T) {
	d := NewDirs()

	t.Run("HomeDir", func(t *testing.T) {
		path, err := d.HomeDir()
		checkPath(t, "HomeDir", path, err)
	})

	t.Run("CacheDir", func(t *testing.T) {
		path, err := d.CacheDir()
		checkPath(t, "CacheDir", path, err)
		if os.Getenv("XDG_CACHE_HOME") == "" && !strings.HasSuffix(path, ".cache") {
			t.Errorf("Default CacheDir path should end with .cache, got: %s", path)
		}
	})

	t.Run("ConfigDir", func(t *testing.T) {
		path, err := d.ConfigDir()
		checkPath(t, "ConfigDir", path, err)
		if os.Getenv("XDG_CONFIG_HOME") == "" && !strings.HasSuffix(path, ".config") {
			t.Errorf("Default ConfigDir path should end with .config, got: %s", path)
		}
	})

	t.Run("DataDir", func(t *testing.T) {
		path, err := d.DataDir()
		checkPath(t, "DataDir", path, err)
		if os.Getenv("XDG_DATA_HOME") == "" && !strings.HasSuffix(path, ".local/share") {
			t.Errorf("Default DataDir path should end with .local/share, got: %s", path)
		}
	})

	t.Run("DataLocalDir", func(t *testing.T) {
		// Should be same as DataDir on Linux
		path, err := d.DataLocalDir()
		checkPath(t, "DataLocalDir", path, err)
		dataPath, _ := d.DataDir()
		if path != dataPath {
			t.Errorf("DataLocalDir (%s) should be the same as DataDir (%s) on Linux", path, dataPath)
		}
	})

	t.Run("ExecutableDir", func(t *testing.T) {
		path, err := d.ExecutableDir()
		checkPath(t, "ExecutableDir", path, err)
		if os.Getenv("XDG_BIN_HOME") == "" && !strings.HasSuffix(path, ".local/bin") {
			t.Errorf("Default ExecutableDir path should end with .local/bin, got: %s", path)
		}
	})

	t.Run("PreferenceDir", func(t *testing.T) {
		// Should be same as ConfigDir on Linux
		path, err := d.PreferenceDir()
		checkPath(t, "PreferenceDir", path, err)
		configPath, _ := d.ConfigDir()
		if path != configPath {
			t.Errorf("PreferenceDir (%s) should be the same as ConfigDir (%s) on Linux", path, configPath)
		}
	})

	t.Run("RuntimeDir", func(t *testing.T) {
		// This might be empty if XDG_RUNTIME_DIR is not set, which is valid.
		path, err := d.RuntimeDir()
		checkPath(t, "RuntimeDir", path, err) // checkPath allows empty for RuntimeDir
	})

	t.Run("StateDir", func(t *testing.T) {
		path, err := d.StateDir()
		checkPath(t, "StateDir", path, err)
		if os.Getenv("XDG_STATE_HOME") == "" && !strings.HasSuffix(path, ".local/state") {
			t.Errorf("Default StateDir path should end with .local/state, got: %s", path)
		}
	})

	// User Dirs
	userDirs := map[string]struct {
		getter        func() (string, error)
		envVar        string
		defaultSuffix string
	}{
		"AudioDir":    {d.AudioDir, "XDG_MUSIC_DIR", "Music"},
		"DesktopDir":  {d.DesktopDir, "XDG_DESKTOP_DIR", "Desktop"},
		"DocumentDir": {d.DocumentDir, "XDG_DOCUMENTS_DIR", "Documents"},
		"DownloadDir": {d.DownloadDir, "XDG_DOWNLOAD_DIR", "Downloads"},
		"PictureDir":  {d.PictureDir, "XDG_PICTURES_DIR", "Pictures"},
		"PublicDir":   {d.PublicDir, "XDG_PUBLICSHARE_DIR", "Public"},
		"TemplateDir": {d.TemplateDir, "XDG_TEMPLATES_DIR", "Templates"},
		"VideoDir":    {d.VideoDir, "XDG_VIDEOS_DIR", "Videos"},
	}

	for name, data := range userDirs {
		t.Run(name, func(t *testing.T) {
			path, err := data.getter()
			checkPath(t, name, path, err)
			if os.Getenv(data.envVar) == "" && !strings.HasSuffix(path, data.defaultSuffix) {
				t.Errorf("Default %s path should end with %s, got: %s", name, data.defaultSuffix, path)
			}
		})
	}

	t.Run("FontDir", func(t *testing.T) {
		path, err := d.FontDir()
		checkPath(t, "FontDir", path, err)
		// Check it's relative to DataDir
		dataDir, _ := d.DataDir()
		expected := filepath.Join(dataDir, "fonts")
		if path != expected {
			t.Errorf("FontDir path expected %s, got: %s", expected, path)
		}
	})
}

func TestLinuxXdgOverrides(t *testing.T) {
	// Requires Go 1.17+ for t.Setenv
	if !strings.HasPrefix(os.Getenv("GOVERSION"), "go1.17") && !strings.HasPrefix(os.Getenv("GOVERSION"), "go1.18") && !strings.HasPrefix(os.Getenv("GOVERSION"), "go1.19") && !strings.HasPrefix(os.Getenv("GOVERSION"), "go1.2") { // Adjust check for future Go versions
		t.Skip("Skipping XDG override tests, requires Go 1.17+ for t.Setenv")
	}

	d := NewDirs()
	testDir := t.TempDir() // Create a temporary directory for testing

	tests := []struct {
		name   string
		envVar string
		getter func() (string, error)
	}{
		{"CacheDir", "XDG_CACHE_HOME", d.CacheDir},
		{"ConfigDir", "XDG_CONFIG_HOME", d.ConfigDir},
		{"DataDir", "XDG_DATA_HOME", d.DataDir},
		{"StateDir", "XDG_STATE_HOME", d.StateDir},
		{"RuntimeDir", "XDG_RUNTIME_DIR", d.RuntimeDir},
		{"ExecutableDir", "XDG_BIN_HOME", d.ExecutableDir},
		{"AudioDir", "XDG_MUSIC_DIR", d.AudioDir},
		{"DesktopDir", "XDG_DESKTOP_DIR", d.DesktopDir},
		{"DocumentDir", "XDG_DOCUMENTS_DIR", d.DocumentDir},
		{"DownloadDir", "XDG_DOWNLOAD_DIR", d.DownloadDir},
		{"PictureDir", "XDG_PICTURES_DIR", d.PictureDir},
		{"PublicDir", "XDG_PUBLICSHARE_DIR", d.PublicDir},
		{"TemplateDir", "XDG_TEMPLATES_DIR", d.TemplateDir},
		{"VideoDir", "XDG_VIDEOS_DIR", d.VideoDir},
		// PreferenceDir uses XDG_CONFIG_HOME
		// DataLocalDir uses XDG_DATA_HOME
		// FontDir uses XDG_DATA_HOME
	}

	for _, tt := range tests {
		t.Run(tt.name+" Override", func(t *testing.T) {
			expectedPath := filepath.Join(testDir, tt.name+"_override")
			t.Setenv(tt.envVar, expectedPath)

			path, err := tt.getter()
			if err != nil {
				t.Fatalf("getter for %s returned error: %v", tt.name, err)
			}
			if path != expectedPath {
				t.Errorf("Expected path %s from env var %s, but got %s", expectedPath, tt.envVar, path)
			}
		})
	}

	// Test dependent dirs explicitly
	t.Run("PreferenceDir Override", func(t *testing.T) {
		expectedPath := filepath.Join(testDir, "PreferenceDir_override")
		t.Setenv("XDG_CONFIG_HOME", expectedPath)
		path, err := d.PreferenceDir()
		checkPath(t, "PreferenceDir", path, err)
		if path != expectedPath {
			t.Errorf("Expected PreferenceDir %s from XDG_CONFIG_HOME, got %s", expectedPath, path)
		}
	})

	t.Run("DataLocalDir Override", func(t *testing.T) {
		expectedPath := filepath.Join(testDir, "DataLocalDir_override")
		t.Setenv("XDG_DATA_HOME", expectedPath)
		path, err := d.DataLocalDir()
		checkPath(t, "DataLocalDir", path, err)
		if path != expectedPath {
			t.Errorf("Expected DataLocalDir %s from XDG_DATA_HOME, got %s", expectedPath, path)
		}
	})

	t.Run("FontDir Override", func(t *testing.T) {
		expectedDataPath := filepath.Join(testDir, "FontDir_override_data")
		t.Setenv("XDG_DATA_HOME", expectedDataPath)
		path, err := d.FontDir()
		checkPath(t, "FontDir", path, err)
		expectedFontPath := filepath.Join(expectedDataPath, "fonts")
		if path != expectedFontPath {
			t.Errorf("Expected FontDir %s based on XDG_DATA_HOME, got %s", expectedFontPath, path)
		}
	})
}
