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

     string nameOfParty;
     string locationOfParty;
     uint numberOfOrganizers;
     address[] organizersAdresses;
     address[] signedConfiguration;
     address nullAddress = 0x0000000000000000000000000000000000000000;


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

  //not very elegant but works. To do : remove index
  //change this functions such that it returns the organizer index instead of the address registered
  function signConfig() returns (address){

    //verify if msg.sender address is in organizersAdresses
    bool pres = false;
    uint i=0;
    for (i; i< numberOfOrganizers; i++){
      if(msg.sender == organizersAdresses[i]){
        pres = true ;
      }
    }

    //add the address only if valid organizer.
    if(pres){
      signedConfiguration[i] = msg.sender;
      return msg.sender;
    } else{
      return nullAddress;
    }
  }

}
