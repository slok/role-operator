package main

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

// Defaults.
const (
	resyncIntervalSecondsDef = 30
	dryRunDef                = false
	developmentDef           = false
	debugDef                 = false
)

// Defaults.
var (
	kubehomeDef = filepath.Join(homedir.HomeDir(), ".kube", "config")
)

// Flags are the flags of the program.
type Flags struct {
	ResyncIntervalSeconds int
	KubeConfig            string
	DryRun                bool
	Development           bool
	Debug                 bool
}

// NewFlags returns the flags of the commandline.
func NewFlags() *Flags {
	flags := &Flags{}
	fl := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	fl.IntVar(&flags.ResyncIntervalSeconds, "resync-interval", resyncIntervalSecondsDef, "resync  seconds of the controller")
	fl.StringVar(&flags.KubeConfig, "kubeconfig", kubehomeDef, "kubernetes configuration path, only used when development mode enabled")
	fl.BoolVar(&flags.DryRun, "dry-run", dryRunDef, "run in dry-run mode")
	fl.BoolVar(&flags.Development, "development", developmentDef, "development flag will allow to run outside a kubernetes cluster")
	fl.BoolVar(&flags.Debug, "debug", debugDef, "enable debug mode")

	fl.Parse(os.Args[1:])

	return flags
}
