# Proof of Personhood on the Ethereum blockchain
Concile anonymity and accoutability using proof of personhood tokens and the ethereum blockchain.

## Setting up the configuration of the party

### Step 1, setting up the paramaters :

Organizer 1 (admin) deploys the popcontract and sends the configuration of the party using a private function, registering the different data needed : name, place, duration and the Ethereum public addresses of the n other organizers.

### Step 2, signing the configuration:

Each organizer individually signs off individually the configuration using is public key (the contract matches the address with the initial addresses of the configuration).
The admin then calls a function to sign the whole configuration.
Note : Might be possible to modify the function such that the other organizers can sign using their private key, removing the need of the admin to sign it manually.

### Step 3, deposit of public keys of the participants :

Every organizer sends a (sorted) list of public attendees keys to the contract. Until the limit time/date, they are free to update the list.

### Step 4, reach consensus :

The contract decides which set of keys to register by taking the one submitted by most organizers. Other ways to reach consensus are also possible. The organizer could include the one choosen in the configuration of the party. Once this is done, the contract enters the state "Locked".

## How to deploy using the truffle framework :

http://truffleframework.com/

Run a local version of the Ethereum blockchain using :

`testrpc`

Then :

`truffle compile`

`truffle migrate`


## To do :

* Write tests
* Change the format of the publicKeySet into ed25519
* Write consensus function
