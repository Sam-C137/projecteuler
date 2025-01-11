import {describe, expect, test} from "bun:test";

function run(target : number): number{
    let max = 0;


    for (let i = target; i > 0; i++) {
        for (let j = i; j > 0; j++) {
            const prod = i * j;

            if (prod < max) break;

            if(isPalindrome(max)) max = prod
        }
    }

    return max
}

function isPalindrome(num: number) {
    return String(num) === String(num).split('').reverse().join('')
}

describe('largest palindrome product', () => {
    test("Largest 3 digit number", () => {
        expect(run(999)).toBe(906609)
    })
});