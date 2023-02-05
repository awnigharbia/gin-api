package utils

import "os"


func CheckFile(fileName string, fileTypes []string, fileSizeMin int64, fileSizeMax int64) string {
	// Check if file extension is supported
	var supported bool
	for _, fileType := range fileTypes {
		if fileHasExtension(fileName, fileType) {
			supported = true
			break
		}
	}

	if !supported {
		return fileName + " has an unsupported file type."
	}

	// Check if file size is within specified range
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return fileName + " does not exist."
	}

	fileSize := fileInfo.Size()
	if fileSizeMin <= fileSize && fileSize <= fileSizeMax {
		return fileName + " is a valid file."
	} else {
		return fileName + " is outside the specified file size range."
	}
}

func fileHasExtension(fileName string, extension string) bool {
	return len(fileName) > len(extension) && fileName[len(fileName)-len(extension):] == extension
}
