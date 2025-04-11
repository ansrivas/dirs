package dirs

// Dirs defines methods to retrieve platform-specific directories.
type Dirs interface {
	HomeDir() (string, error)
	CacheDir() (string, error)
	ConfigDir() (string, error)
	DataDir() (string, error)
	DataLocalDir() (string, error)
	ExecutableDir() (string, error)
	PreferenceDir() (string, error)
	RuntimeDir() (string, error)
	StateDir() (string, error)
	AudioDir() (string, error)
	DesktopDir() (string, error)
	DocumentDir() (string, error)
	DownloadDir() (string, error)
	FontDir() (string, error)
	PictureDir() (string, error)
	PublicDir() (string, error)
	TemplateDir() (string, error)
	VideoDir() (string, error)
}
