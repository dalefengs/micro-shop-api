package response

import (
	"fmt"
	"time"
)

type UserResponse struct {
	Id        uint32   `json:"id"`
	NickName  string   `json:"nickname"`
	Mobile    string   `json:"mobile"`
	Gender    int32    `json:"gender"`
	Birthday  JsonTime `json:"birthday"`
	OrangeKey string   `json:"orange_key"`
}

type UserTokenResponse struct {
	Id       uint   `json:"id"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Token    string `json:"token"`
}

// JsonTime 解析时间戳日期
type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	fmt.Printf("stap: %#v \n", stmp)
	return []byte(stmp), nil
}
