package mock

import (
	"github.com/sirupsen/logrus"
	"github.com/szkiba/xk6-mock/muxpress"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/httpServer", New())
}

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

type RootModule struct {
	srv *Factory
}

func New() *RootModule {
	return &RootModule{
		srv: &Factory{},
	}
}

func (rm *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu:  vu,
		srv: rm.srv,
	}
}

type ModuleInstance struct {
	vu  modules.VU
	srv *Factory
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.srv,
	}
}

type Factory struct {
	vu modules.VU
}

func (f *Factory) New() *Server {
	l := logrus.StandardLogger().Error
	return &Server{
		Server: muxpress.NewServer(l),
	}
}

type Server struct {
	*muxpress.Server
}
