package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	operatork8scli "github.com/slok/role-operator/pkg/client/k8s/clientset/versioned"
	"github.com/slok/role-operator/pkg/log"
	"github.com/slok/role-operator/pkg/operator"
	k8sservice "github.com/slok/role-operator/pkg/service/k8s"
)

const (
	gracePeriod = 5 * time.Second
)

// Main is the main program.
type Main struct {
	flags  *Flags
	stopC  chan struct{}
	logger log.Logger
}

// Run will run the main program.
func (m *Main) Run() error {
	defer m.stop()

	if m.flags.Debug {
		m.logger.Set("debug")
	}

	// Create k8s service.
	k8ssvc, err := m.createKubernetesService()
	if err != nil {
		return err
	}

	// Create operator.
	errC := make(chan error)
	cfg := m.createOperatorConfig()
	op := operator.New(cfg, k8ssvc, m.logger)
	go func() {
		errC <- op.Run(m.stopC)
	}()

	sigC := m.createSignalChan()
	select {
	case err := <-errC:
		if err != nil {
			m.logger.Infof("operator finished with error: %s", err)
			return err
		}
		m.logger.Infof("controller finished successfuly")
	case s := <-sigC:
		m.logger.Infof("signal %s received", s)
	}

	return nil
}

func (m *Main) stop() {
	m.logger.Infof("stopping everything, waiting %s...", gracePeriod)

	// stop everything and let them time to stop.
	close(m.stopC)
	time.Sleep(gracePeriod)
}

func (m *Main) createKubernetesService() (k8sservice.Service, error) {
	_, opcli, aecli, err := m.createKubernetesClients()
	if err != nil {
		return nil, err
	}

	svc := k8sservice.New(aecli, opcli, m.logger)

	return svc, nil
}

func (m *Main) createKubernetesClients() (kubernetes.Interface, operatork8scli.Interface, apiextensionscli.Interface, error) {
	config, err := m.loadKubernetesConfig()
	if err != nil {
		return nil, nil, nil, err
	}

	k8scli, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	opcli, err := operatork8scli.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	aecli, err := apiextensionscli.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	return k8scli, opcli, aecli, nil
}

func (m *Main) createOperatorConfig() operator.Config {
	return operator.Config{
		ResyncDuration: time.Duration(m.flags.ResyncIntervalSeconds) * time.Second,
	}
}

// loadKubernetesConfig loads kubernetes configuration based on flags.
func (m *Main) loadKubernetesConfig() (*rest.Config, error) {
	var cfg *rest.Config
	// If devel mode then use configuration flag path.
	if m.flags.Development {
		config, err := clientcmd.BuildConfigFromFlags("", m.flags.KubeConfig)
		if err != nil {
			return nil, fmt.Errorf("could not load configuration: %s", err)
		}
		cfg = config
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("error loading kubernetes configuration inside cluster, check app is running outside kubernetes cluster or run in development mode: %s", err)
		}
		cfg = config
	}

	return cfg, nil
}

func (m *Main) createSignalChan() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	return c
}

func main() {
	m := Main{
		logger: log.Base(),
		flags:  NewFlags(),
		stopC:  make(chan struct{}),
	}

	err := m.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	os.Exit(0)
}
