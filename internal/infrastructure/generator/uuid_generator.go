package generator

import (
	"OTakumi/todogo/internal/domain/service"

	"github.com/google/uuid"
)

// uuidGeneratorはIDGeneratorインターフェースの具象オブジェクトです
type uuidGenerator struct{}

// NewUUIDGeneratorはuuidGeneratorの新しいインスタンスを生成する
func NewUUIDGenerator() service.IDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) NewID() string {
	return uuid.NewString()
}
