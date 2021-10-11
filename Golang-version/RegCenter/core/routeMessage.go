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
	pos := hash & 1023
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

func findInstance(pos uint16) (string, error) {
	idx := int(pos) * (len(instances) / 1024)
	instance := instances[idx]
	for instance == "E" {
		idx--
		count := 0
		for idx == -1 {
			if count >= 10 {
				return "", myErr.NewError(500, "instances not found.")
			}
			idx = len(instances) - 1
			count++
		}
		instance = instances[idx]
	}
	return instance, nil
}
