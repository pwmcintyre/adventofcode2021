// tools allows us to vendor non-runtime dependencies
//
// More info:
//   https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
//   https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
//

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
