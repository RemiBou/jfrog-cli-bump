package processor

import (
	"context"
	"github.com/jfrog/jfrog-cli-bump/commands/processor/model"
)

type Processor interface {
	Process(ctx context.Context, versionToUpgrade model.PackageVersion, inVersionFile string, inChecksumFile string) (versionFile string, checksumFile string, err error)
}
