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

import "math/big"

type Backend interface {
	// StakingRewards.sol
	RewardsToken(pool string) (string, error)
	StakingToken(pool string) (string, error)
	PeriodFinish(pool string) (int64, error)
	RewardRate(pool string) (*big.Int, error)

	// UniswapV2Pair.sol
	BalanceOf(lpToken, pool string) (*big.Int, error)
	TotalSupply(lpToken string) (*big.Int, error)
	Token0(lpToken string) (string, error)
	Token1(lpToken string) (string, error)
	GetReserves(lpToken string) (reserve0, reserve1 *big.Int, err error)
}
