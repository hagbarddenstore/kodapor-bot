package version

// Package is the overall, canonical project import path under which the
// package was built.
var Package = "github.com/hagbarddenstore/kodapor-bot"

// Version indicates which version of the binary is running. This is set to
// the latest release tag by hand, always suffixed by "+unknown". During
// build, it will be replaced by the actual version. The value here will be
// used if the application is run after a go get based install.
var Version = "5831959+unknown"

// Name is the overall application name.
var Name = "Kodapor-bot"
