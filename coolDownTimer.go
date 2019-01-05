package utils

import(
//  "math"
 "fmt"
)

type CooldownTimer struct{
  TimeLeftMillis float64
  Interval float64
  GoneOff bool
}

func New(interval float64) *CooldownTimer {
  return &CooldownTimer{interval, interval, false}
}


func (c *CooldownTimer) PhaseShift(delta float64){
  (*c).TimeLeftMillis += delta;
}

func (c *CooldownTimer) CoolDown(coolDown float64){
  interval := (*c).Interval
  timeLeft := &((*c).TimeLeftMillis)
  *timeLeft -= coolDown;
  if(*timeLeft < -interval){
    fmt.Println("uh-oh, we're badly behind", *timeLeft - interval)
  }
  for *timeLeft < 0.0{
    *timeLeft += interval
    (*c).GoneOff = true
  }
}

func (c *CooldownTimer) CheckAndReset() bool{
  if((*c).GoneOff){
    (*c).GoneOff = false;
    return true;
  }
  return false;
}
