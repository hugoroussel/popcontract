pragma solidity ^0.4.0;

contract Migrations {
  address public owner;
  // A function wih the signature last_completed_migration and returning a uint is necessary
  uint public last_completed_migration;

  modifier restricted() {
    if (msg.sender == owner) _;
  }

  function Migrations() {
    owner = msg.sender;
  }

  // A function with the signature `setCompleted(uint)` is required.
  function setCompleted(uint completed) restricted {
    last_completed_migration = completed;
  }

  function upgrade(address new_address) restricted {
    Migrations upgraded = Migrations(new_address);
    upgraded.setCompleted(last_completed_migration);
  }
}
