package core

import (
	myErr "RegCenter/error"
	"RegCenter/proto/execTxnRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
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
	if len(instanceList) == 0 {
		return 0, myErr.NewError(500, "No instances!")
	}
	if pos < instanceList[0] {
		idx = instanceList[0]
	} else {
		found := false
		for i := 1; i < len(instanceList); i++ {
			if pos < instanceList[i] {
				idx = instanceList[i]
				found = true
				break
			}
		}
		if !found {
			idx = instanceList[0]
		}
	}
	return idx, nil
}
