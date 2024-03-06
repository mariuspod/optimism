// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { SafeCall } from "src/libraries/SafeCall.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { L1Block } from "src/L2/L1Block.sol";
import { ISemver } from "src/universal/ISemver.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000022
/// @title CrossL2Inbox
/// @notice The CrossL2Inbox is responsible for executing a cross chain message on the destination
///         chain. It is permissionless to execute a cross chain message on behalf of any user.
contract CrossL2Inbox is ISemver {
    struct Identifier {
        address origin;
        uint256 blocknumber;
        uint256 logIndex;
        uint256 timestamp;
        uint256 chainId;
    }

    address public l1Block;
}
