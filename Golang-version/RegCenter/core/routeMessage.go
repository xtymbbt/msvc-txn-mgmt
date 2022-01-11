package core

import (
	"RegCenter/proto/execTxnRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"math"
)

func RouteMessage(message *execTxnRpc.TxnMessage) (*execTxnRpc.TxnStatus, error) {
	hash := hashCode(message.TreeUuid)
	log.Debugf("hash is: %v\n", hash)
	pos := hash
	log.Debugf("pos is: %v\n", pos)
	instance, err := findInstance(pos)
	if err != nil {
		return &execTxnRpc.TxnStatus{
			Status:  501,
			Message: "Route message failed. Cannot find instance, please wait and retry.",
		}, err
	}
	conn := set[instance].conn
	c := execTxnRpc.NewExecTxnRpcClient(conn)
	return c.ExecTxn(context.Background(), message)
}

func findInstance(pos uint32) (uint32, error) {
	var idx uint32
	var i uint32
	for i = 0; ; i++ {
		if i == math.MaxUint32 {
			i = 0
		}
		if _, ok := set[i]; ok {
			idx = i
			break
		}
	}
	return idx, nil
}
