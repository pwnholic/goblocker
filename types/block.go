package types

import (
	"crypto/sha256"

	pb "google.golang.org/protobuf/proto"
)

func HashBlock(block *proto.Block) []byte {
    b,err := pb.Marshal(block)
    if err !=nil {
         panic(err) 
    }

    hash:= sha256.Sum256(b)
    return hash[:]
}
