// hevm: flattened sources of src/RewardsDistributionRecipient.sol

pragma solidity >=0.5.16 <0.6.0;

////// src/RewardsDistributionRecipient.sol
/* pragma solidity ^0.5.16; */

contract RewardsDistributionRecipient {
    address public rewardsDistribution;
    function notifyRewardAmount(uint256 reward) external;
    modifier onlyRewardsDistribution() {
        require(msg.sender == rewardsDistribution, "Caller is not RewardsDistribution contract");
        _;
    }
}

