package send

import "testing"
	
func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello, World3" {
		t.Errorf("HelloWorlf = %s, want \"Hello, World\"", HelloWorld())
	}
}
	