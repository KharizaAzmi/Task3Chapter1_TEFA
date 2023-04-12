const foods = ['Sate', 'Bakso', 'Dimsum', 'Rames'];
const drinks = ['Jeruk','Teh', 'Kelapa','Cendol'];

function logPairs(arrayl, array2, words) {
    let counter = 0;
    for (let i = 0; i < arrayl.length; i++) { //0(n)
        for (let j = 0; j < array2.length; j++) { //0(n = m)
            counter++; //0(n = m)
            console.log(words + ' ' + counter + ' ' + arrayl[i] + ' dan ' + array2[j]); //0(n = m)
        }
    }
}

logPairs(foods, drinks, 'Menu');

