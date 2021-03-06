package e2e

import (
	"flag"
	"path/filepath"

	"github.com/appscode/go/flags"
	logs "github.com/appscode/go/log/golog"
	"k8s.io/client-go/util/homedir"
)

type E2EOptions struct {
	KubeContext string
	KubeConfig  string
}

var (
	options = &E2EOptions{
		KubeConfig: filepath.Join(homedir.HomeDir(), ".kube", "config"),
	}
	brokerImageFlag = "shudipta/service-broker:latest"
	storageClass    = "standard"
)

func init() {
	flag.StringVar(&options.KubeConfig, "kubeconfig", options.KubeConfig, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	flag.StringVar(&options.KubeContext, "kube-context", "", "Name of kube context")

	flag.StringVar(&brokerImageFlag, "broker-image", brokerImageFlag,
		"The container image for the broker to test against")
	flag.StringVar(&storageClass, "storage-class", storageClass,
		"name of the storage-class for database storage")
	//framework.RegisterParseFlags()

	enableLogging()
	flag.Parse()
}

func enableLogging() {
	defer func() {
		logs.InitLogs()
		defer logs.FlushLogs()
	}()
	flag.Set("logtostderr", "true")
	logLevelFlag := flag.Lookup("v")
	if logLevelFlag != nil {
		if len(logLevelFlag.Value.String()) > 0 && logLevelFlag.Value.String() != "0" {
			return
		}
	}
	flags.SetLogLevel(2)
}
