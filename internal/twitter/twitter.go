package twitter

import (
	"context"
	"go-twitter-exporter/internal/config"

	"github.com/sivchari/gotwtr"
)

func GetFollowersCount(targetId string) int {
	client := gotwtr.New(config.GetConfig().TWITTER.TWITTER_BEARER_TOKEN)

	opts := &gotwtr.RetrieveUserOption{
		UserFields: []gotwtr.UserField{
			gotwtr.UserFieldPublicMetrics,
		},
	}
	un, err := client.RetrieveSingleUserWithUserName(context.Background(), targetId, opts)
	if err != nil {
		panic(err)
	}
	return *&un.User.PublicMetrics.FollowersCount
}
