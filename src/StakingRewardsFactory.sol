pragma solidity ^0.5.16;

import 'zeppelin-solidity/token/ERC20/IERC20.sol';
import 'ds-auth/auth.sol';

import './StakingRewards.sol';

contract StakingRewardsFactory is DSAuth {
    // immutables
    address public rewardsToken;
    uint public stakingRewardsGenesis;

    // the staking tokens for which the rewards contract has been deployed
    address[] public stakingTokens;

    // rewards info by staking token
    mapping(address => address) public stakingRewardsInfoByStakingToken;

    constructor(
        address _rewardsToken,
        uint _stakingRewardsGenesis
    ) DSAuth() public {
        require(_stakingRewardsGenesis >= block.timestamp, 'StakingRewardsFactory::constructor: genesis too soon');

        rewardsToken = _rewardsToken;
        stakingRewardsGenesis = _stakingRewardsGenesis;
    }

    ///// permissioned functions

    // deploy a staking reward contract for the staking token, and store the reward amount
    // the reward will be distributed to the staking reward contract no sooner than the genesis
    function deploy(address stakingToken) public auth {
        require(stakingRewardsInfoByStakingToken[stakingToken] == address(0), 'StakingRewardsFactory::deploy: already deployed');

        stakingRewardsInfoByStakingToken[stakingToken] = address(new StakingRewards(/*_rewardsDistribution=*/ address(this), rewardsToken, stakingToken));
        stakingTokens.push(stakingToken);
    }

    function recoverERC20(address tokenAddress) public auth {
        for (uint i = 0; i < stakingTokens.length; i++) {
            address stakingRewards = stakingRewardsInfoByStakingToken[stakingTokens[i]];
            require(stakingRewards != address(0), 'StakingRewardsFactory::notifyRewardAmount: not deployed');
            uint256 tokenAmount = IERC20(tokenAddress).balanceOf(stakingRewards);
            if (tokenAmount > 0) {
                StakingRewards(stakingRewards).recoverERC20(tokenAddress, tokenAmount);
                IERC20(tokenAddress).transfer(owner, tokenAmount);
            }
        }
    }

    function setRewardsDuration(uint256 _rewardsDuration) public auth {
        for (uint i = 0; i < stakingTokens.length; i++) {
            address stakingRewards = stakingRewardsInfoByStakingToken[stakingTokens[i]];
            require(stakingRewards != address(0), 'StakingRewardsFactory::notifyRewardAmount: not deployed');
            StakingRewards(stakingRewards).setRewardsDuration(_rewardsDuration);
        }
    }

    // call notifyRewardAmount for all staking tokens.
    function notifyRewardAmounts(uint256 rewardAmount) public auth {
        require(stakingTokens.length > 0, 'StakingRewardsFactory::notifyRewardAmounts: called before any deploys');
        uint256 rewardEachAmount = rewardAmount / stakingTokens.length;
        for (uint i = 0; i < stakingTokens.length; i++) {
            notifyRewardAmount(stakingTokens[i], rewardEachAmount);
        }
    }

    // notify reward amount for an individual staking token.
    // this is a fallback in case the notifyRewardAmounts costs too much gas to call for all contracts
    function notifyRewardAmount(address stakingToken, uint256 rewardAmount) public auth {
        require(block.timestamp >= stakingRewardsGenesis, 'StakingRewardsFactory::notifyRewardAmount: not ready');

        address stakingRewards = stakingRewardsInfoByStakingToken[stakingToken];
        require(stakingRewards != address(0), 'StakingRewardsFactory::notifyRewardAmount: not deployed');
        require(rewardAmount > 0, 'StakingRewardsFactory::notifyRewardAmount: reward is zero');
        require(
            IERC20(rewardsToken).transfer(stakingRewards, rewardAmount),
            'StakingRewardsFactory::notifyRewardAmount: transfer failed'
        );
        StakingRewards(stakingRewards).notifyRewardAmount(rewardAmount);
    }

    function recover(address tokenAddress) public auth {
        uint256 tokenAmount = IERC20(tokenAddress).balanceOf(address(this));
        IERC20(tokenAddress).transfer(owner, tokenAmount);
    }
}
