// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrow} from "../src/Escrow.sol";

contract DepositEscrow is Script {
    address constant ESCROW_ADDRESS = 0x1e34128a274115D63Fe7a6d6A8A7342531076cc8;
    address constant USER = 0xFF3f9d0E76D30577D7A0bD35E14259a76078230e;

    function run() external {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        uint256 deadline = block.timestamp + 1 days;
        uint256 campaignId = 2;
        uint256 minViews = 6;
        string memory contentText = "Hi This is the second test";

        vm.startBroadcast(privateKey);

        AdEscrow escrow = AdEscrow(ESCROW_ADDRESS);
        escrow.deposit{value: 0.01 ether}(campaignId, USER, contentText, minViews, deadline);

        console.log("Deposited for campaign", campaignId);

        vm.stopBroadcast();
    }
}
