package programs

import (
	"fmt"
	"regexp"
	"time"

	"github.com/samber/lo"
	"github.com/traPtitech/go-traq"

	v2 "traQ-Script/pkg/api/traQ/v2"
)

const pokemonSleepChannel = "364b9ba9-d71a-4aba-9b3b-f613d70dfdc3"

var postChannel = v2.MyBotChannelId

func PokemonSleep() {
	message, _, err := v2.PostMessage(v2.Bot, postChannel, "test", true)
	if err != nil {
		fmt.Println(err)
		return
	}
	messages, err := getMessagesByPeriod(time.Now().Add(-12*time.Hour*24), time.Now(), message.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	messageIds := make([]string, 0)
	result := "## フレンドコードまとめ(自動取得)\n\n|名前|フレンドコード|\n|-|-|\n"

	for _, m := range messages {
		if v2.UserIdToUserNameFunc(m.UserId) == "noc7t" {
			continue
		}
		r := regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}")
		r_uuid := regexp.MustCompile("(-\\d{4}-\\d{4}-\\d{4})|(\\d{4}-\\d{4}-\\d{4}-)")
		messageWithoutUuid := r_uuid.ReplaceAllString(m.Content, "")
		if r.MatchString(messageWithoutUuid) {
			result += fmt.Sprintf("|:@%s:%s|%s|\n", v2.UserIdToUserNameFunc(m.UserId), v2.UserIdToUserNameFunc(m.UserId), r.FindString(messageWithoutUuid))
			fmt.Println(m.Id)
			messageIds = append(messageIds, m.Id)
		}
	}

	messageIds = lo.Map(messageIds, func(messageId string, _ int) string {
		return "https://q.trap.jp/messages/" + messageId
	})

	v2.PostMessage(v2.Bot, postChannel, result, false)
}

func getMessagesByPeriod(after time.Time, before time.Time, progressMessageID string) ([]*traq.Message, error) {
	var messages []*traq.Message
	var searchBefore = before
	for {
		t1 := time.Now()
		res, _, err := v2.SearchMessagesByTime(after, searchBefore)
		fmt.Println(time.Since(t1))
		if err != nil {
			return nil, err
		}
		if len(res.Hits) == 0 {
			break
		}

		for i := range res.Hits {
			messages = append(messages, &res.Hits[i])
		}
		time.Sleep(time.Millisecond * 100)
		searchBefore = messages[len(messages)-1].CreatedAt
		v2.EditMessage(v2.Bot, progressMessageID, fmt.Sprintf("Searching...(%d):loading:", len(messages)))
	}

	return messages, nil
}
