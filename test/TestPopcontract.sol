pragma solidity ^0.4.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/popcontract.sol";

contract TestFirst {
  Popcontract popcontract = Popcontract(DeployedAddresses.Popcontract());

}
