package embedassets

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

//go:embed assets/.config
//go:embed assets/.themes
var embeddedFiles embed.FS

func ExtractAssets() error {
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser == "" {
		return fmt.Errorf("this script must be run with sudo")
	}
	u, err := user.Lookup(sudoUser)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return err
	}
	home := u.HomeDir

	// Delete existing directories if they exist
	for _, dirName := range []string{".config", ".themes"} {
		destDir := filepath.Join(home, dirName)
		if err := os.RemoveAll(destDir); err != nil {
			return err
		}
	}

	mappings := map[string]string{
		"assets/.config": filepath.Join(home, ".config"),
		"assets/.themes":  filepath.Join(home, ".themes"),
	}

	for srcRoot, destRoot := range mappings {
		err := fs.WalkDir(embeddedFiles, srcRoot, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(srcRoot, path)
			if err != nil {
				return err
			}

			outPath := filepath.Join(destRoot, relPath)

			if d.IsDir() {
				if err := os.MkdirAll(outPath, 0755); err != nil {
					return err
				}
				return os.Chown(outPath, uid, gid)
			}

			// For files
			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				return err
			}
			// Chown the parent dir (may be redundant but safe)
			if err := os.Chown(filepath.Dir(outPath), uid, gid); err != nil {
				return err
			}

			data, err := embeddedFiles.ReadFile(path)
			if err != nil {
				return err
			}

			if err := os.WriteFile(outPath, data, 0644); err != nil {
				return err
			}

			return os.Chown(outPath, uid, gid)
		})
		if err != nil {
			return err
		}
	}
	return nil
}