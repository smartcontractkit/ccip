// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";

// setup
contract Any2EVMTollSenderDappSetup is TokenSetup {
  function setUp() public virtual override {
    TokenSetup.setUp();
  }
}

/// @notice #constructor
contract Any2EVMTollSenderDapp_constructor is Any2EVMTollSenderDappSetup {
  // Success
  function testSuccess() public {}
  // it('should set the onRamp', async () => {
  //   it('#should set the destination contract', async () => {
  //   it('should send a request to the onRamp', async () => {
  //   it('should fail if the destination address is zero address', async () => {
  //   it('should return the correct type and version', async () => {
}
