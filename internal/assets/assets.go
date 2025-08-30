package assets

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"sprout/internal/utils/system"
	"strconv"
)

//go:embed assets/.config
//go:embed assets/.themes
var embeddedFiles embed.FS

func ExtractAssets(executables []string) error {
	
	u, err := system.GetUser()
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

	mappings := map[string]string{
		"assets/.config": filepath.Join(home, ".config"),
		"assets/.themes": filepath.Join(home, ".themes"),
	}

	execMap := make(map[string]struct{})
	for _, e := range executables {
		execMap[e] = struct{}{}
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

			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				return err
			}
			if err := os.Chown(filepath.Dir(outPath), uid, gid); err != nil {
				return err
			}

			data, err := embeddedFiles.ReadFile(path)
			if err != nil {
				return err
			}

			mode := os.FileMode(0644)

			if _, ok := execMap[path]; ok {
				mode = 0755
			}

			if err := os.WriteFile(outPath, data, mode); err != nil {
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
