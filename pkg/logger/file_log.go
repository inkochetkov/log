package logger

import (
	"path"

	"github.com/inkochetkov/exist/pkg/exist"
)

const (
	defaultPatch         = "./"
	defaultPatchFileName = "info.log"

	empty = ""
)

func checkFile(patch, patchName string) (string, error) {

	if patch == empty {
		patch = defaultPatch
	}

	if patchName == empty {
		patchName = defaultPatchFileName
	}

	filePatch := path.Join(patch, patchName)

	if ok := exist.CheckFile(filePatch); ok {
		return filePatch, nil
	}

	_, err := exist.InitDirFile(patch, patchName)
	if err != nil {
		return empty, err
	}

	return filePatch, nil
}
