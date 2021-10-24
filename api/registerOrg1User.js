/*
 * SPDX-License-Identifier: Apache-2.0
 */

"use strict";

const FabricCAServices = require('fabric-ca-client');
const { Wallets } = require('fabric-network');
const path = require("path");
const fs = require("fs");

const ccpPath = path.resolve(__dirname, "connection-org1.json");
const ccpJSON = fs.readFileSync(ccpPath, "utf8");
const ccp = JSON.parse(ccpJSON);
async function main() {
  try {
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet");
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.get("org1user");
    if (userExists) {
      console.log(
        'An identity for the user "org1user" already exists in the wallet'
      );
      return;
    }

    // Check to see if we've already enrolled the admin user.
    const adminExists = await wallet.get("admin-org1");
    if (!adminExists) {
      const caInfo = ccp.certificateAuthorities['ca.org1.example.com'];
      const caTLSCACerts = caInfo.tlsCACerts.pem;
      const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

      // Enroll the admin user, and import the new identity into the wallet.
      const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
      const x509Identity = {
        credentials: {
          certificate: enrollment.certificate,
          privateKey: enrollment.key.toBytes(),
        },
        mspId: 'Org1MSP',
        type: 'X.509',
      };
      await wallet.put('admin-org1', x509Identity);
    }

    const adminIdentity = await wallet.get('admin-org1');


    // Get the CA client object from the gateway for interacting with the CA.
    const caURL = ccp.certificateAuthorities['ca.org1.example.com'].url;
    const ca = new FabricCAServices(caURL);

    const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, 'admin-org1');

    const secret = await ca.register({
      affiliation: 'org1.department1',
      enrollmentID: 'org1user',
      enrollmentSecret: 'org1userpw',
      role: 'client'
    }, adminUser);
    const enrollment = await ca.enroll({
      enrollmentID: 'org1user',
      enrollmentSecret: secret
    });
    const x509Identity = {
      credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes(),
      },
      mspId: 'Org1MSP',
      type: 'X.509',
    };
    await wallet.put('org1user', x509Identity);
    console.log('Successfully registered and enrolled admin user "org1user" and imported it into the wallet');

  } catch (error) {
    console.error(`Failed to register user "org1user": ${error}`);
    process.exit(1);
  }
}

main();
