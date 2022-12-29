package apdu

import "testing"

func TestSendMsg(t *testing.T){
	apduObj:=NewApdu("http://172.28.76.42:5000")
	msg:=[]byte{0xb0,0x01,0x00,0x00,0x00}
	sendMsg, err := apduObj.SendMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("reply msg:%v",string(sendMsg))
}
