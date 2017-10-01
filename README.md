# Proof of Personhood on the Ethereum blockchain
Concile anonymity and accoutability using proof of Personhood tokens and the Ethereum blockchain

# Setting up the configuration

## Step 1 :

Organizer 1 deploys the PoP contracts and sends configuration of the party using a private function, registering the different data needed : name, place, date and the Ethereum public addresses of the n other organizers.

## Step 2 :

Every organizer signs off the config by calling a public function only accessible by the public addresses mentioned earlier. Once, all organizers signed/voted, the contract enters a “set” state.

Possibility of voting as well on the way to reach consensus later on?

Every organizer sends a (sorted) list of public attendees keys to the contract.
The party ends once the organizers reach a consensus on which keys to push to the blockchain. Multiple ways to determine the set of keys kept :

The last deposit with the last set is kept
Take the set which is sent by the majority of organizers (with or without threshold)
Wait for consensus (risk of deadlock)
