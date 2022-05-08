// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Key {

address public immutable admin; // 管理员地址

constructor()  {
    admin = msg.sender;
}
// 定义一个用户的结构体
struct UserStruct  {
    address addr; // 用户地址 唯一的
    string account; // 用户账号 唯一的
    string password; // 密码
    string email; // 邮箱 唯一的
    uint256 time;  // 注册时间
    bool exist; // 用户是否存在
}
 //定义用户列表数据结构
struct UserList {   // 通过 Sum 获取全部用户的地址
    uint256 index; // 用户地址的索引
    address userAddr;
}

uint256 public Sum = 0; // 所有地址的总数

//  mapping映射将地址映射到用户数据结构，这个可以初略理解为一个地址所对应的值有哪些
mapping(address => UserStruct) public userStruct; // 存储账户个人信息
// userStruct[address].exist = true;    // 判断用户存在
mapping(string=> address) public accountMap; // 账号与地址的关联
mapping(string => address) public emailMap; // 账号与地址的关联
mapping(uint256 => UserList) public userList; // 通过 []数组下标映射地址来存储所有用户的地址

// 创建一个新用户 只能由管理员身份创建
function addUser(address _addr,string memory _account, string memory _password, string memory _email) public {
   require(msg.sender == admin,"You're not an administrator!");    // 判断是否是管理员
   require(isExitUserAddress(_addr) == false, "User already exists! please use another address");
   require(isExitUserAccount(_account) == false, "Account already exists! please use another account");
   require(isExitUserEmail(_email) == false, "Email already exists! please use another email");
    // 创建一个新的用户
    userStruct[_addr] =  UserStruct({
        addr: _addr,
        account: _account,
        password: _password,
        email: _email,
        time: block.timestamp,
        exist: true
    });
   
   // 将地址与账号和邮箱关联
    accountMap[_account] = _addr;
    emailMap[_email] = _addr;

    // 将地址添加到数组中
    userList[Sum] = UserList({
        index: Sum,     // 考虑使用 num 来代替 index,这样子代表的是元素是第几个,为 0 时代表没有元素
        userAddr: _addr
    });
    Sum++;  // 地址总数加1
}

//判断用户地址是否存在
function isExitUserAddress(address _userAddress) public view returns(bool) {
       return userStruct[_userAddress].exist;    
}

//判断 用户账号 是否存在
function isExitUserAccount(string memory _userAddress) public view returns(bool) {
    // accountMap => account => userStruct{} => exist
    return userStruct[accountMap[_userAddress]].exist;
    
}
//判断 用户邮箱 是否存在
function isExitUserEmail(string memory _email) public view returns(bool) {
       return userStruct[emailMap[_email]].exist;
}

}