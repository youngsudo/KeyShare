// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./conf.sol";

contract User{
    UserType usertype;   // 引入UserType枚举类型
    address public immutable administrator; // 超级管理员地址
    uint256 Sum; // 所有地址的总数
    address[] adminList; // 管理员地址列表  动态数组
    uint8 public immutable adminNum; // 管理员数量
    bool AddMyself = false; // 是否已经添加自己到管理员列表

// 定义一个用户的结构体
struct UserStruct  {
    address addr; // 用户地址 唯一的
    string account; // 用户账号 唯一的
    string password; // 密码
    string email; // 邮箱 唯一的
    uint256 time;  // 注册时间
    uint8 usertype; // 用户类型
    bool exist; // 用户是否存在
}
 //定义用户列表数据结构
struct UserList {   // 通过 Sum 获取全部用户的地址
    uint256 index; // 用户地址的索引
    address userAddr;
}

struct AdminStruct{
    uint8 index;    // 管理员地址在列表中的索引
}
//  mapping映射将地址映射到用户数据结构，这个可以初略理解为一个地址所对应的值有哪些
mapping(address => UserStruct) private userStruct; // 存储账户个人信息
// userStruct[address].exist = true;    // 判断用户存在
mapping(string => address) private accountMap; // 账号与地址的关联
mapping(string => address) private emailMap; // 账号与地址的关联
mapping(uint256 => UserList) private userList; // 通过 []数组下标映射地址来存储所有用户的地址
mapping(address => AdminStruct) private adminIndex; // 存储管理员在 管理员列表中的 索引(下标)
mapping(address => string) private symbol;  // 存储助记符
mapping(string => address) symbolToAddr; // 存储助记符与地址的关联

constructor(uint8 _adminNum)  {
    require(_adminNum >= 1,"At least one administrator");
    administrator = msg.sender;  // 超级管理员地址
    adminNum = _adminNum;   // 管理员总数量,包括超级管理员
}

modifier onlyAdministrator {
 require(msg.sender == administrator ,"You're not an administrator!");    // 判断是否是超级管理员
 _;
}

modifier onlyAdmin {
    require(AddMyself, "Administrator,you must add yourself first!");
    require(isExitUserAddressFunc(msg.sender), "The current admin user address does not exist!"); 
    require(isAdminFunc(msg.sender),"You're not an admin!");    // 判断是否是管理员
 _;
}

// 创建一个新用户
function addUserFunc(address _addr,string memory _account, string memory _password, string memory _email,string memory _symbol) public {
   require(!isExitUserAddressFunc(_addr), "User already exists! please use another address");
   require(!isExitUserAccountFunc(_account), "Account already exists! please use another account");
   require(!isExitUserEmailFunc(_email), "Email already exists! please use another email");
    
    // 将地址与账号和邮箱以及助记符关联
    accountMap[_account] = _addr;
    emailMap[_email] = _addr;
    symbol[_addr] = _symbol;
    Sum++;  // 所有地址总数加1
    // 将地址添加到数组中   索引为 0 时代表的是超级管理员身份,或者没有元素 
    if(_addr == administrator){
            // 创建一个超级管理员用户
            userStruct[_addr] =  UserStruct({
                addr: _addr,
                account: _account,
                password: _password,
                email: _email,
                time: block.timestamp,
                usertype: 0,    // 超级管理员
                exist: true
            });
            userList[0] = UserList({
                index: 0,    
                userAddr: _addr
            });
    }else{
        // 创建一个新的用户
        userStruct[_addr] =  UserStruct({
            addr: _addr,
            account: _account,
            password: _password,
            email: _email,
            time: block.timestamp,
            usertype: 2,    // 普通用户
            exist: true
        });
            userList[Sum] = UserList({
            index: Sum,  
            userAddr: _addr
    });
    }
}

// 判断用户是否为管理员 0 是administrator 超级管理员; 1 admin 管理员; 2 是普通用户
function isAdminFunc(address _addr) public view returns (bool) {
    require(isExitUserAddressFunc(_addr), "User does not exist!"); // 判断用户是否存在
    return userStruct[_addr].usertype <= 1; // 判断用户是否为管理员
}

// 判断用户地址是否存在(是否已经注册)
function isExitUserAddressFunc(address _userAddress) public view returns(bool) {
       return userStruct[_userAddress].exist;    
}

// 判断 用户账号 是否存在
function isExitUserAccountFunc(string memory _userAddress) public view returns(bool) {
    // accountMap => account => userStruct{} => exist
    return userStruct[accountMap[_userAddress]].exist;
    
}
// 判断 用户邮箱 是否存在
function isExitUserEmailFunc(string memory _email) public view returns(bool) {
       return userStruct[emailMap[_email]].exist;
}

// 管理员查看用户的总数量
function getSumFunc() public view onlyAdmin returns(uint256) {
    return Sum;
}

// 超管必须先将自己先升级为管理员  
function addMyselfFunc() public onlyAdministrator{
    require(!AddMyself, "You already add yourself!");
    adminList.push(msg.sender) ;
    adminIndex[msg.sender].index = 0; // 超级管理员索引为0
    AddMyself = true;   // 已经添加自己
}
// 查看所有管理员的地址 onlyAdmin
function getAdminListFunc() public view onlyAdmin returns(address[] memory) {
    return adminList;
}

//  查看用户的地址 onlyAdmin
function getUserListFunc(uint256 _index) public view onlyAdmin returns(address) {
    return userList[_index].userAddr;
}

// 将用户升级为管理员
function addAdminFunc(address _addr) public onlyAdministrator {
    require(AddMyself, "You must add yourself first!");
    require(isExitUserAddressFunc(_addr), "User does not exist!");
    require(!isAdminFunc(_addr), "User is already an admin!");
    require(adminList.length < adminNum, "Administrator number is full!");
    userStruct[_addr].usertype = 1;
    adminList.push(_addr);    // 将管理员添加到数组中
    // 已经修改了管理员数量 (adminList.length +1)
    adminIndex[_addr].index = uint8(adminList.length) - 1; // 管理员索引 = 数组长度 - 1
}
// 将管理员移除
function removeAdminFunc(address _addr) public onlyAdministrator {
    require(_addr != administrator, "You are the administrator!");
    require(isExitUserAddressFunc(_addr) == true, "User does not exist!");
    require(isAdminFunc(_addr) == true, "User is not an administrator!");
    userStruct[_addr].usertype = 2; // 将管理员降级为普通用户
    // 将管理员移除   
    // 当管理员数 = 2 时, 直接将2号管理员移除
    if (adminList.length == 2) {
        adminList.pop(); 
    }else{
        // 当管理员数 >= 3 时, 将最后一个管理员移除的方法
        adminList[adminIndex[_addr].index] = adminList[adminList.length - 1]; // 将最后一个管理员地址替代该管理员地址的位置
        adminIndex[adminList[adminList.length - 1]].index = adminIndex[_addr].index; // 将最后一个管理员索引替代该管理员的索引
        delete adminIndex[_addr]; // 删除该管理员的索引
        adminList.pop(); // 将最后一个管理员地址从数组中移除
    }
}



// 账号密码管理 
// 
// 用户登陆  账号或地址登陆,返回用户账号密码是否正确与 成功时 用户的类型,用于判断是否为管理员
// 与合约本身无关系, 只是为了方便后台管理
function loginFunc(address _addr,string memory _account, string memory _password) public view returns(bool,uint8) {
    if (keccak256(abi.encode(_account)) == keccak256(abi.encode(""))){  // 如果传递的账号为空,则表示用地址登陆
        if (isExitUserAddressFunc(_addr)) {
            // 通过地址获取到存储到用户信息的密码与用户输入的密码进行比较
            if(keccak256(abi.encode(userStruct[_addr].password))== keccak256(abi.encode(_password))){
                return (true,userStruct[_addr].usertype);    // 登陆成功
            }
        }        
        return (false,3); // 登陆失败
    }else{
        if (isExitUserAccountFunc(_account)) {
            // 通过账号获取到存储到用户信息的密码与用户输入的密码进行比较
            if(keccak256(abi.encode(userStruct[accountMap[_account]].password))== keccak256(abi.encode(_password))){
                return (true,userStruct[_addr].usertype);    // 登陆成功
            }
        }  
        return (false,3); // 登陆失败
    }
}

// 用户查看自己的信息   // 只能返回三个值?!!!
function getMyselfInfoFunc() public view returns(address,string memory,uint8) {
    require(isExitUserAddressFunc(msg.sender), "You are not register!"); // 判断用户是否存在
    // require(isLogin, "You are not login!"); // 判断用户是否登陆
    return (
    userStruct[msg.sender].addr,
    userStruct[msg.sender].account,
    userStruct[msg.sender].usertype
    );
}
}