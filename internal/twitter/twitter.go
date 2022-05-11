package twitter

import (
	"context"
	"go-twitter-exporter/internal/config"
	"log"

	"github.com/sivchari/gotwtr"
)

func GetMultipleUserProfiles(targetIds []string) gotwtr.UsersResponse {
	client := gotwtr.New(config.GetConfig().TWITTER.TWITTER_BEARER_TOKEN)

	opts := &gotwtr.RetrieveUserOption{
		UserFields: []gotwtr.UserField{
			gotwtr.UserFieldPublicMetrics,
		},
	}
	un, err := client.RetrieveMultipleUsersWithUserNames(context.Background(), targetIds, opts)
	if err != nil {
		log.Println(err)
	}
	return *un
}
