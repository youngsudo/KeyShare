// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Start {

// 定义一个用户的结构体
struct UserStruct  {
    address addr; // 用户地址
    string username; // 用户名
    string password; // 密码
    string email; // 邮箱
    uint time;  // 注册时间
    bool exist; // 用户是否存在
}

// 存储已经存在的全部用户信息   information
address[] public userAddresses; // 所有地址集合 // 判断用户地址是否存在
string[] public usernames; // 所有用户名集合   // 判断用户名是否存在
string[] public emails; // 所有邮箱集合   // 判断邮箱是否已存在


//  mapping映射将地址映射到用户数据结构，这个可以初略理解为一个地址所对应的值有哪些
mapping(address => UserStruct) public userStruct; // 存储账户个人信息
// userStruct[address].exist = true;    // 判断用户存在

// 创建一个新用户 只能由管理员身份创建
function addUser(address _addr,string memory _username, string memory _password, string memory _email) public {
    // 创建一个新的用户
    userStruct[_addr] =  UserStruct({
        addr: _addr,
        username: _username,
        password: _password,
        email: _email,
        time: block.timestamp,
        exist: true
    });
   
    // 将用户信息添加到 information 中
    userAddresses.push(_addr);
    usernames.push(_username);
    emails.push(_email);

}
}