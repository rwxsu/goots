package game

import "testing"

func TestHealthPercent(t *testing.T) {
	p := NewPlayer(2, "babu", Position{32000, 32000, 7})
	p.SetHealthMax(500)
	p.SetHealthNow(0)
	if HealthPercent(p) != 0 {
		t.Error("health percent is never 0")
	}
	p.SetHealthNow(500)
	if HealthPercent(p) != 100 {
		t.Error("health percent is never 100")
	}
	p.SetHealthNow(250)
	hp := HealthPercent(p)
	if hp != 50 {
		t.Errorf("expected 50 percent got %d percent", hp)
	}
}
