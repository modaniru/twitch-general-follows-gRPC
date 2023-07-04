package service

import (
	"errors"
	"time"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/client"
	"github.com/modaniru/tgf-gRPC/src/utils"
)

type Service struct {
	twitchClient *client.Queries
}

func NewService(twitchClient *client.Queries) *Service{
	return &Service{twitchClient: twitchClient}
}

func (s *Service) GetGeneralFollows(nicknames []string) (*pkg.GetTGFResponse, error){
	users, err := s.GetUsersInfo(nicknames, "login")
	if err != nil{
		return nil, err
	}
	usersMap := utils.ReponseUserToHashMap(users)
	generalFollows := make(map[string]*pkg.OldestUser)
	channel := make(chan []client.FollowInfo)
	//Получение списка подписок первого пользователя
	go s.twitchClient.GetFollows(users[0].Id, channel)
	//Получение подписок следующих пользователей
	for _, v := range users[1:]{
		go s.twitchClient.GetFollows(v.Id, channel)
	}
	//Инициализация списка
	followList := <-channel
	if followList == nil {
		return nil, errors.New("error")
	}
	for _, v := range followList[1:] {
		generalFollows[v.ToId] = &pkg.OldestUser{
			User: usersMap[v.FromId],
			Date: v.FollowedAt,
		}
	}
	for i := 1; i < len(users); i++{
		nextGeneralFollows  := make(map[string]*pkg.OldestUser)
		followList = <-channel
		if followList == nil {
			return nil, errors.New("error")
		}
		for _, v := range followList{
			prev, ok := generalFollows[v.ToId]
			if ok{
				prevTime, err := time.Parse("2006-01-02T15:04:05Z", prev.Date)
				if err != nil{
					return nil, err
				}
				nowTime, err := time.Parse("2006-01-02T15:04:05Z", v.FollowedAt)
				if err != nil{
					return nil, err
				}
				oldestUser := generalFollows[v.ToId]
				if prevTime.Compare(nowTime) > 0{
					oldestUser = &pkg.OldestUser{
						User: usersMap[v.FromId],
						Date: v.FollowedAt,
					}
				}
				nextGeneralFollows[v.ToId] = oldestUser
			}
		}
		generalFollows = nextGeneralFollows
	}
	ids := make([]string, 0, len(generalFollows))
	for k := range generalFollows{
		ids = append(ids, k)
	}
	streamersInfo, err := s.GetUsersInfo(ids, "id")
	if err != nil{
		return nil, err
	}
	generalStreamers := make([]*pkg.Streamer, 0, len(streamersInfo))
	for _, v := range streamersInfo{
		generalStreamers = append(generalStreamers, &pkg.Streamer{
			Streamer: v,
			OldestUser: generalFollows[v.Id],
		})
	}
	return &pkg.GetTGFResponse{
		InputedUsers: users,
		GeneralStreamers: generalStreamers,
	}, nil
}

func (s *Service) GetUsersInfo(nicknames []string, searchType string) ([]*pkg.ResponseUser, error){
	users, err := s.twitchClient.GetUsersInfo(nicknames, searchType)
	if err != nil{
		return nil, err
	}
	return utils.UserInfoToResponseUser(users), nil
}
