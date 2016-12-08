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
	"net/http"
)

func main() {
	http.HandleFunc("/postpage", func(w http.ResponseWriter, r *http.Request) {
		//接受post请求，然后打印表单中key和value字段的值
		if r.Method == "POST" {
			var (
				key   string = r.PostFormValue("key")
				value string = r.PostFormValue("value")
			)
			fmt.Printf("key is  : %s\n", key)
			fmt.Printf("value is: %s\n", value)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
