package core

import (
	"crypto/md5"
	log "github.com/sirupsen/logrus"
)

func hashCode(s string) uint32 {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(s))
	md5Result := Md5Inst.Sum([]byte(""))
	log.Debugf("%x\n\n", md5Result)

	var result uint32
	for i := 0; i < len(md5Result); i++ {
		result = 31*result + uint32(md5Result[i])
	}
	return result
}
