package process

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)


type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"` //消息ID
	UserId  int64  `json:"userid,omitempty" form:"userid"` //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` //群聊还是私聊
	DstId   int64  `json:"dstId,omitempty" form:"dstid"`//对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"` //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"` //预览图片
	Url     string `json:"url,omitempty" form:"url"` //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"` //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"` //其他和数字相关的
}

type Node struct {
	Conn *websocket.Conn

}

var clientMap map[int64]*Node = make(map[int64]*Node,0)

var rwLocker sync.RWMutex



func Chat(writer http.ResponseWriter,request *http.Request) {

}




