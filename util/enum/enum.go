package enum

import pb "github.com/rexposadas/learning-protocol-buffers/proto"

func DoEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}
