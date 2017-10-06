# Proof of Personhood on the Ethereum blockchain
Concile anonymity and accoutability using proof of Personhood tokens and the Ethereum blockchain

## Setting up the configuration of the party

### Step 1, setting up the paramaters :

Organizer 1 (admin) deploys the popcontract and sends the configuration of the party using a private function, registering the different data needed : name, place, date and the Ethereum public addresses of the n other organizers. The contract passes in the state "configuration set".

### Step 2, signing the configuration:

Each organizer individually signs off individually the configuration using is public key (the contract matches the address with the initial addresses of the configuration).
The admin then calls a function to sign the whole configuration. The contract passes in the state "configuration signed".
NOTE : In the future modify the function such that the other organizers can sign using their private key, thus removing the need of the admin to sign it manually.

### Step 3, deposit of public keys of the participants :

Every organizer sends a (sorted) list of public attendees keys to the contract. Until the limit time/date, they are free to update the list.
Once one set of key has been deposited the contract passes in the state "keyDeposit".

### Step 4, reach consensus :

The contract decides which set of keys to keep using either :

* The last deposit with the last set is kept
* Take the set which is sent by the majority of organizers (with or without threshold)
* Wait for consensus (risk of deadlock)

Once this is done, the contract enters the state "Locked".

## How to deploy using the truffle framework :

http://truffleframework.com/

Run a local version of the Ethereum blockchain using :
`testrpc`

Then : 
`truffle compile`

`truffle migrate`
