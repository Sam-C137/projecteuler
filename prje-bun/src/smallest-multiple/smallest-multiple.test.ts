import {describe, expect, test} from "bun:test";

function run(upperbound = 10) {
    let smallest =1
    for (let i = 1; i <= upperbound; i++) {
        smallest = lcm(smallest, i)
    }
    return smallest
}

function gcd(a: number, b: number): number {
    while (b !== 0) {
        [a, b] = [b, a%b];
    }

    return a
}

function lcm(a: number, b: number): number {
    return (a * b) / gcd(a, b)
}

describe('smallest multiple', () => {
    test("default", () => {
        expect(run()).toBe(2520)
    })

    test("with arg 1000", () => {
        expect(run(20)).toBe(232792560)
    })
});