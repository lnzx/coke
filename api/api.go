package api

import (
	"fmt"
	"github.com/bytedance/sonic"
)

const (
	list_node   = "headscale nodes list -o json-line"
	rename_node = "headscale nodes rename %s -i %d"
	delete_node = "headscale nodes delete -i %d --force"
	list_users  = "headscale users list -o json-line"
)

func ListUsers() []User {
	var users []User
	bytes := run(list_users)
	err := sonic.Unmarshal(bytes, &users)
	if err != nil {
		return []User{}
	}
	return users
}

func GetNodes(user string) []Node {
	var nodes []Node
	if user == "" {
		bytes := run(list_node)
		if err := sonic.Unmarshal(bytes, &nodes); err != nil {
			return []Node{}
		}
		return nodes
	}
	command := list_node + " -u " + user
	bytes := run(command)
	if err := sonic.Unmarshal(bytes, &nodes); err != nil {
		return nil
	}
	return nodes
}

func RenameNode(newName string, id int) []byte {
	command := fmt.Sprintf(rename_node, newName, id)
	return run(command)
}

func DeleteNode(id int) []byte {
	command := fmt.Sprintf(delete_node, id)
	return run(command)
}
