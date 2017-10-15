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

  //States of the contract
   enum contractState {
     initialState,
     configurationSet,
     configurationSigned,
     keyDeposited,
     locked
   }

   //can change some of the variables to public to delete getters
   contractState public currentState;
   string nameOfParty;
   string locationOfParty;
   uint numberOfOrganizers;
   address[] organizersAdresses;
   address[] signedConfiguration;
   address nullAddress = 0x0000000000000000000000000000000000000000;

   //for the moment the public addresses are on the format Ethereum
   mapping (address => address[]) public publicKeySet;


  modifier onlyState(contractState expectedState){
    if(expectedState == currentState){_;}
    else{revert();}
  }

  //constructor to initialize contract state
  function popcontract(){
    currentState = contractState.initialState;
  }

  //to do : add modifier so only accessible in initialState
  function setConfiguration (string name, string place, uint organizers, address[] data) onlyState(contractState.initialState) returns (bool) {
    if(msg.sender == owner){
    nameOfParty = name ;
    locationOfParty = place;
    numberOfOrganizers = organizers;
    organizersAdresses.length = numberOfOrganizers;
    organizersAdresses = data;
    currentState = contractState.configurationSet;
    return true;
    }else{
    revert();
    }

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

  function configSignOrganizers() onlyState(contractState.configurationSet) returns (bool){
    //verify if msg.sender address is in organizersAdresses and sign if it is case
    for (uint i; i< numberOfOrganizers; i++){
      if(msg.sender == organizersAdresses[i]){
        signedConfiguration[i] = msg.sender;
        return true;
      }
    }
    return false;
  }

  //after the signing of the configuration by the organizers the admin changes the state of the contract
  function signWholeConfiguration() onlyState(contractState.configurationSet) returns (bool){
    bool correct = false;
    if(msg.sender == owner){
      for(uint i =0; i< numberOfOrganizers; i++){
        if(signedConfiguration[i]==nullAddress){
          revert();
        } else{
          correct = true;
        }
      }
      if(correct){
        currentState = contractState.configurationSigned;
        return true;
      }
    }
  }

  function isOrganizer(address sender) onlyState(contractState.configurationSigned) returns (bool){
    for(uint i=0; i<numberOfOrganizers; i++){
      if(organizersAdresses[i]==sender){
        return true;
      }
    }
    return false;
  }

  function depositPublicKeys (address [] _publicKeySet) onlyState(contractState.configurationSigned) returns (bool){
    if(isOrganizer(msg.sender)){
       _publicKeySet = publicKeySet[msg.sender];
       return true;
    }
    return false; 
  }







}
