package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func Get[T Employee](value T) string {
	return value.GetName()
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return m.Name
}

func TestInheritance(t *testing.T) {
	manager := &Manager{
		Name: "fahmi",
	}

	assert.Equal(t, "fahmi", Get(manager))
}