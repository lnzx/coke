package api

import (
	"fmt"
)

const (
	list_node   = "headscale nodes list -o json-line"
	rename_node = "headscale nodes rename %s -i %d"
	delete_node = "headscale nodes delete -i %d --force"
	list_users  = "headscale users list -o json"
)

func ListUsers() string {
	return run(list_users)
}

func GetNodes(user string) string {
	if user == "" {
		return run(list_node)
	}
	command := list_node + " -u " + user
	return run(command)
}

func RenameNode(newName string, id int) string {
	command := fmt.Sprintf(rename_node, newName, id)
	return run(command)
}

func DeleteNode(id int) string {
	command := fmt.Sprintf(delete_node, id)
	return run(command)
}
