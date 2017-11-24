# Proof of Personhood on the Ethereum blockchain
Concile anonymity and accountability using proof of personhood tokens and the ethereum blockchain.

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

## How to deploy using Go :

You will need to have geth and to run a full node. The current specified path is for the rinkeby network. Organizers will need to create accounts using geth, and the administrator
of the party needs to input is JSON file as well as his password into the app.go file. Manual input of nonce is required in the app.go file as well. To get current nonce run :

`geth attach ipc:"path/to/your/geth.ipc"`

`web3.eth.getTransactionCount("administrator address")`

To run, clone project & in directory run :

`go build && ./popcontract`

Currently displays a demo of a party setup.

To apply modification to popcontract.sol :

`abigen --sol=popcontract.sol --pkg=main --out=pop.go`

![Image](/doc/fsm.jpg)


## To do :

* Test with non empty keyset
* Connect CLI to app.go
* Write consensus function
* Final statement function
