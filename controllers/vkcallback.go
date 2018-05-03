package controllers

import (
    "net/http"
    "encoding/json"
    "bytes"
    "strconv"
    "github.com/astaxie/beego"
    "vknotifier/models"
    "net/url"
)

var communities map[int]string
var communityTypes map[int]string

type VkCallbackController struct {
    beego.Controller
}

func (c *VkCallbackController) Post() {
    if len(communities) == 0 {
        communities, communityTypes = models.GetCommunitiesMap()
    }

    var msg models.VkCallbackRequest
    json.Unmarshal(c.Ctx.Input.RequestBody, &msg)

    var confirmationResponse, message string
    confirmationResponse = ""

    communityName, issetCommunityName := communities[msg.GroupID]
    if issetCommunityName != true {
        message = "в группе " + strconv.Itoa(msg.GroupID) + " "
    } else {
        message = "в группе " + communityName + " "
    }

    if msg.Type == "confirmation" {
        confirmationResponse = models.ConfirmationResponse(msg)
    } else if msg.Type == "photo_comment_new" {
        message += "добавлен комментарий к фотографии http://vk.com/" + communityTypes[msg.GroupID]
        message += strconv.Itoa(msg.GroupID) + "?z=photo-" + strconv.Itoa(msg.GroupID)
        message += "_" + strconv.Itoa(msg.Object.PhotoID)
    } else if msg.Type == "wall_post_new" {
        message += "добавлен пост на стену https://vk.com/wall-" + strconv.Itoa(msg.GroupID) + "_"
        message += strconv.Itoa(msg.Object.ID)
    } else if msg.Type == "wall_reply_new" {
        message += "добавлен комментарий в пост https://vk.com/wall-" + strconv.Itoa(msg.GroupID) + "_"
        message += strconv.Itoa(msg.Object.ID)
    } else if msg.Type == "message_new" {
        message += "новое сообщение «" + msg.Object.Body + "»"
    } else if msg.Type == "board_post_new" {
        message += "новое cообщение в обсуждениях «" + msg.Object.Text + "» https://vk.com/topic-"
        message += strconv.Itoa(msg.GroupID) + "_" + strconv.Itoa(msg.Object.TopicID)
        message += "?post=" + strconv.Itoa(msg.Object.ID)
    } else {
        message = "в группе " + strconv.Itoa(msg.GroupID) + " произошло " + msg.Type
    }

    if len(confirmationResponse) > 1 {
        c.Data["response"] = confirmationResponse
        c.TplName = "plainResponse.html"
    } else {
        SendMessage(message)
        c.Data["response"] = "ok"
        c.TplName = "plainResponse.html"
    }
}

func SendMessage(message string) {
    apiUrl := "https://api.telegram.org/bot" + beego.AppConfig.String("botId") + ":"
    apiUrl += beego.AppConfig.String("botToken") + "/sendMessage"
    var jsonStr = []byte(`{"chat_id":"-` + beego.AppConfig.String("chatId") + `","text":"` + message + `"}`)

    req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    proxyUrl, err := url.Parse("socks5://127.0.0.1:9050")
    client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    // TODO get full body
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
}
