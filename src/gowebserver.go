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
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func WebServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", loginTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("192.168.1.27:8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")

	//模拟延时
	time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	result := NewBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}
