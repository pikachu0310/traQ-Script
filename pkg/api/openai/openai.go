package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/sashabaranov/go-openai"
)

type Models int

const (
	GPT3dot5Turbo Models = iota
	GPT4
)

type FinishReason int

const (
	stop FinishReason = iota
	length
	errorHappen
)

var blobs = [...]string{":blob_bongo:", ":blob_crazy_happy:", ":blob_grin:", ":blob_hype:", ":blob_love:", ":blob_lurk:", ":blob_pyon:", ":blob_pyon_inverse:", ":blob_slide:", ":blob_snowball_1:", ":blob_snowball_2:", ":blob_speedy_roll:", ":blob_speedy_roll_inverse:", ":blob_thinking:", ":blob_thinking_fast:", ":blob_thinking_portal:", ":blob_thinking_upsidedown:", ":blob_thonkang:", ":blob_thumbs_up:", ":blobblewobble:", ":blobenjoy:", ":blobglitch:", ":blobbass:", ":blobjam:", ":blobkeyboard:", ":bloblamp:", ":blobmaracas:", ":blobmicrophone:", ":blobthinksmart:", ":blobwobwork:", ":conga_party_thinking_blob:", ":Hyperblob:", ":party_blob:", ":partyparrot_blob:", ":partyparrot_blob_cat:"}
var amazed = [...]string{":amazed_fuzzy:", ":amazed_amazed_fuzzy:", ":amazed_god_enel:", ":amazed_hamutaro:"}
var blobsAndAmazed = append(blobs[:], amazed[:]...)
var warnings = [...]string{":warning:", ":ikura-hamu_shooting_warning:"}

type Message = openai.ChatCompletionMessage

var Messages = make([]Message, 0)
var SystemRoleMessage = "あなたはサークル である東京工業大学デジタル創作同好会traPの部内SNS、traQのユーザーを楽しませる娯楽用 途のBOTの中に作られた、openaiのモデルgpt-3.5-turboを用いた対話型AIです。身内しかいないSNSで、ユーザーに緩く接してください。そ して、ユーザーの言う事に出来る限り従うようにしてください。"

const deltaTime = 500 * time.Millisecond
const GptSystemString = "FirstSystemMessageを変更しました。/gptsys showで確認できま す。\nFirstSystemMessageとは、常に履歴の一番最初に入り、最初にgptに情報や状況を説明 するのに使用する文字列です"

func Stream(messages []Message, openaiModel Models, do func(string)) (responseMessage string, finishReason FinishReason, err error) {
	c := Client
	ctx := context.Background()

	var model string
	switch openaiModel {
	case GPT3dot5Turbo:
		model = openai.GPT3Dot5Turbo
	case GPT4:
		model = openai.GPT4
	default:
		model = openai.GPT3Dot5Turbo
	}

	req := openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Stream:   true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	lastDoTime := time.Now()
	for {
		response, err := stream.Recv()

		if err != nil {
			do(responseMessage + warnings[rand.Intn(len(warnings))] + ":blobglitch: Error: " + fmt.Sprint(err))
			finishReason = errorHappen
			break
		}
		if errors.Is(err, io.EOF) {
			err = errors.New("stream closed")
			finishReason = errorHappen
			break
		}

		if response.Choices[0].FinishReason == "stop" {
			time.Sleep(200 * time.Millisecond)
			do(responseMessage)
			finishReason = stop
			break
		} else if response.Choices[0].FinishReason == "length" {
			do(responseMessage + "\n" + amazed[rand.Intn(len(amazed))] + "トークン(履歴を含む文字数)が上限に達したので履歴の最初のメッセージを削除して続きを出力します:loading:")
			finishReason = length
			break
		}

		responseMessage += response.Choices[0].Delta.Content

		if time.Now().Sub(lastDoTime) >= deltaTime {
			lastDoTime = time.Now()
			do(blobsAndAmazed[rand.Intn(len(blobs))] + responseMessage + ":loading:")
		}
	}
	AddMessageAsAssistant(responseMessage)
	return
}

func AddMessageAsUser(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
}

func AddMessageAsAssistant(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleAssistant,
		Content: message,
	})
}

func AddMessageAsSystem(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	})
}

func AddSystemMessageIfNotExist(message string) {
	for _, m := range Messages {
		if m.Role == "system" {
			return
		}
	}
	Messages = append([]Message{{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	}}, Messages...)
}

func UpdateSystemRoleMessage(message string) {
	AddSystemMessageIfNotExist(message)
	Messages[0] = Message{
		Role:    "system",
		Content: message,
	}
}

func ChatReset() {
	Messages = make([]Message, 0)
}
