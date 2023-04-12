const telco = ['Telkom', 'Indosat', 'XL', 'Verizon', 'AT&T', 'Nippon',
                'Vodafone', 'Orange', 'KDDI', 'Telefonica', 'T-Mobile'];

function findCompany(array, input) {
    for (let i= 0; i < array.length; i++) {
        if (array[i] === array[input]) {
            console. log('Company Found : ' + array[input]);
            break;
        }

        console.log('Searching Company...' + (i+1));
    }
}
const company = Math.round(Math.random() * 10);
findCompany(telco, company);