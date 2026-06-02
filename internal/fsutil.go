package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

const NpmrcFilename = ".npmrc"
const NpmrcmDirName = ".npmrcm"
const NpmrcmConfigFilename = "config.json"
const ProfilesDirName = "profiles"

type FsUtil struct {
	UserHomeDir string
}

func (f *FsUtil) GetNpmrc() (*Npmrc, error) {
	npmrcPath := filepath.Join(f.UserHomeDir, NpmrcFilename)

	content, err := f.ReadFile(npmrcPath)

	if err != nil {
		return nil, err
	}

	return &Npmrc{
		Path:    npmrcPath,
		Content: content,
	}, nil
}

func (f *FsUtil) ReplaceNpmrcContent(content string) error {
	npmrc, err := f.GetNpmrc()

	if err != nil {
		return err
	}

	err = os.WriteFile(npmrc.Path, []byte(content), 0666)

	if err != nil {
		return err
	}

	return nil
}

func (f *FsUtil) GetProfiles() ([]NpmrcProfile, error) {
	npmrc, err := f.GetNpmrc()

	if err != nil {
		return nil, err
	}

	appConfigDir, err := f.GetAppConfigDir()

	if err != nil {
		return nil, err
	}

	profileFiles, err := os.ReadDir(appConfigDir)

	if err != nil {
		return nil, err
	}

	var profiles []NpmrcProfile

	for _, profileFile := range profileFiles {
		if profileFile.IsDir() {
			continue
		}

		path := filepath.Join(appConfigDir, profileFile.Name())
		content, err := f.ReadFile(path)

		if err != nil {
			return nil, err
		}

		profile := NpmrcProfile{
			Name:    profileFile.Name(),
			Path:    path,
			Active:  npmrc.Content == content,
			Content: content,
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (f *FsUtil) GetAppConfigDir() (string, error) {
	var home string

	if f.UserHomeDir == "" {
		osHome, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		home = osHome
	} else {
		home = f.UserHomeDir
	}

	return filepath.Join(home, NpmrcmDirName, ProfilesDirName), nil
}

func (f *FsUtil) IsConfigured() bool {
	path, err := f.GetAppConfigDir()

	if err != nil {
		return false
	}

	path = filepath.Join(path, NpmrcmConfigFilename)

	_, err = f.ReadFile(path)

	if err != nil {
		return false
	}

	return true
}

func (f *FsUtil) ReadFile(path string) (string, error) {
	result, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Could not read file at path [%s]\n", path)
		}

		return "", err
	}

	return string(result), nil
}
