package goredis

import (
	"bytes"
	"strconv"
)

// ==============================
// 代表一条客户端指令
// 对于 SET name Latermoon
// cmd.StringAtIndex(0) == cmd.Name() == "SET"
// cmd.StringAtIndex(1) == "name"
// cmd.StringAtIndex(2) == "Latermoon"
// ==============================
type Command struct {
	Args [][]byte
}

// 指令名称
// cmd.StringAtIndex(0) == cmd.Name() == "SET"
func (cmd *Command) Name() string {
	return string(cmd.Args[0])
}

// 参数按字符串返回
func (cmd *Command) StringAtIndex(i int) string {
	return string(cmd.Args[i])
}

func (cmd *Command) IntAtIndex(i int) (j int) {
	j, _ = strconv.Atoi(cmd.StringAtIndex(i))
	return
}

func (cmd *Command) String() string {
	buf := bytes.Buffer{}
	for _, arg := range cmd.Args {
		buf.Write(arg)
		buf.WriteString(" ")
	}
	return buf.String()
}