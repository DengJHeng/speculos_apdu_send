package apdu

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Apdu struct {
	Address string
}

type MsgData struct {
	Data string `json:"data"`
}

func NewApdu(address string) *Apdu {
	return &Apdu{
		address,
	}
}

func (a *Apdu) SendMsg(msg []byte) (res []byte, err error) {
	if msg == nil || len(msg) == 0 {
		err = fmt.Errorf("msg is nil")
		return
	}
	msgData := MsgData{Data: hex.EncodeToString(msg)}
	msgJsonStr, err := json.Marshal(msgData)
	if err != nil {
		err = fmt.Errorf("json marshal msgData err:%v", err)
		return
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%v/apdu", a.Address), bytes.NewBuffer(msgJsonStr))
	if err != nil {
		err = fmt.Errorf("http.NewRequest err:%v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("send msg to server err:%v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read body data err:%v", err)
		return
	}
	replyData := MsgData{}
	err = json.Unmarshal(body, &replyData)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal reply data err:%v", err)
		return
	}
	dataByte, err := hex.DecodeString(replyData.Data)
	if err != nil {
		err = fmt.Errorf("hex.DecodeString reply data err:%v", err)
		return
	}

	swOffset := len(dataByte) - 2
	sw := binary.BigEndian.Uint16(dataByte[swOffset:])
	if sw != 0x9000 {
		err = fmt.Errorf("apdu reply data code not equal 9000")
		res = dataByte[:swOffset]
		return
	}
	res = dataByte[:swOffset]
	return
}
