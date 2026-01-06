// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrow} from "../src/Escrow.sol";

contract DepositEscrow is Script {
    address constant ESCROW_ADDRESS = 0x502c3323D62425316bE18C09B751c5Ce09FCC943;
    address constant USER = 0xFF3f9d0E76D30577D7A0bD35E14259a76078230e;

    function run() external {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        uint256 expiryDeadline = block.timestamp + 1 days;
        uint256 minViews = 6;
        uint64 campaignDuration = 0; // 0 means duration check is disabled
        string memory contentText = "Hi This is the second test";
        address token = address(0); // Native ETH (NATIVE_ETH_ADDRESS)

        vm.startBroadcast(privateKey);

        AdEscrow escrow = AdEscrow(ESCROW_ADDRESS);
        uint256 campaignId = escrow.deposit{value: 0.01 ether}(
            token,
            USER,
            0.01 ether,
            contentText,
            minViews,
            expiryDeadline,
            campaignDuration
        );

        console.log("Deposited for campaign", campaignId);
        console.log("Next campaign ID will be:", escrow.campaignCounter());

        vm.stopBroadcast();
    }
}
