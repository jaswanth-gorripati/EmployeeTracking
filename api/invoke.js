/*
 * SPDX-License-Identifier: Apache-2.0
 */

"use strict";

const { Gateway, Wallets } = require('fabric-network');
const path = require("path");
const fs = require("fs")

const ccpPath = path.resolve(__dirname, "connection-org1.json");
let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

async function invoke(fcn, args) {
  try {
    // Create a new file system based wallet for managing identities.
    console.log(fcn, args);
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.get("org1user");
    if (!userExists) {
      console.log(
        'An identity for the user "org1user" does not exist in the wallet'
      );
      console.log("Run the registerUser.js application before retrying");
      return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
      wallet,
      identity: "org1user",
      discovery: { enabled: true, asLocalhost: true }
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork("trackingchannel");

    // Get the contract from the network.
    const contract = network.getContract("emptrackcc");

    // Submit the specified transaction.
    await contract.submitTransaction(fcn, args);
    console.log("Transaction has been submitted");

    // Disconnect from the gateway.
    await gateway.disconnect();
    return "Transaction has been submitted Successfully";
  } catch (error) {
    console.error(`Failed to submit transaction: ${error}`);
    throw error;
  }
}

exports.invoke = invoke;
