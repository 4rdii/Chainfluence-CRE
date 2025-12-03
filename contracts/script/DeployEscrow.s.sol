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

        // Start broadcasting transactions
        vm.startBroadcast(deployerPrivateKey);

        // Deploy the contract
        AdEscrow escrow = new AdEscrow(forwarder);

        console.log("AdEscrow deployed at:", address(escrow));
        console.log("Forwarder:", forwarder);
        console.log("Deployer:", vm.addr(deployerPrivateKey));

        vm.stopBroadcast();

        return escrow;
    }
}
