function generateData(n){
    const baseNumber = '082' ;
    let customers = [];
    let mobileNumber;

    for (let i = 0 ; i < n; i++){
        mobileNumber = baseNumber + (String(i).padStart(9, '0'));
        customers.push(mobileNumber)
    }
    return customers;
}

function sendPromoDiscount(input1, input2){
    for(let i = 0; i < input1.length; i++){
        console.log('Sending Promo to ' + input1[i]);
    }
    console.log('Its Finished to send Promo to ' + input1.length + 'Customers');

    for (let i = 0; i < input2.length; i++){
        console.log('Sending Discount to ' + input2[i]);
    }
    console.log('Its Finished to send Discount to ' + input2.length + ' Customers');
}

let dataPromo = generateData(100000);
let dataDiscount = generateData(1000);

sendPromoDiscount(dataPromo, dataDiscount);

