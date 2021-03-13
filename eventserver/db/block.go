package db

import (
	"context"

	"github.com/xuperchain/xuperchain/core/pb"
)

func InsertProtoBlock(block *pb.FilteredBlock) error {
	_, err := MDB.Collection("block").InsertOne(context.TODO(), &block)
	return err
}
