// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Count{
    uint256 public  count;

    function Add() public  returns (uint256)    {
        count+=1;
        return count;
    }

}