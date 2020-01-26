// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"Kratos-Grpc-Server/internal/dao"
	"Kratos-Grpc-Server/internal/server/grpc"
	"Kratos-Grpc-Server/internal/server/http"
	"Kratos-Grpc-Server/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
