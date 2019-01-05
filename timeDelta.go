package utils

import(
  "time"
)

type TimeDelta struct {
  lastTime float64
  currentTime float64
}

func NewTimeDelta() *TimeDelta{
  return &TimeDelta{
    float64(time.Now().UnixNano()),
    float64(time.Now().UnixNano())}
}

func (td *TimeDelta) CalculateDelta() float64{
  (*td).lastTime = (*td).currentTime
  (*td).currentTime = float64(time.Now().UnixNano())
  return ((*td).currentTime - (*td).lastTime) / 1000000.0
}
