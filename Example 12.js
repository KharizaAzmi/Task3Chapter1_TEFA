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

function sendPromoDiscount(input){
    for(let i = 0; i < input.length; i++){
        console.log('Sending Promo to ' + input[i]);
    }
    console.log('Its Finished to send Promo to ' + input.length + 'Customers');

    for (let i = 0; i < input.length; i++){
        console.log('Sending Discount to ' + input[i]);
    }
    console.log('Its Finished to send Discount to ' + input.length + ' Customers');
}

let data = generateData(1000);
sendPromoDiscount(data);

