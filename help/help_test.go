package help

import "testing"

func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()
	t.Log(uuid)
}
