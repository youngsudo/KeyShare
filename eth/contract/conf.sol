//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

enum UserType {
    administrator, // 超级管理员地址 0
    admin,  // 管理员 1
    normal  // 普通用户 2
}

enum KeyType {
   privateKey,
   publicKey
}