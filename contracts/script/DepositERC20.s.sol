// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrow} from "../src/Escrow.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract DepositERC20 is Script {
    address constant ESCROW_ADDRESS = 0x502c3323D62425316bE18C09B751c5Ce09FCC943;
    address constant USER = 0xFF3f9d0E76D30577D7A0bD35E14259a76078230e;
    address constant TOKEN_ADDRESS = 0x12394F560E5645E18204e67c2C7c4AaED0b1dC4B;
    
    // 1000 tokens with 18 decimals
    uint256 constant AMOUNT = 1000 * 10**18;
    
    // 30 days in seconds
    uint64 constant CAMPAIGN_DURATION = 60 days;

    function run() external {
        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        address depositor = vm.addr(privateKey);
        
        uint256 expiryDeadline = block.timestamp + 1 days;
        uint256 minViews = 6;
        string memory contentText = "Hi This is the second test";

        vm.startBroadcast(privateKey);

        AdEscrow escrow = AdEscrow(ESCROW_ADDRESS);
        IERC20 token = IERC20(TOKEN_ADDRESS);
        
        // Check if token is whitelisted
        bool isWhitelisted = escrow.whitelistedTokens(TOKEN_ADDRESS);
        require(isWhitelisted, "Token not whitelisted");
        
        // Check token balance
        uint256 balance = token.balanceOf(depositor);
        console.log("Depositor balance:", balance);
        require(balance >= AMOUNT, "Insufficient token balance");
        
        // Approve escrow to spend tokens
        console.log("Approving escrow to spend tokens...");
        token.approve(ESCROW_ADDRESS, AMOUNT);
        
        // Deposit tokens to escrow
        console.log("Depositing tokens to escrow...");
        uint256 campaignId = escrow.deposit(
            TOKEN_ADDRESS,
            USER,
            AMOUNT,
            contentText,
            minViews,
            expiryDeadline,
            CAMPAIGN_DURATION
        );

        console.log("Deposited for campaign", campaignId);
        console.log("Token:", TOKEN_ADDRESS);
        console.log("Amount:", AMOUNT);
        console.log("Campaign Duration:", CAMPAIGN_DURATION, "seconds (30 days)");
        console.log("Next campaign ID will be:", escrow.campaignCounter());

        vm.stopBroadcast();
    }
}
