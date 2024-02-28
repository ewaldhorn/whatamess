const myMap = new Map();

const reportOn = (map) => {
    if (map.size > 0) {
        let report = ``;
        map.forEach((value, key) => report += `${key}:${value},`);
        console.log("Map contains: ", `[${report}]`);
    } else {
        console.log("Map is empty!");
    }
}

const newline = () => console.log();

reportOn(myMap);

myMap.set(1, "Item");
newline();
reportOn(myMap);

myMap.set(1, "First Item");
myMap.set(2, "Second Item");
myMap.set(3, "Third Item");
myMap.set(4, "Fourth Item");
newline();
reportOn(myMap);

newline();