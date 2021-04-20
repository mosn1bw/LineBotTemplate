// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
					case "m1": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好!我是自動回覆的機器人,在CoC台灣英雄聯盟為您服務")).Do()
					case "m2": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "m3": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大姓張不姓謝喔")).Do()
				}
			}
		}
	}
					
			case message.Text=="a7":
				if len(profile) > 0{
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
				}
				
			case "e3e" {
				messages := []linebot.Message{
				linebot.NewTextMessage("友達登録ありがとうございます。"),
			}
		}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
