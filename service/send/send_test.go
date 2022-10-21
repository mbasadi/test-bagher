package send

import "testing"
	
func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello, World7" {
		t.Errorf("HelloWorlf = %s, want \"Hello, World\"", HelloWorld())
	}
}
	