package patterns

import (
	"reflect"
	"testing"
)

func TestNewServerDiscarded(t *testing.T) {
	got := NewServerDiscarded("localhost", 10000, 10000, false)
	want := &server{
		"localhost",
		default_port,
		default_timeout,
		default_encryption,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost, 1000, 1000, false) = %v, want %v", got, want)
	}
}

func TestNewServer(t *testing.T) {
	got := NewServer("localhost")
	want := &server{
		"localhost",
		default_port,
		default_timeout,
		default_encryption,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost) = %v, want %v", got, want)
	}
}

func TestWithEncryption(t *testing.T) {
	got := NewServer("localhost", WithEncryption())
	want := &server{
		"localhost",
		default_port,
		default_timeout,
		true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost, WithEncryption()) = %v, want %v", got, want)
	}
}

func TestWithPort(t *testing.T) {
	got := NewServer("localhost", WithPort(65535))
	want := &server{
		"localhost",
		65535,
		default_timeout,
		false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost, WithPort(65535)) = %v, want %v", got, want)
	}
}

func TestWithTimeoutMs(t *testing.T) {
	got := NewServer("localhost", WithTimeoutMs(1000))
	want := &server{
		"localhost",
		default_port,
		1000,
		false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost, WithTimeoutMs(1000)) = %v, want %v", got, want)
	}
}

func TestWithoutEncryption(t *testing.T) {
	got := NewServer("localhost", WithoutEncryption())
	want := &server{
		"localhost",
		default_port,
		default_timeout,
		false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewServer(localhost, WithoutEncryption()) = %v, want %v", got, want)
	}
}
