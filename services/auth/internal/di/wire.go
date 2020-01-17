// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"ilinkcloud/services/auth/internal/dao"
	"ilinkcloud/services/auth/internal/server/grpc"
	"ilinkcloud/services/auth/internal/server/http"
	"ilinkcloud/services/auth/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
