// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AdEscrow} from "../src/Escrow.sol";

address constant MOCK_FORWARDER = 0x15fC6ae953E024d975e77382eEeC56A9101f9F88;

contract DeployEscrow is Script {
    function run() external returns (AdEscrow) {
        // Read environment variables
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

        // Optional: Read authorized withdrawer from env, default to mock Keystone forwarder
        address forwarder = vm.envOr("AUTHORIZED_WITHDRAWER", MOCK_FORWARDER);

        // Optional: Read whitelisted tokens from env (comma-separated addresses)
        // If not provided, only native ETH will be whitelisted
        address[] memory whitelistedTokens = _parseWhitelistedTokens();

        // Start broadcasting transactions
        vm.startBroadcast(deployerPrivateKey);

        // Deploy the contract
        AdEscrow escrow = new AdEscrow(forwarder, whitelistedTokens);

        console.log("AdEscrow deployed at:", address(escrow));
        console.log("Forwarder:", forwarder);
        console.log("Deployer:", vm.addr(deployerPrivateKey));
        console.log("Whitelisted tokens count:", whitelistedTokens.length);
        for (uint256 i = 0; i < whitelistedTokens.length; i++) {
            console.log("  Token", i, ":", whitelistedTokens[i]);
        }

        vm.stopBroadcast();

        return escrow;
    }

    function _parseWhitelistedTokens() internal view returns (address[] memory) {
        // Try to read from env variable (comma-separated addresses)
        // Format: "0x123...,0x456...,0x789..."
        try vm.envString("WHITELISTED_TOKENS") returns (string memory tokensStr) {
            if (bytes(tokensStr).length == 0) {
                return new address[](0);
            }
            
            // Count commas to determine array size
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
            // If env variable doesn't exist, return empty array (only native ETH whitelisted)
            return new address[](0);
        }
    }
}
