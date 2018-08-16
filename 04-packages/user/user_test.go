package user

import (
	"testing"
)

func TestUserSayHello(t *testing.T) {
	s := SayHello()

	if s != "hello world!!" {
		t.Fail()
	}

}
