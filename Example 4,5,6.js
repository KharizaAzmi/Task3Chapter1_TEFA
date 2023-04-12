const address = 'DKI Jakarta';
const addresses = new Array(10).fill(address)

function findAddress(addresses) {
	 const tx = performance.now();
	 console.log('The Default Address is ' + (addresses[0]));
	 const ty = performance.now();
	 console.log('The Performance is ' + (ty-tx) + ' ms');
}
findAddress(addresses);
