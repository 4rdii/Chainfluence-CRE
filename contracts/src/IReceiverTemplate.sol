// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {IERC165} from "@openzeppelin/contracts/interfaces/IERC165.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IReceiver} from "./IReceiver.sol";

/**
 * @title IReceiverTemplate
 * @notice Abstract receiver with optional permission controls for CRE workflows
 * @dev Permission fields default to zero (disabled). Enable checks via the provided setters.
 */
abstract contract IReceiverTemplate is IReceiver, Ownable {
    // Optional permission fields (zero value disables the check)
    address public forwarderAddress; // Restricts sender
    address public expectedAuthor; // Restricts workflow owner
    bytes10 public expectedWorkflowName; // Restricts workflow name
    bytes32 public expectedWorkflowId; // Restricts workflow ID

    // Custom errors
    error InvalidSender(address sender, address expected);
    error InvalidAuthor(address received, address expected);
    error InvalidWorkflowName(bytes10 received, bytes10 expected);
    error InvalidWorkflowId(bytes32 received, bytes32 expected);

    constructor() Ownable(msg.sender) {}

    /**
     * @inheritdoc IReceiver
     */
    function onReport(bytes calldata metadata, bytes calldata report) external override {
        if (forwarderAddress != address(0) && msg.sender != forwarderAddress) {
            revert InvalidSender(msg.sender, forwarderAddress);
        }

        if (expectedWorkflowId != bytes32(0) || expectedAuthor != address(0) || expectedWorkflowName != bytes10(0)) {
            (bytes32 workflowId, bytes10 workflowName, address workflowOwner) = _decodeMetadata(metadata);

            if (expectedWorkflowId != bytes32(0) && workflowId != expectedWorkflowId) {
                revert InvalidWorkflowId(workflowId, expectedWorkflowId);
            }
            if (expectedAuthor != address(0) && workflowOwner != expectedAuthor) {
                revert InvalidAuthor(workflowOwner, expectedAuthor);
            }
            if (expectedWorkflowName != bytes10(0) && workflowName != expectedWorkflowName) {
                revert InvalidWorkflowName(workflowName, expectedWorkflowName);
            }
        }

        _processReport(report);
    }

    /**
     * @notice Updates the allowed forwarder address
     * @param _forwarder Address of the Chainlink forwarder (set to zero to disable)
     */
    function setForwarderAddress(address _forwarder) external onlyOwner {
        forwarderAddress = _forwarder;
    }

    /**
     * @notice Updates the expected workflow owner
     * @param _author Address of the workflow owner (set to zero to disable)
     */
    function setExpectedAuthor(address _author) external onlyOwner {
        expectedAuthor = _author;
    }

    /**
     * @notice Updates the expected workflow name
     * @param _name Workflow registry name (empty string disables the check)
     */
    function setExpectedWorkflowName(string calldata _name) external onlyOwner {
        if (bytes(_name).length == 0) {
            expectedWorkflowName = bytes10(0);
            return;
        }

        bytes32 hash = sha256(bytes(_name));
        bytes memory hexString = _bytesToHexString(abi.encodePacked(hash));
        bytes memory first10 = new bytes(10);
        for (uint256 i = 0; i < 10; i++) {
            first10[i] = hexString[i];
        }
        // Casting to bytes10 is safe because we explicitly build a 10-byte array.
        // forge-lint: disable-next-line(unsafe-typecast)
        expectedWorkflowName = bytes10(first10);
    }

    /**
     * @notice Updates the expected workflow ID
     * @param _id Workflow ID (set to zero to disable)
     */
    function setExpectedWorkflowId(bytes32 _id) external onlyOwner {
        expectedWorkflowId = _id;
    }

    function _bytesToHexString(bytes memory data) private pure returns (bytes memory) {
        bytes memory hexChars = "0123456789abcdef";
        bytes memory hexString = new bytes(data.length * 2);

        for (uint256 i = 0; i < data.length; i++) {
            hexString[i * 2] = hexChars[uint8(data[i] >> 4)];
            hexString[i * 2 + 1] = hexChars[uint8(data[i] & 0x0f)];
        }

        return hexString;
    }

    function _decodeMetadata(bytes memory metadata)
        internal
        pure
        returns (bytes32 workflowId, bytes10 workflowName, address workflowOwner)
    {
        assembly {
            workflowId := mload(add(metadata, 32))
            workflowName := mload(add(metadata, 64))
            workflowOwner := shr(96, mload(add(metadata, 74)))
        }
    }

    /**
     * @notice Implement business logic for decoded report data
     * @param report Encoded report payload
     */
    function _processReport(bytes calldata report) internal virtual;

    /**
     * @inheritdoc IERC165
     */
    function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
        return interfaceId == type(IReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
    }
}
