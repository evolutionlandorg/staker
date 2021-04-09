pragma solidity ^0.6.7;

import "ds-test/test.sol";

import "./Staker.sol";

contract StakerTest is DSTest {
    Staker staker;

    function setUp() public {
        staker = new Staker();
    }

    function testFail_basic_sanity() public {
        assertTrue(false);
    }

    function test_basic_sanity() public {
        assertTrue(true);
    }
}
