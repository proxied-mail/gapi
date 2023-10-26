package entityId

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encoder := &Encoder{}
	encoded := encoder.Encode(12345, "exampleEntity")
	if encoded != "3B9D9303-0000-0000-0000A0D3" {
		t.Error("encoder failed")
	}

	encoded2 := encoder.Encode(380915, "proxy_bindings")
	if encoded2 != "2FFF3FFC-5000-0000-00000BAE" {
		t.Error("encoder2 failed")
	}
}

func TestEncode2(t *testing.T) {
	encoder := &Encoder{}
	encoded2 := encoder.Encode(380915, "proxy_bindings")
	if encoded2 != "2FFF3FFC-5000-0000-00000BAE" {
		t.Error("encoder2 failed")
	}
}

func TestDecode(t *testing.T) {
	encoder := &Encoder{}
	encoded, _ := encoder.Decode("3B9D9303-0000-0000-0000A0D3", "exampleEntity")
	if encoded != 12345 {
		t.Error("Decoder failed")
	}
}
