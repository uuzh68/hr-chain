package workers

import (
	"context"

	"gitlab.inn4science.com/gophers/service-kit/api/infoworker"
	"gitlab.inn4science.com/gophers/service-kit/routines"
	"github.com/uuzh68/hr-chain/config"
	"github.com/uuzh68/hr-chain/info"
	"github.com/uuzh68/hr-chain/workers/api"
	"github.com/uuzh68/hr-chain/workers/foobar"
)

var WorkerChief routines.Chief

func GetChief() *routines.Chief {
	WorkerChief = routines.Chief{EnableByDefault: true}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})

	//add worker only in dev mode
	if config.Config().InfoWorker != nil {
		//create context with pointer to chief
		ctx := context.Background()
		ctx = context.WithValue(ctx, "chief", &WorkerChief)
		worker := infoworker.GetInfoWorker(*config.Config().InfoWorker, ctx, info.App)
		WorkerChief.AddWorker(config.WorkerInfoServer, worker)
		if !WorkerChief.EnableByDefault {
			WorkerChief.EnableWorker(config.WorkerInfoServer)
		}

	}
	return &WorkerChief
}
