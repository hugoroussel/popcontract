pragma solidity ^0.4.0;


contract mortal {
    /* Define variable owner of the type address */
    address owner;

    /* This function is executed at initialization and sets the owner of the contract */
    function mortal() { owner = msg.sender; }

    /* Function to recover the funds on the contract */
    function kill() { if (msg.sender == owner) selfdestruct(owner); }
}

contract popcontract is mortal {

    public string nameOfParty;
    public string locationOfParty;
    public uint numberOfOrganizers;
    public address[] organizersAdresses;


  function setConfiguration(string name, string place, uint organizers){
    nameOfParty = name ;
    locationOfParty = place;
    numberOfOrganizers = organizers;

  }

  function setOrganizersAddresses(address[] data){
    organizersAdresses.length = numberOfOrganizers;
    organizersAdresses = data;
  }

    function getName() constant returns (string) {
        return nameOfParty;
    }
    function getPlace() constant returns (string) {
        return locationOfParty;
    }
    function getOrganizersAddresses() constant returns (address[]){
        return organizersAdresses;
    }

}
