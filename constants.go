package main

/**
 * RPC通信的配置
 */
const RPCURL = "http://127.0.0.1:8332"
const RPCUSER = "user"
const RPCPASSWORD = "pwd"
const RPCVERSION = "2.0"

/**
 * 比特币节点的命令
 */
//获取比特币节点的区块数量
const GETBLOCKCOUNT = "getblockcount"

//获取节点的最新区块的hash值
const GETBESTBLOCKHASH  = "getbestblockhash"

const STOP = "stop"

const GETNEWADDRESS = "getnewaddress"

const GETBLOCKHASH = "getblockhash"

const GETBLOCK = "getblock"
