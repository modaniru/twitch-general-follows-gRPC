package utils

import (
	"reflect"
	"testing"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/client"
)

type Test struct {
	title  string
	given  []client.UserInfo
	result []*pkg.ResponseUser
}

func TestMapUserInfoToResponseUser(t *testing.T) {
	given := []Test{
		{title: "mapping normal data",
			given: []client.UserInfo{
				{
					Id:              "123123",
					Login:           "test",
					DisplayName:     "Test",
					BroadcasterType: "partner",
					Description:     "channel description",
					ProfileImageURL: "url",
					OfflineImageURL: "url",
					ViewCount:       0,
					Email:           "",
					CreatedAt:       "12-12-1212",
				},
			},
			result: []*pkg.ResponseUser{
				{
					DisplayName:     "Test",
					ImageLink:       "url",
					Id:              "123123",
					BroadcasterType: "partner",
				},
			}},
		{
			title:  "empty given test",
			given:  []client.UserInfo{},
			result: []*pkg.ResponseUser{},
		},
	}
	for _, g := range given {
		res := UserInfoToResponseUser(g.given)
		ok := reflect.DeepEqual(res, g.result)
		if !ok {
			t.Errorf("test name: %s: require %+v, response %+v", given[0].title, given[0].result, res)
		}
	}
}

func TestResponseUserToHashMap(t *testing.T){
	array := []*pkg.ResponseUser{
		{DisplayName: "test1", ImageLink: "url", Id: "123123", BroadcasterType: "partner"},
		{DisplayName: "test2", ImageLink: "url", Id: "123124", BroadcasterType: "partner"},
		{DisplayName: "test3", ImageLink: "url", Id: "123125", BroadcasterType: "partner"},
	}
	except := make(map[string]*pkg.ResponseUser)
	for _, u := range array{
		except[u.Id] = u
	}
	actual := ReponseUserToHashMap(array)
	ok := reflect.DeepEqual(except, actual)
	if !ok{
		t.Errorf("test name: %s: require \n%+v, response \n%+v", "TestResponseUserToHashMap", except, actual)
	}
}