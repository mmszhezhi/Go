// scripts/index.js

async function local_op(){
    const accounts = await ethers.provider.listAccounts();
    console.log(accounts);
    const address = '0x5FbDB2315678afecb367f032d93F642f64180aa3';
    const Box = await ethers.getContractFactory('Box');
    const box = await Box.attach(address);
    // Call the retrieve() function of the deployed Box contract
    const value = await box.retrieve();
    console.log('Box value is', value.toString());

    // Send a transaction to store() a new value in the Box
    await box.store(23);

    const value2 = await box.retrieve();
    console.log('Box value is', value2.toString());
}

async function box_op (){
    // goerli contract
    const addr = "0x02AE9D41dC084e2D4E85cd2C253090EC9797dc28"
    const write = false

    // local upgradeble
    // const addr = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
    // const write = true


    console.log("fuck box contract interact")
    const Box = await ethers.getContractFactory('Box');
    const box = await Box.attach(addr);

    if(write){
        await box.store(42);
    }
    const value2 = await box.retrieve();
    console.log('Box value is', value2.toString());
}


async function box_upgrade_op (){
    // goerli contract
    // const addr = "0x02AE9D41dC084e2D4E85cd2C253090EC9797dc28"
    // const write = false

    // local upgradeble
    const addr = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
    const write = true


    console.log("fuck box contract interact")
    const Box = await ethers.getContractFactory('BoxV2');
    const box = await Box.attach(addr);

    await box.increment();
    // if(write){
    //     await box.store(42);
    // }
    const value2 = await box.retrieve();
    console.log('Box value is', value2.toString());
}

async function main () {
    await box_op()
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });