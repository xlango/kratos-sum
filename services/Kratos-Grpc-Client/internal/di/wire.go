// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"Kratos-Grpc-Client/internal/dao"
	"Kratos-Grpc-Client/internal/server/grpc"
	"Kratos-Grpc-Client/internal/server/http"
	"Kratos-Grpc-Client/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
