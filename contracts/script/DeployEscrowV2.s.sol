// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrowV2} from "../src/AdEscrowV2.sol";

contract DeployEscrowV2 is Script {
    function run() external returns (AdEscrowV2) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        // CRE Keystone forwarder
        address forwarder = vm.envOr("AUTHORIZED_WITHDRAWER", address(0));

        // Platform fee in basis points (default 250 = 2.5%)
        uint256 feeBps = vm.envOr("PLATFORM_FEE_BPS", uint256(250));

        // Fee recipient (defaults to deployer)
        address feeRecipient = vm.envOr("FEE_RECIPIENT", deployer);

        // Whitelisted ERC20 tokens (native ETH is always whitelisted in constructor)
        address[] memory whitelistedTokens = _parseWhitelistedTokens();

        vm.startBroadcast(deployerPrivateKey);

        AdEscrowV2 escrow = new AdEscrowV2(
            forwarder,
            whitelistedTokens,
            feeBps,
            feeRecipient
        );

        console.log("AdEscrowV2 deployed at:", address(escrow));
        console.log("Deployer / Owner:", deployer);
        console.log("Forwarder:", forwarder);
        console.log("Platform fee:", feeBps, "bps");
        console.log("Fee recipient:", feeRecipient);
        console.log("Whitelisted tokens:", whitelistedTokens.length);
        for (uint256 i = 0; i < whitelistedTokens.length; i++) {
            console.log("  Token", i, ":", whitelistedTokens[i]);
        }

        vm.stopBroadcast();

        return escrow;
    }

    function _parseWhitelistedTokens() internal view returns (address[] memory) {
        try vm.envString("WHITELISTED_TOKENS") returns (string memory tokensStr) {
            if (bytes(tokensStr).length == 0) {
                return new address[](0);
            }

            uint256 count = 1;
            for (uint256 i = 0; i < bytes(tokensStr).length; i++) {
                if (bytes(tokensStr)[i] == ",") {
                    count++;
                }
            }

            address[] memory tokens = new address[](count);
            bytes memory strBytes = bytes(tokensStr);
            uint256 currentIndex = 0;
            uint256 startIndex = 0;

            for (uint256 i = 0; i <= strBytes.length; i++) {
                if (i == strBytes.length || strBytes[i] == ",") {
                    if (i > startIndex) {
                        bytes memory addrBytes = new bytes(i - startIndex);
                        for (uint256 j = startIndex; j < i; j++) {
                            addrBytes[j - startIndex] = strBytes[j];
                        }
                        tokens[currentIndex] = vm.parseAddress(string(addrBytes));
                        currentIndex++;
                    }
                    startIndex = i + 1;
                }
            }

            return tokens;
        } catch {
            return new address[](0);
        }
    }
}
