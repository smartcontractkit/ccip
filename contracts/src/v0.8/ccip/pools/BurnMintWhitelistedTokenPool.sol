// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintWhitelisted} from "../interfaces/IBurnMintWhitelisted.sol";

import {TokenPool} from "./TokenPool.sol";
import {BurnMintTokenPoolAbstract} from "./BurnMintTokenPoolAbstract.sol";

contract BurnMintWhitelistedTokenPool is BurnMintTokenPoolAbstract, ITypeAndVersion {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "BurnMintTokenPool 1.2.0";

  constructor(
    IBurnMintWhitelisted token,
    address[] memory allowlist,
    address armProxy
  ) TokenPool(token, allowlist, armProxy) {}

  /// @inheritdoc BurnMintTokenPoolAbstract
  function _mint(address receiver, uint256 amount) internal virtual override {
    IBurnMintWhitelisted token = IBurnMintWhitelisted(address(i_token));
    // NOTE: the mintFromWhitelistedContract call mints tokens to the caller,
    // so we need an additional transfer call to move funds to the receiver address.
    token.mintFromWhitelistedContract(amount);
    token.transfer(receiver, amount);
  }

  /// @inheritdoc BurnMintTokenPoolAbstract
  function _burn(uint256 amount) internal virtual override {
    IBurnMintWhitelisted(address(i_token)).burnFromWhitelistedContract(amount);
  }
}
