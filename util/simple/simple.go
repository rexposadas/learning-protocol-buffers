package simple

import pb "github.com/rexposadas/learning-protocol-buffers/proto"

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id: 33,
	}
}
