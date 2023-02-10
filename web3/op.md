npx hardhat run --network localhost ./scripts/index.js
npx hardhat run --network localhost ./scripts/deploy.js
npx hardhat run --network localhost ./scripts/inter-box.js
npx hardhat run --network goerli ./scripts/trans.js
npx hardhat run --network goerli ./scripts/deploy.js

npx hardhat run --network mainnet ./scripts/deploy.js
npx hardhat run --network goerli ./scripts/deploy.js


npx hardhat run --network goerli ./scripts/inter-box.js
## compile
npx hardhat compile


0x3072639306798dED58408e625e5536f17503892B
to 
 0xC2339bAF3FffE5dC1f9B7d344F902747369aE5fA
const box = await Box.attach('0xC2339bAF3FffE5dC1f9B7d344F902747369aE5fA');




npx hardhat console --network goerli

(await ethers.provider.getBalance("0x19fC239A0D676f8DDEE7555b453dD8eD77732C41")).toString()
(await ethers.provider.getBalance("0x3072639306798dED58408e625e5536f17503892B")).toString()
