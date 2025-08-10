package service

// IDGeneratorはIDを生成する機能のインターフェース
type IDGenerator interface {
	NewID() string
}

