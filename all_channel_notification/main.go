package main

import (
	"fmt"
	"time"

	"github.com/samber/lo"

	"traQ-Script/api"
	"traQ-Script/util"
)

type ChannelsWithSubscriptionLevel lo.Tuple2[string, int]

func main() {
	_, _, err := util.Login()
	if err != nil {
		panic(err)
	}

	allChannelDetails, err := api.GetChannels()
	if err != nil {
		panic(err)
	}
	allPublicChannels := lo.Map(allChannelDetails.Public, func(c api.GetChannelResponsePublic, _ int) string { return c.Id })
	allSubscribedChannelsWithLevel, err := api.GetSubscriptions()
	if err != nil {
		panic(err)
	}
	allSubscribedChannels := lo.Map(*allSubscribedChannelsWithLevel, func(c api.GetMeSubscriptions, _ int) string { return c.ChannelId })
	notSubscribedChannels := difference(allPublicChannels, allSubscribedChannels)
	notSubscribedChannelsWithSubscriptionLevel := lo.Map(notSubscribedChannels, func(c string, _ int) ChannelsWithSubscriptionLevel {
		return ChannelsWithSubscriptionLevel(lo.T2(c, 1))
	})

	for _, channelWithSubscriptionLevel := range notSubscribedChannelsWithSubscriptionLevel {
		fmt.Printf("Subscribe to %s\n", channelWithSubscriptionLevel.A)
		err = api.PutSubscriptions(channelWithSubscriptionLevel.A, channelWithSubscriptionLevel.B)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
