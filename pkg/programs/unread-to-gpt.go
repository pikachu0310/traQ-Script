package programs

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/traPtitech/go-traq"

	"traQ-Script/pkg/api/openai"
	v2 "traQ-Script/pkg/api/traQ/v2"
)

var myBotChannelId = "9f551eae-0e50-4887-984e-ce9d8b3919cc"

func UnreadToGPT() {

	unreadChannels, _, err := v2.GetMyUnreadChannels()
	if err != nil {
		panic(err)
	}

	for i, channel := range unreadChannels {
		if i >= 1 {
			break
		}

		messages, err := getMessagesByUnreadChannel(channel)
		if err != nil {
			return
		}

		formattedMessages := lo.Map(messages, func(message traq.Message, _ int) string {
			return fmt.Sprintf(":@%s: %s", v2.UserIdToUserNameFunc(message.UserId), message.Content)
		})

		inputFormattedMessages := formatMessagesToInputFormat(formattedMessages)

		postedMessage, _, err := v2.PostMessage(myBotChannelId, "#"+v2.ChannelIdToAllParentChannelName(channel.ChannelId), true)
		if err != nil {
			return
		}

		promptMessage := "あるSNS上のあるチャンネルの書き込み全てを入力として渡すので、全体を要約して出力してください。\n入力フォーマットの例は以下の通りです。\n```\nユーザー名1 メッセージ内容1\nユーザー名2 メッセージ内容2\nユーザー名1 メッセージ内容3\nユーザー名3 メッセージ内容4\n```\n上のような入力が与えられるので、以下のように出力してください。\n---\n要約の内容\n\n---"
		openai.ChatReset()
		// openai.UpdateSystemRoleMessage(promptMessage)
		openai.AddMessageAsUser(promptMessage)
		openai.AddMessageAsAssistant("分かりました。以下のような出力でよろしいですか？\n---\n### :@user1::@user2::@user3:が雑談をしています。\n:@user1: メッセージ1とメッセージ3の要約\n:@user2: メッセージ2\n:@user3: メッセージ3\n\n---")
		// openai.AddMessageAsUser("例えば、以下のように出力してください。\n\n---\n\n:@pika: 面白い動画見つける\n:@toki: 一緒に見よ\n:@pikachu: 面白かった\n\n\n---")
		openai.AddMessageAsUser("完璧です！それでは、次で入力を送ります。")
		openai.AddMessageAsAssistant("はい。入力をお願いします。")
		openai.AddMessageAsUser(inputFormattedMessages)
		channelMessage := fmt.Sprintf("!{\"type\":\"channel\",\"raw\":\"#%s\",\"id\":\"%s\"}", v2.ChannelIdToAllParentChannelName(channel.ChannelId), channel.ChannelId)
		_, _, err = openai.Stream(openai.Messages, openai.GPT3dot5Turbo, func(message string) {
			_, err := v2.EditMessage(postedMessage.Id, channelMessage+"\n"+message)
			if err != nil {
				fmt.Println(err)
			}
		})
		if err != nil {
			fmt.Println(err)
		}
		_, err = v2.ReadChannel(channel.ChannelId)
		if err != nil {
			return
		}
	}
}

func getMessagesByUnreadChannel(channel traq.UnreadChannel) (messages []traq.Message, err error) {
	messages = make([]traq.Message, 0)
	lastSearchTime := channel.Since
	var searchResult *traq.MessageSearchResult
	for {
		searchResult, _, err = v2.SearchMessagesByUnread(lastSearchTime, channel.ChannelId)
		if err != nil {
			return
		}
		if searchResult.TotalHits == 0 {
			break
		}
		// TODO 微妙な処理してる↓ので治したいね
		lastSearchTime = searchResult.Hits[len(searchResult.Hits)-1].CreatedAt.Add(1 * time.Second)
		messages = append(messages, searchResult.Hits...)
	}
	return
}

func formatMessagesToInputFormat(formattedMessages []string) string {
	var lines []string
	lines = append(lines, "```")

	latestMessages := LatestMessages(formattedMessages, 3000)
	fmt.Println(len(latestMessages))

	lines = append(lines, latestMessages...)
	lines = append(lines, "```")
	return strings.Join(lines, "\n")
}

func LatestMessages(formattedMessages []string, maxCharacterCount int) (messages []string) {
	count := 0
	for _, message := range reverseArray[string](formattedMessages) {
		count += len(message)
		if count > maxCharacterCount {
			break
		}
		messages = append([]string{message}, messages...)
	}
	return
}

func reverseArray[T any](array []T) []T {
	reversedArray := make([]T, len(array))
	for i, v := range array {
		reversedArray[len(array)-1-i] = v
	}
	return reversedArray
}
