# Plutus Uniswap Demo

## Pre requisites
- Golang minimum 1.17.0
- npm minimum 8.0.0
- ghc 8.10.5
- cabal 3.4.0.0


### Project structure
- gui-app (Source code for the Web-app GUI)
- middleware-app (Source code for the Server who will serve the static Web-app files and as a proxy between the Plutus PAB and the GUI)
- uniswap-pab (Source Code containing the PAB (Plutus Application Backend) demo and the smart contract)

## Start the demo (development mode)
1. PAB
  - navigate to uniswap-pab `cd uniswap-pab`
  - start the application `cabal run uniswap-pab`
  - wait until the Wallets are initialized __log Message: `Uniswap user contract started for Wallet 4`__
2. Start Middleware
  - navigate to middleware-app `cd middleware-app`
  - start the application with `go run main.go`
3. Start the Web Application
  - navigate to gui-app `cd gui-app`
  - start the application with `npm install && npm run start`


## Demo from Plutus Pioneer programm by Lars Brunjes



