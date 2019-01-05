Cooldown Timer

### Usage
from the tests:
``` golang
  cd100 := NewCooldownTimer(100) // goes off after 100 is subtracted
  cd100.CoolDown(100) // subtract 100 from the timer
  if cd100.CheckAndReset() {
     t.Errorf("Timer went off and it shouldn't have...")
  }
  cd100.CoolDown(1)
  if !cd100.CheckAndReset() {
     t.Errorf("Timer didn't go off and it should have...")
  }
```

Doesn't use any units because it's cranked manually, interpret the units as anything you want.
