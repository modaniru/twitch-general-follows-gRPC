package utils

import (
	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/client"
)

// TODO documentation

func UserInfoToResponseUser(usersInfo []client.UserInfo) []*pkg.ResponseUser{
	res := make([]*pkg.ResponseUser, len(usersInfo))
	for i, v := range usersInfo{
		res[i] = &pkg.ResponseUser{
			Id: v.Id,
			ImageLink: v.ProfileImageURL,
			DisplayName: v.DisplayName,
			BroadcasterType: v.BroadcasterType,
		}
	}
	return res
}

func ReponseUserToHashMap(users []*pkg.ResponseUser) map[string]*pkg.ResponseUser{
	res := make(map[string]*pkg.ResponseUser)
	for _, v := range users{
		res[v.Id] = v
	}
	return res
}