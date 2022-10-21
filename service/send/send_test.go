package send

import "testing"
	
func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello, World8" {
		t.Errorf("HelloWorlf = %s, want \"Hello, World\"", HelloWorld())
	}
}
	