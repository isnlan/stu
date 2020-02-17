package dagutils

import (
	"context"
	dag "github.com/ipfs/go-merkledag"
)

func IpfsHash(data string) (string, error) {
	root := new(dag.ProtoNode)
	e := NewDagEditor(root, nil)

	err := insert(e, "a", data, false)
	if err != nil {
		return "", err
	}
	c := e.GetNode().Cid()
	return c.String(), nil
}

func insert(e *Editor, path, data string, create bool) error {
	child := dag.NodeWithData([]byte(data))
	err := e.GetDagService().Add(context.Background(), child)
	if err != nil {
		return err
	}

	var c func() *dag.ProtoNode
	if create {
		c = func() *dag.ProtoNode {
			return &dag.ProtoNode{}
		}
	}

	return e.InsertNodeAtPath(context.Background(), path, child, c)
}
