module.exports = {
  networks: {
    rinkeby: {
      network_id: 4,
      host: '127.0.0.1',
      port: 8545,
      gas: 4000000,
      from: "0x5600a6AEb296506d51Da2CAF592ac98b7879C968",
    },
  },
  rpc: {
    // Use the default host and port when not using rinkeby
    host: 'localhost',
    port: 8080,
  },
};
