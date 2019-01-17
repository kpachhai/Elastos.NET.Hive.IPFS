// +build !windows,nofuse

package node

import (
	"errors"

	core "github.com/elastos/Elastos.NET.Hive.IPFS/core"
)

func Mount(node *core.IpfsNode, fsdir, nsdir string) error {
	return errors.New("not compiled in")
}
