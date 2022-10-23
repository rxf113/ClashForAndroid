package statistic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"os"
	"sync"
)

type Dto struct {
	Download      int64  `json:"download"`
	RemoteAddress string `json:"address"`
	ExitTime      int64  `json:"exitTime"`
}

var wg sync.WaitGroup

var conn2 *websocket.Conn = nil

var url = "ws://127.0.0.1:8888/ws"

func initClient() {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("错误信息:", err)
	}
	conn2 = conn
	wg.Wait()
}

func SendMsg(dto *Dto) {
	bytes, err := json.Marshal(dto)
	if err != nil {
		log.Println(err)
	}

	if conn2 == nil {
		initClient()
	}
	conn2.WriteMessage(1, bytes)
}

func read(conn *websocket.Conn) {
	defer wg.Done()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("错误信息:", err)
			break
		}
		if err == io.EOF {
			continue
		}
		fmt.Println("获取到的信息:", string(msg))
	}
}
func writeM(conn *websocket.Conn) {
	defer wg.Done()
	for {
		fmt.Print("请输入:")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		conn.WriteMessage(1, []byte(data))
	}
}
