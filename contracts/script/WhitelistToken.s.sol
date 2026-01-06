// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrow} from "../src/Escrow.sol";

contract WhitelistToken is Script {
    address constant ESCROW_ADDRESS = 0x502c3323D62425316bE18C09B751c5Ce09FCC943;
    address constant TOKEN_ADDRESS = 0x12394F560E5645E18204e67c2C7c4AaED0b1dC4B;

    function run() external {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(privateKey);

        AdEscrow escrow = AdEscrow(ESCROW_ADDRESS);
        
        // Check if token is already whitelisted
        bool isWhitelisted = escrow.whitelistedTokens(TOKEN_ADDRESS);
        if (isWhitelisted) {
            console.log("Token already whitelisted:", TOKEN_ADDRESS);
        } else {
            escrow.whitelistToken(TOKEN_ADDRESS);
            console.log("Token whitelisted:", TOKEN_ADDRESS);
        }

        vm.stopBroadcast();
    }
}
