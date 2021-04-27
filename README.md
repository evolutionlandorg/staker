## Staker

Forked from 
[https://github.com/Uniswap/liquidity-staker](https://github.com/Uniswap/liquidity-staker)

Staking pool for Uniswap GOLD-RING, RING-WOOD, HHO-RING, FIRE-RING and SIOO-RING liquidity tokens for receiving X RING daily reward.

### API

#### `totalSupply()` 
Return the total liquidity token amount in the staking pool

#### `balanceOf(address account)`
Return the liquidity token balance of `account`

#### `earned(address account)`
Return the earned RING amount of `account`

#### `stake(uint256 amount)`
Stake `amount` liquidity token for receiving RING reward

#### `withdraw(uint256 amount)`
Withdraw `amount` liquidity token

#### `getReward()`
Claim earned RING

#### `exit()`
Withdraw all staked liquidity token and Claim earned RING to exit

