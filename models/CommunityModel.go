package models

import (
    "github.com/astaxie/beego/orm"
    "strconv"
)

func init() { orm.RegisterModel(new(Community)) }

func (u *Community) TableName() string {
    return "communities"
}

type Community struct {
    Id            int
    VkId          int
    Name          string
    Type          string
    Key           string
    CheckResponse string
}

type VkCallbackRequest struct {
    Type string `json:"type"`
    Object struct {
        ID           int    `json:"id"`
        FromID       int    `json:"from_id,omitempty"`
        Date         int    `json:"date,omitempty"`
        Text         string `json:"text,omitempty"`
        PostOwnerID  int    `json:"post_owner_id,omitempty"`
        PostID       int    `json:"post_id,omitempty"`
        PhotoOwnerID int    `json:"photo_owner_id,omitempty"`
        PhotoID      int    `json:"photo_id,omitempty"`
        TopicID      int    `json:"topic_id,omitempty"`
        Body         string `json:"body,omitempty"`
    } `json:"object,omitempty"`
    GroupID int    `json:"group_id"`
    Secret  string `json:"secret,omitempty"`
}

var community Community

func GetCommunities() []Community {
    var communities []Community
    o := orm.NewOrm()
    o.QueryTable("communities").All(&communities)

    return communities
}

func ConfirmationResponse(msg VkCallbackRequest) string {
    o := orm.NewOrm()
    o.QueryTable("communities").Filter("VkId", msg.GroupID).Filter("Key", msg.Secret).One(&community)

    return community.CheckResponse
}

func GetCommunity(id int) Community {
    var community Community
    var o = orm.NewOrm()
    o.QueryTable("communities").Filter("Id", id).One(&community)

    return community
}

func DelCommunity(id int) {
    var o = orm.NewOrm()
    o.QueryTable("communities").Filter("Id", id).Delete()
}

func GetCommunitiesMap() (map[int]string, map[int]string) {
    var communities map[int]string
    var communityTypes map[int]string
    communities = make(map[int]string)
    communityTypes = make(map[int]string)
    query := GetCommunities()
    for _, v := range query {
        communities[v.VkId] = v.Name
        communityTypes[v.VkId] = v.Type
    }

    return communities, communityTypes
}

func EditCommunity(commuintyRow map[string]string) bool {
    var o = orm.NewOrm()
    var community Community
    o.QueryTable("communities").Filter("Id", commuintyRow["Id"]).One(&community)
    community.VkId, _ = strconv.Atoi(commuintyRow["communityId"])
    community.Name = commuintyRow["communityName"]
    community.CheckResponse = commuintyRow["responseCode"]
    community.Type = commuintyRow["communityType"]
    community.Key = commuintyRow["secretKey"]

    o.Update(&community)

    return true
}