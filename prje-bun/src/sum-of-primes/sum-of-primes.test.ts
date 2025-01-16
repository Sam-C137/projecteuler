import { describe, expect, test } from "bun:test";

function run(max = 10): number {
    let acc = 0;

    for (let i = 0; i < max; i++) {
        if (is_prime(i)) acc += i;
    }

    return acc;
}

function is_prime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;

    for (let i = 3; i ** 2 <= num; i += 2) {
        if (num % i === 0) return false;
    }

    return true;
}

describe("sum of primes", () => {
    test("below 10", () => {
        expect(run()).toBe(17);
    });

    test("below 2_000_000", () => {
        expect(run(2_000_000)).toBe(142913828922);
    });
});
