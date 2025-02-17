package container

import "github.com/ChristianSilvaDev/GoMail/src/internal/route/v1"

func ProvideRoutesV1() *route.RoutesV1 {
	return &route.RoutesV1{PingController: ProvidePingController()}
}
