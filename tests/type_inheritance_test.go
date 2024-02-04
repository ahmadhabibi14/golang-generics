package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePrecident struct {
	Name string
}

func (m *MyVicePrecident) GetName() string {
	return m.Name
}

func (m *MyVicePrecident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Habi", GetName[Manager](&MyManager{Name: "Habi"}))
	assert.Equal(t, "Habi", GetName[VicePresident](&MyVicePrecident{Name: "Habi"}))
}
