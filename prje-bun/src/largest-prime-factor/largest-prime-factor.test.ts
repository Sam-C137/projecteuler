import { describe, expect, test } from "bun:test";

function run(target: number): number {
    let max = 0;
    let n = target;

    while (n % 2 === 0) {
        max = 2;
        n /= 2;
    }

    for (let i = 3; i < Math.sqrt(n); i += 2) {
        while (n % i === 0) {
            max = i;
            n /= i;
        }
    }

    if (n > 2) max = n;

    return max;
}

describe("largest prime factor", () => {
    test("600_851_475_143", () => {
        expect(run(600_851_475_143)).toBe(6857);
    });
});
