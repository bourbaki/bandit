package bandit

import "math/rand"

type Arm interface {
	EReward() float64
}

type Bandit interface {
	Choose(arms []Arm) Arm
}

type EpsilonGreedy struct {
	epsilon float64
}

type ConstArm float64

func (c ConstArm) EReward() float64 {
	return float64(c)
}

func (b *EpsilonGreedy) Choose(arms []Arm) Arm {
	u := rand.Float64()
	if u < (1 - b.epsilon) {
		maxArm := arms[0]
		max := arms[0].EReward()
		for _, x := range arms[1:] {
			if x.EReward() > max {
				max = x.EReward()
				maxArm = x
			}
			return maxArm
		}
	}
	return arms[rand.Intn(len(arms))]
}
