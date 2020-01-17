// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"ilinkcloud/services/permission/internal/dao"
	"ilinkcloud/services/permission/internal/server/grpc"
	"ilinkcloud/services/permission/internal/server/http"
	"ilinkcloud/services/permission/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
