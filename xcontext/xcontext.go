package xcontext

import (
	"github.com/ory/hydra/sdk/go/hydra"
	"context"
)

var (
	keyHydraCodeGenSDK = "hydraCodeGenSDK"
)

func InitHydraCodeGenSDK(ctx context.Context, sdk *hydra.CodeGenSDK) context.Context {

	return context.WithValue(ctx, keyHydraCodeGenSDK, sdk)
}

func GetHydraCodeGenSDK(ctx context.Context) *hydra.CodeGenSDK {

	return ctx.Value(keyHydraCodeGenSDK).(*hydra.CodeGenSDK)
}
