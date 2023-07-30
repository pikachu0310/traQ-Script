package v2

import (
	"net/http"
	"time"

	"github.com/traPtitech/go-traq"
)

func PostMessage(client int, channelId string, content string, embed bool) (*traq.Message, *http.Response, error) {
	return GetClient(client).MessageApi.PostMessage(GetContext(client), channelId).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
		Embed:   &embed,
	}).Execute()
}

func EditMessage(client int, messageId string, content string) (*http.Response, error) {
	return GetClient(client).MessageApi.EditMessage(GetContext(client), messageId).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
}

func SearchMessages(word string, after, before time.Time, in, to, from, citation string, bot, hasURL, hasAttachments, hasImage, hasVideo, hasAudio bool, limit, offset int32, sort string) (*traq.MessageSearchResult, *http.Response, error) {
	return GetClient(Bot).MessageApi.SearchMessages(GetContext(Bot)).Word(word).After(after).Before(before).In(in).To(to).From(from).Citation(citation).Bot(bot).HasURL(hasURL).HasAttachments(hasAttachments).HasImage(hasImage).HasVideo(hasVideo).HasAudio(hasAudio).Limit(limit).Offset(offset).Sort(sort).Execute()
}

func SearchMessagesByUnread(after time.Time, channelId string) (*traq.MessageSearchResult, *http.Response, error) {
	return GetClient(Bot).MessageApi.SearchMessages(GetContext(Bot)).After(after).In(channelId).Limit(100).Sort("-createdAt").Execute()
}

func SearchMessagesByTime(after time.Time, before time.Time) (*traq.MessageSearchResult, *http.Response, error) {
	return GetClient(Bot).MessageApi.SearchMessages(GetContext(Bot)).After(after).Before(before).Limit(100).Bot(false).Execute()
}
