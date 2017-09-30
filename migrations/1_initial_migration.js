var Migrations = artifacts.require("./Migrations.sol");

module.exports = function(deployer) {
  // Deploy the migrations contract as our only task
  deployer.deploy(Migrations);
};
