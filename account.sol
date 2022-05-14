// 需求:给 msg.sender 添加分类数据
// 添加的分类只属于该用户，不同的用户可以添加相同的分类

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AccountClass{

// string 分类数组
// string[] private accountTypeList; // 存储分类结构体的数组
// 通过地址获取分类数组
mapping (address => string[]) public accountClassListMap;
// 通过地址获取分类数组的索引
mapping(address => mapping(string => uint256)) private accountClassIndexMap;  // 存储分类结构体的索引

// 添加分类
function addAccountClassFunc(string memory _accountClass) public  {
    // 判断是否已经存在该分类
    require(!isAccountClassFunc(_accountClass), "Account type already exists!"); // 判断账号分类是否存在
    if (accountClassListMap[msg.sender].length == 0) {
        // 添加默认分类
        accountClassListMap[msg.sender].push("default");
        accountClassIndexMap[msg.sender][_accountClass] = 0;
    }
    // 添加分类
    // accountTypeList.push(_accountClass);
    accountClassListMap[msg.sender].push(_accountClass);
    // 添加分类索引
    accountClassIndexMap[msg.sender][_accountClass] = accountClassListMap[msg.sender].length ;
}

// 判断分类是否存在 
function isAccountClassFunc(string memory _accountClass) internal view returns(bool){
    if(accountClassListMap[msg.sender].length == 0){return false;}  // 判断分类数组是否为空,为空则肯定不存在
    // return keccak256(abi.encode(accountTypeList[accountClassIndexMap[msg.sender][_accountClass]])) == keccak256(abi.encode(_accountClass));
    return accountClassIndexMap[msg.sender][_accountClass] != 0;    // 判断分类索引是否为0,为0则肯定不存在
}

// 返回所有分类(数组)
function returnClassAll() public view returns (string[] memory){
    return accountClassListMap[msg.sender];
}

// 删除分类
function deleteAccountClassFunc(string memory _accountClass) public {
    // 判断是否存在该分类
    require(isAccountClassFunc(_accountClass), "Account type does not exist!");
    // 默认分类不能删除
    require(!str(_accountClass, "default"), "Default type cannot be deleted!");
  
}

// 判断两个字符串是否相等
function str(string memory _str1,string memory _str2) public pure returns(bool){
    return keccak256(abi.encode(_str1)) == keccak256(abi.encode(_str2));
}
}
