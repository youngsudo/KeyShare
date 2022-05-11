// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./user.sol";
import "./account.sol";

contract Key is AccountClass, User  {
    constructor(uint8 _adminNum) User(_adminNum) {
    }
}