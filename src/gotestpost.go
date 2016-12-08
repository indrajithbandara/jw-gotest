package src

/**
 *@ProjectName: ${project_name}
 *@FileName: ${file_name}
 *@Date: ${date}-${time}
 *@Copyright: ${year} www.jointwisdom.com Inc. All rights reserved.

 * Copyright (c) 2016 jointwisdom Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
*/
import (
  "fmt"
  "net"
)

func main() {
	//因为post方法属于HTTP协议，HTTP协议以TCP为基础，所以先建立一个
	//TCP连接，通过这个TCP连接来发送我们的post请求
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	//构造post请求
	var post string
	post += "POST /postpage HTTP/1.1\r\n"
	post += "Content-Type: application/x-www-form-urlencoded\r\n"
	post += "Content-Length: 37\r\n"
	post += "Connection: keep-alive\r\n"
	post += "Accept-Language: zh-CN,zh;q=0.8,en;q=0.6\r\n"
	post += "\r\n"
	post += "key=this is key&value=this is value\r\n"

	if _, err := conn.Write([]byte(post)); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("post send success.")
}
