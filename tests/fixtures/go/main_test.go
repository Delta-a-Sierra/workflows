package fixture

import "testing"

func TestGreeting(t *testing.T) {
	t.Parallel()

	if got := Greeting(); got != "workflow hub fixture" {
		t.Fatalf("Greeting() = %q, want %q", got, "workflow hub fixture")
	}
}
