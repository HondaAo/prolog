package discovery

import (
	"fmt"
	"testing"
)

func setupMember(t testing.T, members []*Membership) ([]*Membership, *handler) {
	id := len(members)
	ports := dynaport.Get(1)
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", ports[0])
}

type handler struct {
	joins chan map[string]string
	leaves chan string
}