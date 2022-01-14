package core

import (
	myErr "algorithm-test/error"
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type TxnMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online      bool              `protobuf:"varint,1,opt,name=online,proto3" json:"online,omitempty"`
	TreeUuid    string            `protobuf:"bytes,2,opt,name=tree_uuid,json=treeUuid,proto3" json:"tree_uuid,omitempty"`
	ServiceUuid string            `protobuf:"bytes,3,opt,name=service_uuid,json=serviceUuid,proto3" json:"service_uuid,omitempty"`
	ParentUuid  string            `protobuf:"bytes,4,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
	Children    []string          `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	DbName      string            `protobuf:"bytes,6,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	TableName   string            `protobuf:"bytes,7,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Method1     bool              `protobuf:"varint,8,opt,name=method1,proto3" json:"method1,omitempty"`
	Method2     bool              `protobuf:"varint,9,opt,name=method2,proto3" json:"method2,omitempty"`
	Query       string            `protobuf:"bytes,10,opt,name=query,proto3" json:"query,omitempty"` // 若query有多个值，则使用","分隔开。
	Data        map[string]string `protobuf:"bytes,11,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func RouteMessage(message *TxnMessage) (*IstsInfo, error) {
	hash := hashCode(message.TreeUuid)
	//log.Debugf("hash is: %v\n", hash)
	pos := hash
	//log.Debugf("pos is: %v\n", pos)
	instance, err := findInstance(pos)
	istsInfo := set[instance]
	fmt.Printf("instance is: %#v\n", istsInfo)
	mutex.Lock()
	istsInfo.Conn.TxnNum++
	mutex.Unlock()
	return istsInfo, err
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
	fmt.Println(idx)
	return idx, nil
}
