// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AccountClass{

// string 分类数组
// string[] private accountTypeList; // 存储分类结构体的数组
// 通过地址获取分类数组
mapping (address => string[]) public accountClassListMap;
// 通过地址获取分类数组的索引
mapping(address => mapping(string => uint256)) private accountClassIndexMap;  // 存储分类结构体的索引

// 添加分类事件
event addAccountClassEvent(address indexed userAddr,string accountClass);
// 修改分类事件
event modifyAccountClassEvent(address indexed userAddr,string accountClass,string oldAccountClass); 
// 删除分类事件
event deleteAccountClassEvent(address indexed userAddr,string accountClass);
// 添加账号(Key)事件
event addKeyEvent(address indexed userAddr,string accountClass,string key,string createtime);
// 删除账号(Key)事件
event deleteKeyEvent(address indexed userAddr,string accountClass,string key);
// 修改账号(Key)事件
event modifyKeyEvent(address indexed userAddr,string accountClass,string key);
// 移动账号(Key)事件
event moveKeyEvent(address indexed userAddr,string accountClass,string key,string oldAccountClass);

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
    // 添加分类索引 数组已经提前加了 1,所以索引得减一
    accountClassIndexMap[msg.sender][_accountClass] = accountClassListMap[msg.sender].length - 1;

    emit addAccountClassEvent(msg.sender,_accountClass);
}

// 判断分类是否存在 
function isAccountClassFunc(string memory _accountClass) internal view returns(bool){
    if(accountClassListMap[msg.sender].length == 0){return false;}  // 判断分类数组是否为空,为空则肯定不存在
    // return keccak256(abi.encode(accountTypeList[accountClassIndexMap[msg.sender][_accountClass]])) == keccak256(abi.encode(_accountClass));
    return (str(_accountClass,"default") || accountClassIndexMap[msg.sender][_accountClass] != 0);    // 判断分类索引是否为0,为0则不是default就肯定不存在
}

// 返回所有分类(数组)
function returnClassAll() public view returns (string[] memory){
    return accountClassListMap[msg.sender];
}

// 修改分类名
function changeClass(string memory _oldClass,string memory _newClass) public {
    // 判断是否存在该分类
    require(isAccountClassFunc(_oldClass), "Account type does not exist!");
    require(!isAccountClassFunc(_newClass), "Account type is exist!");
    // 判断是否是默认分类
    require(!str(_oldClass, "default"), "Default type cannot be modified!");
    require(!str(_newClass, "default"), "Default type cannot be modified!");
    // 修改分类名,只是修改了该文件夹的名字
    uint256 _index = accountClassIndexMap[msg.sender][_oldClass]; 
    accountClassListMap[msg.sender][_index] = _newClass;    // 换名称
    accountClassIndexMap[msg.sender][_newClass] = _index;  // 修改索引
    delete accountClassIndexMap[msg.sender][_oldClass]; // 删除分类索引

    // 如果分类下有Key,则必须修改分类下的 key
    if (keyListMap[msg.sender][_oldClass].length > 0) {
        for (uint i = 0; i < keyListMap[msg.sender][_oldClass].length; i++) {
            KeyStruct memory key = keyListMap[msg.sender][_oldClass][i];
            keyListMap[msg.sender][_newClass].push(KeyStruct({
                key:key.key,
                password:key.password,
                time:key.time
            }));
            keyIndexMap[msg.sender]["default"][key.key] = keyListMap[msg.sender]["default"].length;
        }
        delete keyListMap[msg.sender][_oldClass];
    }

    emit modifyAccountClassEvent(msg.sender,_newClass,_oldClass);
}

// 删除分类
function deleteAccountClassFunc(string memory _accountClass) public {
    // 判断是否存在该分类
    require(isAccountClassFunc(_accountClass), "Account type does not exist!");
    // 默认分类不能删除
    require(!str(_accountClass, "default"), "Default type cannot be deleted!");
    
    // 删除分类
     if (accountClassListMap[msg.sender].length == 2) {
        accountClassListMap[msg.sender].pop(); 
    }else{
        // 当分类数 >= 3 时, 将最后一个移除的方法
        //    要删除的分类索引
        uint256 deleteIndex = accountClassIndexMap[msg.sender][_accountClass]; 
        // 用最后一个分类占据要删除的分类的位置
        string memory _last = accountClassListMap[msg.sender][accountClassListMap[msg.sender].length - 1];
        accountClassListMap[msg.sender][deleteIndex] = _last;
        // 删除最后一个分类
        accountClassListMap[msg.sender].pop();      // 删除最后一个分类
        accountClassIndexMap[msg.sender][_last] = deleteIndex;  // 修改索引
        delete accountClassIndexMap[msg.sender][_accountClass]; // 删除分类索引
    }

     // 如果分类下有Key,则必须将分类下的 key 移动到 default
    if (keyListMap[msg.sender][_accountClass].length > 0) {
        for (uint i = 0; i < keyListMap[msg.sender][_accountClass].length; i++) {
                KeyStruct memory key = keyListMap[msg.sender][_accountClass][i];
                keyListMap[msg.sender]["default"].push(KeyStruct({
                    key:key.key,
                    password:key.password,
                    time:key.time
                }));
                keyIndexMap[msg.sender]["default"][key.key] = keyListMap[msg.sender]["default"].length;
            }
        delete keyListMap[msg.sender][_accountClass];
    }

    emit deleteAccountClassEvent(msg.sender,_accountClass);
}

// 判断两个字符串是否相等
function str(string memory _str1,string memory _str2) internal pure returns(bool){
    return keccak256(abi.encode(_str1)) == keccak256(abi.encode(_str2));
}

// 分类下的 Key

struct KeyStruct{ // 保存的账号和密码
    string key;
    string password;
    string time;
}
// 通过地址和分类名 获取Key数组
mapping (address => mapping(string => KeyStruct[])) public keyListMap;
// 地址 分类名 账号列表(Key) 账号索引
mapping(address => mapping(string => mapping(string => uint256))) public keyIndexMap; 

// 添加账号 不要索引
function addKeyFunc(string memory _accountClass,string memory _key,string memory _pass,string memory _time) public {
    require(isAccountClassFunc(_accountClass), "Account type does not exist!");

    // 判断是否已经存在该账号
    require(!isKeyFunc(_accountClass,_key), "Account Key already exists!");

    // 添加账号
    keyListMap[msg.sender][_accountClass].push(KeyStruct({
        key:_key,
        password:_pass,
        time: _time
    }));
    // 添加分类下Key索引,不要索引,使用第几个
    keyIndexMap[msg.sender][_accountClass][_key] = keyListMap[msg.sender][_accountClass].length; // 不要索引!!!!!!

    emit addKeyEvent(msg.sender,_accountClass,_key,_time);
}

// 判断账号(Key)是否存在
function isKeyFunc(string memory _accountClass,string memory _key) public view returns(bool){
    uint256 index = keyIndexMap[msg.sender][_accountClass][_key];   // 获取该Key的位置而不是索引
    return index != 0;  // 个数绝对不会为0,因为添加时已经判断了
}

// 返回所有Key(数组)
function returnClassAll(string memory _accountClass) public view returns (KeyStruct[] memory){
    return keyListMap[msg.sender][_accountClass];
}

// 删除Key
function deleteKeyFunc(string memory _accountClass,string memory _key) public {
    require(isAccountClassFunc(_accountClass), "Account type does not exist!");
    require(isKeyFunc(_accountClass,_key), "Account key does not exist!");
   
    uint256 index = keyIndexMap[msg.sender][_accountClass][_key];   // 获取该Key的位置而不是索引
    delete keyListMap[msg.sender][_accountClass][index - 1 ]; 
    delete keyIndexMap[msg.sender][_accountClass][_key];

    emit deleteKeyEvent(msg.sender,_accountClass,_key);
}

// 修改Key的密码
function updateKeyFunc(string memory _accountClass,string memory _key,string memory _pass) public {
    require(isAccountClassFunc(_accountClass), "Account type does not exist!");
    require(isKeyFunc(_accountClass,_key), "Account key does not exist!");
    keyListMap[msg.sender][_accountClass][keyIndexMap[msg.sender][_accountClass][_key]].password = _pass;

    emit modifyKeyEvent(msg.sender,_accountClass,_key);
}

// 移动某个 Key 到其他分类
function moveFunc(string memory _oldClass,string memory _newClass,string memory _key) public {
    require(isAccountClassFunc(_oldClass), "Account type does not exist!");
    require(isAccountClassFunc(_newClass), "Account type does not exist!");
            uint256 index = keyIndexMap[msg.sender][_oldClass][_key];
            KeyStruct memory key = keyListMap[msg.sender][_oldClass][index - 1];
            keyListMap[msg.sender][_newClass].push(KeyStruct({
                key:key.key,
                password:key.password,
                time:key.time
            }));

        keyIndexMap[msg.sender][_newClass][key.key] = keyListMap[msg.sender][_newClass].length;
        delete keyListMap[msg.sender][_oldClass][index - 1 ];
    emit moveKeyEvent(msg.sender,_oldClass,_newClass,_key);
}
}