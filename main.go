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
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			case *linebot.ImageMessage:
				if err := app.handleImage(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				if err := app.handleVideo(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.AudioMessage:
				if err := app.handleAudio(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.FileMessage:
				if err := app.handleFile(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if err := app.handleLocation(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if err := app.handleSticker(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyText(event.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			if err := app.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			if err := app.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

		case "e2e":
			switch message := event.Message.(type) {
				return nil
			}
			messages = append(messages,
				linebot.NewTextMessage("クーポンをゲットしよう!!!"),
				linebot.NewTextMessage(os.Getenv("WEB_CAMPAIGN_URL")),
			)
		}

		return messages
	}

	//石斑魚的code
						case Contains(message.Text,"a1")||Contains(message.Text,"班"):
							food = ""
							switch{
								case Contains(message.Text,"a2"):
									food = "龍虎石斑"
								case Contains(message.Text,"a01"):
									food = "青斑"
								case Contains(message.Text,"a02"):
									food = "珍珠石斑"
								default:
									if Contains(message.Text,"一斤多少")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算")||Contains(message.Text,"還有多少")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"){
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("拍謝啦! 我是笨笨的電腦，不知道您要問哪種石斑捏，我們有龍虎石斑、青斑、還有珍珠石斑")).Do()
									}
							}
							if food != ""{
								i:=0
								for i<=len(list_array){
									var menu []string
									menu = strings.Split(list_array[i], " ")
									if menu[0] == food{
										price=menu[1]
										stock=menu[2]
										break
									}
									i++
								}
								switch{
									case Contains(message.Text,"w5")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
									case Contains(message.Text,"w6")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "大概還有" + stock + "尾可以買，賣完就沒了喔!! 趕快來電088953096/0939220743黃先生")).Do() 
									case Contains(message.Text,"w7"):
										if len(profile) > 0{
											bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買" + food)).Do() 
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do()
										}else{
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
										}**/
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買石斑嗎? 趕快撥打 0939 220 743 跟石斑膠膠哥-黃大哥買魚喔!!")).Do()
								}
							}else{
								if Contains(message.Text,"w4"){
									if len(profile) > 0{
										bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買菜")).Do() 
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do() 	
									}else{
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
									}**/
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買石斑嗎? 趕快撥打 0939 220 743 跟石斑膠膠哥-黃大哥買魚喔!!")).Do()
								}
							}

	//以下是喇賽的code
						case Contains(message.Text,"w1")||Contains(message.Text,"洗眼")||Contains(message.Text,"乳牛")||Contains(message.Text,"淨化"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
						case Contains(message.Text,"w2"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "6569950-1490833625.jpg", url + "6569950-1490833625.jpg")).Do() 
						case Contains(message.Text,"w3"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "blackman.jpg", url + "blackman.jpg")).Do() 					
					}
				}
			}

		}
	}
}

if text == "Buttons" {
	    message := linebot.NewTextMessage(text + "じゃねぇよ！")
			linebot.NewTextMessage("Select your favorite food category or send me your location!").
			}
		}
		case "e3e":
			if e.Source.Type != linebot.EventSourceTypeUser {
				return nil
			}
			messages = append(messages,
				linebot.NewTextMessage("クーポンをゲットしよう!!!"),
				linebot.NewTextMessage(os.Getenv("WEB_CAMPAIGN_URL")),
			)
		}

		return messages
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
