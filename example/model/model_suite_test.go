package model

import (
	"testing"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redis-orm")
}
