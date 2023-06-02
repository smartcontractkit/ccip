// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../shared/token/ERC20/IBurnMintERC20.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {BurnMintTokenPool} from "./BurnMintTokenPool.sol";
import {Router} from "../Router.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableSet.sol";

/// @notice This pool mints and burns a 3rd-party token. This pool is not owned by the DON
// and therefor has an additional check on adding offRamps.
contract ThirdPartyBurnMintTokenPool is BurnMintTokenPool {
  using EnumerableSet for EnumerableSet.AddressSet;

  error InvalidOffRamp(address offRamp);

  /// @notice the trusted Router address to validate new offRamps through.
  address private s_router;

  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig,
    address router
  ) BurnMintTokenPool(token, allowlist, rateLimiterConfig) {
    s_router = router;
  }

  /// @notice Sets permissions for all on and offRamps.
  /// @dev Only callable by the owner
  /// @param onRamps A list of onRamps and their new permission status
  /// @param offRamps A list of offRamps and their new permission status
  function applyRampUpdates(RampUpdate[] memory onRamps, RampUpdate[] memory offRamps) public override onlyOwner {
    for (uint256 i = 0; i < onRamps.length; ++i) {
      RampUpdate memory update = onRamps[i];

      // No need to check the onRamps as they can never extract value from a pool, only add to it.
      if (update.allowed ? s_onRamps.add(update.ramp) : s_onRamps.remove(update.ramp)) {
        emit OnRampAllowanceSet(onRamps[i].ramp, onRamps[i].allowed);
      }
    }

    for (uint256 i = 0; i < offRamps.length; ++i) {
      RampUpdate memory update = offRamps[i];
      if (!update.allowed) {
        if (s_offRamps.remove(update.ramp)) {
          emit OffRampAllowanceSet(offRamps[i].ramp, offRamps[i].allowed);
        }
        continue;
      }
      // If the offRamp is being added do an additional check if the offRamp is
      // permission by the router. If not, we revert because we tried to add an
      // invalid offRamp.
      (bool exists, ) = Router(s_router).isOffRamp(update.ramp);
      if (exists) {
        if (s_offRamps.add(update.ramp)) {
          emit OffRampAllowanceSet(offRamps[i].ramp, offRamps[i].allowed);
        }
        continue;
      }
      revert InvalidOffRamp(update.ramp);
    }
  }
}
