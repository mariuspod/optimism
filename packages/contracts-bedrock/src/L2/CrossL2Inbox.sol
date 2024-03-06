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

    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    function origin() public view returns (address _origin) { }

    function blocknumber() public view returns (uint256 _blocknumber) { }

    function logIndex() public view returns (uint256 _logIndex) { }

    function timestamp() public view returns (uint256 _timestamp) { }

    function chainId() public view returns (uint256 _chainId) { }

    /// @notice Executes a cross chain message on the destination chain
    /// @param _msg The message payload, matching the initiating message.
    /// @param _id A Identifier pointing to the initiating message.
    /// @param _target Account that is called with _msg.
    function executeMessage(Identifier calldata _id, address _target, bytes calldata _msg) public payable {
        require(_id.timestamp <= block.timestamp, "CrossL2Inbox: invalid id timestamp"); // timestamp invariant
        uint256 chainId_ = _id.chainId;
        require(msg.sender == tx.origin, "CrossL2Inbox: Not EOA sender"); // only EOA invariant
    }
}
