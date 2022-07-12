package protomap

import pb "github.com/rexposadas/learning-protocol-buffers/proto"

func DoMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IDWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}
