package main

import "fmt"

type HttpRequest struct {
	count    int    // 1:代表1个,2:代表2部分,3:代表3部分
	position [3]int // 46代表左右,28代表上下,258代表上中下,456代表左中右,246代表上左下右下,135代表下左上右上
	parts    [3]rune
}

type HttpResponse struct {
	code int
	msg  string
}

type ServerNode struct {
	c              rune
	UpNodes        []*ServerNode
	DownNodes      []*ServerNode
	LeftNodes      []*ServerNode
	RightNodes     []*ServerNode
	LeftDownNodes  []*ServerNode
	RightDownNodes []*ServerNode
	LeftUpNodes    []*ServerNode
	RightUpNodes   []*ServerNode
}

func IsValid(count int, position [3]int, parts [3]rune) bool {
	httpRequest := HttpRequest{
		count:    count,
		position: position,
		parts:    parts,
	}
	httpResponse := ServerIsValid(httpRequest)
	if httpResponse.code == 200 {
		return true
	}
	return false
}

var m = make(map[rune]*ServerNode)

func ServerInsert(httpRequest HttpRequest) {
	count := httpRequest.count
	position := httpRequest.position
	parts := httpRequest.parts
	if count == 2 {
		// 左右
		if position[0] == 4 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			m[parts[0]].RightNodes = append(m[parts[0]].RightNodes, m[parts[1]])
		}
		// 上下
		if position[0] == 2 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			m[parts[0]].DownNodes = append(m[parts[0]].DownNodes, m[parts[1]])
		}
	} else if count == 3 {
		// 上中下
		if position[0] == 2 && position[1] == 5 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			if m[parts[2]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[2]
				m[parts[2]] = &newNode
			}
			m[parts[0]].DownNodes = append(m[parts[0]].DownNodes, m[parts[1]])
			m[parts[1]].DownNodes = append(m[parts[1]].DownNodes, m[parts[2]])
			// 左中右
		} else if position[0] == 4 && position[1] == 5 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			if m[parts[2]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[2]
				m[parts[2]] = &newNode
			}
			m[parts[0]].RightNodes = append(m[parts[0]].RightNodes, m[parts[1]])
			m[parts[1]].RightNodes = append(m[parts[1]].RightNodes, m[parts[2]])
			// 上左下右下
		} else if position[0] == 2 && position[1] == 4 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			if m[parts[2]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[2]
				m[parts[2]] = &newNode
			}
			m[parts[0]].LeftDownNodes = append(m[parts[0]].LeftDownNodes, m[parts[1]])
			m[parts[0]].RightDownNodes = append(m[parts[0]].RightDownNodes, m[parts[2]])
			// 下左上右上
		} else if position[0] == 1 && position[1] == 3 {
			if m[parts[0]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[0]
				m[parts[0]] = &newNode
			}
			if m[parts[1]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[1]
				m[parts[1]] = &newNode
			}
			if m[parts[2]] == nil {
				newNode := ServerNode{}
				newNode.c = parts[2]
				m[parts[2]] = &newNode
			}
			m[parts[0]].LeftUpNodes = append(m[parts[0]].LeftUpNodes, m[parts[1]])
			m[parts[0]].RightUpNodes = append(m[parts[0]].RightUpNodes, m[parts[2]])
		}
	}
}

func ServerIsValid(httpRequest HttpRequest) HttpResponse {
	count := httpRequest.count
	position := httpRequest.position
	parts := httpRequest.parts
	httpResponse := HttpResponse{}
	httpResponse.code = 400
	httpResponse.msg = "匹配错误"
	if count == 2 {
		if m[parts[0]] == nil || m[parts[1]] == nil {
			return httpResponse
		}
		// 左右
		if position[0] == 4 {
			if containItem(m[parts[1]], m[parts[0]].RightNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
		}
		// 上下
		if position[0] == 2 {
			if containItem(m[parts[1]], m[parts[0]].DownNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
		}
	} else if count == 3 {
		if m[parts[0]] == nil || m[parts[1]] == nil || m[parts[2]] == nil {
			return httpResponse
		}
		// 上中下
		if position[0] == 2 && position[1] == 5 {
			if containItem(m[parts[1]], m[parts[0]].DownNodes) && containItem(m[parts[2]], m[parts[1]].DownNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
			// 左中右
		} else if position[0] == 4 && position[1] == 5 {
			if containItem(m[parts[1]], m[parts[0]].RightNodes) && containItem(m[parts[2]], m[parts[1]].RightNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
			// 上左下右下
		} else if position[0] == 2 && position[1] == 4 {
			if containItem(m[parts[1]], m[parts[0]].LeftDownNodes) && containItem(m[parts[2]], m[parts[0]].RightDownNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
			// 下左上右上
		} else if position[0] == 1 && position[1] == 3 {
			if containItem(m[parts[1]], m[parts[0]].LeftUpNodes) && containItem(m[parts[2]], m[parts[0]].RightUpNodes) {
				httpResponse.code = 200
				httpResponse.msg = "匹配成功"
				return httpResponse
			}
		}
	}
	return httpResponse
}

func containItem(item *ServerNode, itemArray []*ServerNode) bool {
	flag := false
	for i := 0; i < len(itemArray); i++ {
		if itemArray[i].c == item.c {
			flag = true
			break
		}
	}
	return flag
}

func main() {
	httpRequest := HttpRequest{}
	httpRequest.count = 3
	httpRequest.position = [3]int{2, 4, 6}
	httpRequest.parts = [3]rune{'雨', '革', '月'}
	ServerInsert(httpRequest)
	fmt.Printf("霸：雨、革、月")
	fmt.Println(ServerIsValid(httpRequest))
	httpRequest.count = 2
	httpRequest.position = [3]int{2, 8}
	httpRequest.parts = [3]rune{'大', '亏'}
	ServerInsert(httpRequest)
	fmt.Printf("夸：大、亏")
	fmt.Println(ServerIsValid(httpRequest))
}
