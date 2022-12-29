## apdu_send 使用
针对在wsl启动`speculos`后，并在windows 浏览器能通过地址正确访问该程序的apdu指令发送页面。

use:

```
go mod download github.com/DengJHeng/speculos_apdu_send
```

```go
package main

import "github.com/DengJHeng/speculos_apdu_send/apdu"

func main(){
	apduBuffer:=make([]byte,0)
	apduObj := apdu.NewApdu("http://localhost:5000")
	// send apdu buffer
	apduObj.SendMsg(apduBuffer)
}
```