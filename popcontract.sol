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



   struct publicKeySet{
     address sender;
     address[] keySet;
   }


   publicKeySet [] public allSets;
   publicKeySet public finalKeySet;
   contractState public currentState;
   string public nameOfParty;
   string public locationOfParty;
   uint public endOfParty;
   uint numberOfOrganizers;
   address[] public organizersAdresses;
   address[] public signedConfiguration;
   address nullAddress = 0x0000000000000000000000000000000000000000;
   address testAddress = 0x0000000000000000000000000000000000000012;
   bool public signed;


  modifier onlyState(contractState expectedState){
    if(expectedState == currentState){_;}
    else{revert();}
  }
  modifier beforeDeadline() { if (now <= endOfParty) _; }
  modifier afterDeadline() { if (now >= endOfParty) _; }

  //constructor to initialize contract state
  function popcontract(){
    currentState = contractState.initialState;
  }

  //to do : add modifier so only accessible in initialState
  function setConfiguration (
    string place,
    uint organizers,
    address[] data,
    uint durationInMinutes)
    onlyState(contractState.initialState) returns (bool)
    {
    if(msg.sender == owner){
    endOfParty = now + durationInMinutes * 1 minutes;
    locationOfParty = place;
    numberOfOrganizers = organizers;
    organizersAdresses.length = numberOfOrganizers;
    organizersAdresses = data;
    currentState = contractState.configurationSet;
    signedConfiguration.length = organizers;
    signed = false;
    return true;
    }
    else{
    revert();
    }
  }

  function setName(string _name) onlyState(contractState.initialState) returns (bool){
    nameOfParty = _name;
  }

  function getOrganizersAddresses() constant returns (address[]){
      return organizersAdresses;
  }

  function getSignedConfiguration() constant returns (address[]){
      return signedConfiguration;
  }


//Security risk here
  function configSignOrganizers() onlyState(contractState.configurationSet) returns (address) {

      uint index = 0;

      for(uint i=0;i< numberOfOrganizers;i++){
        if(msg.sender == organizersAdresses[i]){
          index = i;
        }
      }

      signedConfiguration.push(organizersAdresses[index]);
  }

  //after the signing of the configuration by the organizers the admin changes the state of the contract
  function signWholeConfiguration() onlyState(contractState.configurationSet) returns (bool){
    bool correct = false;
    if(msg.sender == owner){
      for(uint i=0; i< numberOfOrganizers; i++){
        if(signedConfiguration[i+3]==nullAddress){
          revert();
        } else{
          correct = true;
          signed = true;
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


  function depositPublicKeys (address [] _publicKeySet) onlyState(contractState.configurationSigned) beforeDeadline returns (bool){
    if(isOrganizer(msg.sender)){
       allSets.push(publicKeySet(msg.sender, _publicKeySet));
       if(currentState != contractState.keyDeposited){
       currentState = contractState.keyDeposited;
     }
       return true;
    }
    return false;
  }

  //Just returns the last keySet (temporary)
  function publicKeyConsensus() onlyState(contractState.keyDeposited) afterDeadline returns (bool){
    if(allSets.length != 0 && msg.sender == owner){
    finalKeySet=allSets[allSets.length];
    currentState = contractState.locked;
    return true;
  }
  return false;
  }










}
