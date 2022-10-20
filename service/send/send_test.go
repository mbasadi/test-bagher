package send

import "testing"
	
func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello, World2" {
		t.Errorf("HelloWorlf = %s, want \"Hello, World\"", HelloWorld())
	}
}
	