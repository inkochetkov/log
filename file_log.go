package log

import (
	"path"

	"github.com/inkochetkov/exist"
)

// checkFile - check directory and file for existence
func checkFile(patch, patchName string) (string, error) {

	if patch == empty {
		// default values
		patch = defaultPatch
	}

	if patchName == empty {
		// default values
		patchName = defaultPatchFileName
	}

	// full patch for file
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
