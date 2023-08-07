package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/modaniru/tgf-gRPC/internal/client"
	"github.com/modaniru/tgf-gRPC/internal/utils"
	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
)

type Service struct {
	twitchClient *client.Queries
}

// Return new service.Service
func NewService(twitchClient *client.Queries) *Service {
	return &Service{twitchClient: twitchClient}
}

// Returns general follow list by []string of logins
func (s *Service) GetGeneralFollows(nicknames []string) (*pkg.GetTGFResponse, error) {
	users, err := s.GetUsersInfo(nicknames, "login")
	if err != nil {
		return nil, err
	}
	if len(users) != len(nicknames) {
		return nil, errors.New("some users was not found")
	}
	usersMap := utils.ReponseUserToHashMap(users)
	generalFollows := make(map[string]*pkg.OldestUser)
	now := time.Now()
	channel := make(chan []client.FollowInfo, 100)
	//Получение подписок следующих пользователей
	for _, v := range users {
		go func() {
			f, err := s.twitchClient.GetFollows(v.Id)
			if err != nil {
				channel <- nil
				return
			}
			channel <- f
		}()
	}
	//Инициализация списка
	followList := <-channel
	if followList == nil {
		return nil, errors.New("error")
	}
	for _, v := range followList[1:] {
		generalFollows[v.ToId] = &pkg.OldestUser{
			Username: usersMap[v.FromId].DisplayName,
			Date:     v.FollowedAt,
		}
	}
	for i := 1; i < len(users); i++ {
		nextGeneralFollows := make(map[string]*pkg.OldestUser)
		followList = <-channel
		if followList == nil {
			return nil, errors.New("error")
		}
		for _, v := range followList {
			prev, ok := generalFollows[v.ToId]
			if ok {
				prevTime, err := time.Parse("2006-01-02T15:04:05Z", prev.Date)
				if err != nil {
					return nil, err
				}
				nowTime, err := time.Parse("2006-01-02T15:04:05Z", v.FollowedAt)
				if err != nil {
					return nil, err
				}
				oldestUser := generalFollows[v.ToId]
				if prevTime.Compare(nowTime) > 0 {
					oldestUser = &pkg.OldestUser{
						Username: usersMap[v.FromId].DisplayName,
						Date:     v.FollowedAt,
					}
				}
				nextGeneralFollows[v.ToId] = oldestUser
			}
		}
		generalFollows = nextGeneralFollows
	}
	ids := make([]string, 0, len(generalFollows))
	for k := range generalFollows {
		ids = append(ids, k)
	}
	streamersInfo, err := s.GetUsersInfo(ids, "id")
	if err != nil {
		return nil, err
	}
	generalStreamers := make([]*pkg.Streamer, 0, len(streamersInfo))
	for _, v := range streamersInfo {
		generalStreamers = append(generalStreamers, &pkg.Streamer{
			Streamer:   v,
			OldestUser: generalFollows[v.Id],
		})
	}
	fmt.Println(time.Since(now).Seconds())
	return &pkg.GetTGFResponse{
		InputedUsers:     users,
		GeneralStreamers: generalStreamers,
	}, nil
}

// Get user info by []string logins
func (s *Service) GetUsersInfo(nicknames []string, searchType string) ([]*pkg.ResponseUser, error) {
	if len(nicknames) == 0 {
		return []*pkg.ResponseUser{}, nil
	}
	users, err := s.twitchClient.GetUsersInfo(nicknames, searchType)
	if err != nil {
		return nil, err
	}
	return utils.UserInfoToResponseUser(users), nil
}
