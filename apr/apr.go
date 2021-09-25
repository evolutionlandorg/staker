/*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 2 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, see <http://www.gnu.org/licenses/>.
*
* Copyright (C) echo, 2021
 */
package apr

import (
	"errors"
	"math"
	"math/big"
	"strings"
	"time"
)

var Big0 = big.NewInt(0)
var Big1 = big.NewInt(1)
var Big18250 = big.NewInt(18250)

type APR struct {
	b       Backend
	decimal int
}

func New(b Backend, decimal int) *APR {
	return &APR{b, decimal}
}

// APR(%) = 365 * 100 * dailyReward / ( 2 * reverseRing)
// APR(%) = 18250 * dailyReward / reserveRingInPool
func (a *APR) Calc(pool, base, ring string, transformer *Fraction) (float64, error) {
	rewardToken, err := a.b.RewardsToken(pool)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	if !strings.EqualFold(rewardToken, ring) {
		return math.SmallestNonzeroFloat64, errors.New("Not support")
	}
	reserveBaseInPool, err := a.getReserveBaseInPool(pool, base)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	reserveRingInPool := new(Fraction).Mul(reserveBaseInPool, transformer)
	dailyReward, err := a.getDailyReward(pool)
	if err != nil {
		return math.SmallestNonzeroFloat64, err
	}
	apr := new(Fraction).Div(NewFraction(new(big.Int).Mul(Big18250, dailyReward), Big1), reserveRingInPool)
	return apr.toFixed(a.decimal), nil
}

func (a *APR) getReserveBaseInPool(pool, base string) (*Fraction, error) {
	lpToken, err := a.b.StakingToken(pool)
	if err != nil {
		return nil, err
	}
	var reserveBase *big.Int
	reserve0, reserve1, _, err := a.b.GetReserves(lpToken)
	if err != nil {
		return nil, err
	}
	token0, err := a.b.Token0(lpToken)
	if err != nil {
		return nil, err
	}
	token1, err := a.b.Token1(lpToken)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(base, token0) {
		reserveBase = reserve0
	} else if strings.EqualFold(base, token1) {
		reserveBase = reserve1
	} else {
		return nil, errors.New("RING not in pair")
	}
	totalStakedLPAmount, err := a.b.PairBalanceOf(lpToken, pool)
	if err != nil {
		return nil, err
	}
	totalLPAmount, err := a.b.TotalSupply(lpToken)
	if err != nil {
		return nil, err
	}
	return &Fraction{
		numerator:   new(big.Int).Mul(reserveBase, totalStakedLPAmount),
		denominator: totalLPAmount,
	}, nil
}

func (a *APR) getDailyReward(pool string) (*big.Int, error) {
	isOver, err := a.isRewardPeriodOver(pool)
	if err != nil {
		return Big0, err
	}
	if isOver {
		return Big0, nil
	}
	rewardRate, err := a.b.RewardRate(pool)
	if err != nil {
		return Big0, err
	}
	return new(big.Int).Mul(big.NewInt(86400), rewardRate), nil
}

func (a *APR) isRewardPeriodOver(pool string) (bool, error) {
	now := time.Now().UTC().Unix()
	periodFinish, err := a.b.PeriodFinish(pool)
	if err != nil {
		return true, err
	}
	return periodFinish < now, nil
}
