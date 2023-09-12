package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/mikcatta/simple_bank/db/sqlc"
	"github.com/mikcatta/simple_bank/pb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
