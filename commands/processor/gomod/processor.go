package gomod

import (
	"context"
	"github.com/jfrog/jfrog-cli-bump/commands/processor/model"
	"strings"
)

type processor struct {
}

func (processor) Process(ctx context.Context, versionToUpgrade model.PackageVersion, inVersionFile string, inChecksumFile string) (versionFile string, checksumFile string, err error) {
	cursor := 0

	versionFile = inVersionFile
	checksumFile = inChecksumFile

	for cursor < len(inVersionFile) {
		startOfPackage := strings.Index(inVersionFile[cursor:], versionToUpgrade.Name)
		if startOfPackage < 0 {
			break
		}

		cursor += startOfPackage
		endOfLine := cursor +strings.Index(inVersionFile[cursor:], "\n")
		startOfVersion := cursor +strings.LastIndex(inVersionFile[cursor:endOfLine], " ")
		versionFile = strings.ReplaceAll(versionFile, inVersionFile[startOfVersion+1:endOfLine], versionToUpgrade.Version)

		// Move one position forward to start from the next character
		cursor++
	}

	return
}