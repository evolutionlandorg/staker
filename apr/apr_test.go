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
	"math/big"
	"strings"
	"testing"
)

type MockBackend struct {
}

// StakingRewards.sol
func (m *MockBackend) RewardsToken(pool string) (string, error) {
	return "0xb52fbe2b925ab79a821b261c82c5ba0814aaa5e0", nil
}

func (m *MockBackend) StakingToken(pool string) (string, error) {
	return "0x7f34ee2a450fc0838955da3f278262da0dea0c67", nil
}

func (m *MockBackend) PeriodFinish(pool string) (int64, error) {
	return 3619914394, nil
}

func (m *MockBackend) RewardRate(pool string) (*big.Int, error) {
	return U256("0x3157def08c0e38"), nil
}

// UniswapV2Pair.sol
func (m *MockBackend) PairBalanceOf(lpToken string, pool string) (*big.Int, error) {
	return U256("0x1acfad219a54b20000"), nil
}

func (m *MockBackend) TotalSupply(lpToken string) (*big.Int, error) {
	return U256("0x29e49d9a95461f0000"), nil
}

func (m *MockBackend) Token0(lpToken string) (string, error) {
	return "0xb52fbe2b925ab79a821b261c82c5ba0814aaa5e0", nil
}

func (m *MockBackend) Token1(lpToken string) (string, error) {
	return "0xb6a07a36fa73758ce9d58a2c6a8da74cecca438d", nil
}

func (m *MockBackend) GetReserves(lpToken string) (reserve0 *big.Int, reserve1 *big.Int, blockTimestampLast int64, err error) {
	return U256("0x5d910fcf4294380000"), U256("0x12bc13a49741276000"), 0, nil
}

func U256(v string) *big.Int {
	v = strings.TrimPrefix(v, "0x")
	bn := new(big.Int)
	n, _ := bn.SetString(v, 16)
	return n
}

func TestAPR(t *testing.T) {
	b := &MockBackend{}
	Apr := New(b, 2)
	apr, err := Apr.Calc("0x0CE88f43bFA23F2627757267B4EEc8f1911e2d75", "0xb52fbe2b925ab79a821b261c82c5ba0814aaa5e0")
	if err != nil {
		t.Failed()
	}
	t.Log("apr: ", apr)
	var expectedAPR = 19825.68
	if apr != expectedAPR {
		t.Errorf("Got %v expected %v", apr, expectedAPR)
	}
}
