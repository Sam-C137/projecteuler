import { describe, expect, test } from "bun:test";

const dict = new Map([
    ["1", "one"],
    ["2", "two"],
    ["3", "three"],
    ["4", "four"],
    ["5", "five"],
    ["6", "six"],
    ["7", "seven"],
    ["8", "eight"],
    ["9", "nine"],
    ["10", "ten"],
    ["11", "eleven"],
    ["12", "twelve"],
    ["13", "thirteen"],
    ["14", "fourteen"],
    ["15", "fifteen"],
    ["16", "sixteen"],
    ["17", "seventeen"],
    ["18", "eighteen"],
    ["19", "nineteen"],
    ["20", "twenty"],
    ["30", "thirty"],
    ["40", "forty"],
    ["50", "fifty"],
    ["60", "sixty"],
    ["70", "seventy"],
    ["80", "eighty"],
    ["90", "ninety"],
    ["100", "onehundred"],
    ["200", "twohundred"],
    ["300", "threehundred"],
    ["400", "fourhundred"],
    ["500", "fivehundred"],
    ["600", "sixhundred"],
    ["700", "sevenhundred"],
    ["800", "eighthundred"],
    ["900", "ninehundred"],
    ["1000", "onethousand"],
]);

function run(target: number): number {
    let acc = 0;

    const safeLookup = (key: string): number => {
        const val = dict.get(key);
        if (!val) throw new Error("panic, map not have" + key);
        return val.length;
    };

    for (let i = 1; i <= target; i++) {
        if (i <= 20) {
            acc += safeLookup(i.toString());
        } else if (i < 100) {
            const tens = i.toString().slice(0, 1) + "0";
            acc += safeLookup(tens);
            const units = i.toString().slice(1);
            if (units !== "0") acc += safeLookup(units);
        } else if (i < 1000) {
            const hundreds = i.toString().slice(0, 1) + "00";
            acc += safeLookup(hundreds);
            const remainder = i % 100;
            if (remainder !== 0) {
                acc += "and".length;
                if (remainder <= 20) {
                    acc += safeLookup(remainder.toString());
                } else {
                    const tens = i.toString().slice(1, 2) + "0";
                    acc += safeLookup(tens);
                    const units = i.toString().slice(2);
                    if (units !== "0") acc += safeLookup(units);
                }
            }
        } else {
            acc += safeLookup(i.toString());
        }
    }

    return acc;
}

describe("number letter counts", () => {
    test("under 1000", () => {
        expect(run(1000)).toBe(21124);
    });
});
