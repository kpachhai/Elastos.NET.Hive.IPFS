package cmdenv

import (
	"fmt"

	"github.com/elastos/Elastos.NET.Hive.IPFS/commands"
	"github.com/elastos/Elastos.NET.Hive.IPFS/core"
	coreiface "github.com/elastos/Elastos.NET.Hive.IPFS/core/coreapi/interface"

	cmds "gx/ipfs/QmSXUokcP4TJpFfqozT69AVAYRtzXVMUjzQVkYX41R9Svs/go-ipfs-cmds"
	config "gx/ipfs/QmVFZsFtfRgn6hxEAyW5rDiuUYPpiCML4XHtz1p7LDsdon/go-ipfs-config"
)

// GetNode extracts the node from the environment.
func GetNode(env interface{}) (*core.IpfsNode, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.GetNode()
}

// GetApi extracts CoreAPI instance from the environment.
func GetApi(env cmds.Environment) (coreiface.CoreAPI, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.GetApi()
}

// GetConfig extracts the config from the environment.
func GetConfig(env cmds.Environment) (*config.Config, error) {
	ctx, ok := env.(*commands.Context)
	if !ok {
		return nil, fmt.Errorf("expected env to be of type %T, got %T", ctx, env)
	}

	return ctx.GetConfig()
}
