// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./user.sol";

contract Key is User  {
    constructor(uint8 _adminNum) User(_adminNum) {
    }
}