package mongoorg

import "testing"

func TestClient(t *testing.T) {
	e := NewEngine()
	err := e.Init()
	if err != nil {
		t.Error(err)
	}
}
