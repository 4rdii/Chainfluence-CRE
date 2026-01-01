#!/usr/bin/env node

/**
 * Script to trigger CRE workflows via HTTP with JWT authentication
 * 
 * Usage:
 *   npm run trigger
 * 
 * Environment variables required:
 *   - WORKFLOW_ID: Your 64-character workflow ID
 *   - PRIVATE_KEY: Private key (hex, with or without 0x prefix)
 *   - CAMPAIGN_ID: Campaign ID to check
 *   - TWEET_URL: (optional) Tweet URL to verify
 *   - GATEWAY_URL: (optional) Defaults to https://01.gateway.zone-a.cre.chain.link
 */

import { createJWT } from "./src/create-jwt.js";
import * as dotenv from "dotenv";
import { randomUUID } from "crypto";
import stringify from "json-stable-stringify";

// Load environment variables
dotenv.config();

interface TriggerInput {
  campaignId: string;
  tweetUrl?: string;
}

async function triggerWorkflow() {
  // Read environment variables
  const workflowID = process.env.WORKFLOW_ID;
  const privateKeyHex = process.env.PRIVATE_KEY;
  const gatewayURL = process.env.GATEWAY_URL || "https://01.gateway.zone-a.cre.chain.link";
  const campaignID = process.env.CAMPAIGN_ID;
  const tweetURL = process.env.TWEET_URL;

  // Validate required variables
  if (!workflowID) {
    console.error("Error: WORKFLOW_ID environment variable is required");
    process.exit(1);
  }
  if (!privateKeyHex) {
    console.error("Error: PRIVATE_KEY environment variable is required");
    process.exit(1);
  }
  if (!campaignID) {
    console.error("Error: CAMPAIGN_ID environment variable is required");
    process.exit(1);
  }

  // Ensure private key has 0x prefix (required by viem)
  const privateKey = privateKeyHex.startsWith("0x") 
    ? privateKeyHex 
    : `0x${privateKeyHex}`;

  // Prepare input payload
  const input: TriggerInput = {
    campaignId: campaignID,
  };
  if (tweetURL) {
    input.tweetUrl = tweetURL;
  }

  // Create JSON-RPC request
  const request = {
    jsonrpc: "2.0" as const,
    id: randomUUID(),
    method: "workflows.execute",
    params: {
      input,
      workflow: {
        workflowID,
      },
    },
  };

  console.log("🚀 Triggering CRE workflow...");
  console.log(`Workflow ID: ${workflowID}`);
  console.log(`Campaign ID: ${campaignID}`);
  if (tweetURL) {
    console.log(`Tweet URL: ${tweetURL}`);
  }
  console.log();

  try {
    // Create JWT token
    const jwt = await createJWT(request, privateKey as `0x${string}`);

    // Send HTTP request
    const response = await fetch(gatewayURL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${jwt}`,
      },
      body: stringify(request),
    });

    const responseData = await response.json();

    if (!response.ok || responseData.error) {
      console.error("❌ Workflow trigger failed!");
      if (responseData.error) {
        console.error("Error:", JSON.stringify(responseData.error, null, 2));
      } else {
        console.error("HTTP Status:", response.status);
        console.error("Response:", JSON.stringify(responseData, null, 2));
      }
      process.exit(1);
    }

    // Success!
    console.log("✅ Workflow triggered successfully!");
    console.log("Response:", JSON.stringify(responseData, null, 2));

    if (responseData.result) {
      const result = responseData.result;
      if (result.workflow_execution_id) {
        console.log(`\n📊 Execution ID: ${result.workflow_execution_id}`);
        console.log(`View execution: https://cre.chain.link/workflows`);
      }
    }
  } catch (error) {
    console.error("❌ Error triggering workflow:", error);
    if (error instanceof Error) {
      console.error("Error message:", error.message);
    }
    process.exit(1);
  }
}

// Run the script
triggerWorkflow().catch((error) => {
  console.error("Fatal error:", error);
  process.exit(1);
});

