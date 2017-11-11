package wiper

import (
	"math"
	"math/rand"
	"fmt"
)

const (
	OFF = iota
	ON
)

const (
	LOW    = 30
	MEDIUM = 60
	HIGH   = 90
)

type Wiper struct {
	dropCount       int
	rainVelocity    int
	cleanEfficiency float64
	mode            int
}

func NewWiper(dropCount, rainVelocity int, cleanEfficiency float64) *Wiper {
	return &Wiper{dropCount, rainVelocity, cleanEfficiency, OFF}
}

func (wiper *Wiper) Step() {
	wiper.rain()
	wiper.setMode()
	if wiper.mode == ON {
		wiper.clean()
	}
	fmt.Printf("mode:%v\n", wiper.mode)
	fmt.Printf("drops:%v\n", wiper.dropCount)
}

func (wiper *Wiper) setMode() {
	isLow := wiper.isLow()
	isHigh := wiper.isHigh()
	if isLow == 1 {
		wiper.mode = OFF
	}
	if isHigh == 1 {
		wiper.mode = ON
	}
	random := rand.Float64()
	if isLow > random {
		wiper.mode = OFF
	} else if isHigh > random {
		wiper.mode = ON
	}
}

func (wiper *Wiper) clean() {
	wiper.dropCount = int(math.Floor(float64(wiper.dropCount) * wiper.cleanEfficiency))
}

func (wiper *Wiper) rain() {
	wiper.dropCount += wiper.rainVelocity
}

func (wiper *Wiper) isLow() float64 {
	if wiper.dropCount <= LOW {
		return 1
	} else if wiper.dropCount >= MEDIUM {
		return 0
	}
	return float64(MEDIUM-wiper.dropCount) / (MEDIUM - LOW)
}

func (wiper *Wiper) isHigh() float64 {
	if wiper.dropCount >= HIGH {
		return 1
	} else if wiper.dropCount <= MEDIUM {
		return 0
	}
	return float64(wiper.dropCount-MEDIUM) / (HIGH - MEDIUM)
}
