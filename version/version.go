package version

import (
	"fmt"
	"runtime"
)


//var needs to be used instead of const for ldflags
var (
	Version           = "v0.0.2"
	GoVersion         = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)
