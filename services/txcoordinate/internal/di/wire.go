// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"ilinkcloud/services/txcoordinate/internal/dao"
	"ilinkcloud/services/txcoordinate/internal/server/grpc"
	"ilinkcloud/services/txcoordinate/internal/server/http"
	"ilinkcloud/services/txcoordinate/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
