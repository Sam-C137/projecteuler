import { describe, expect, test } from "bun:test";

function run(fib: bigint): bigint {
    if (fib < 2) return fib;

    let prev = BigInt(0);
    let curr = BigInt(1);
    let even = BigInt(0);

    while (curr < fib) {
        if ((curr & 1n) !== 1n) {
            even += curr;
        }

        if (prev + curr < prev) break;

        [prev, curr] = [curr, prev + curr];
    }

    return even;
}

describe("fibonacci", () => {
    test("passes", () => {
        expect(run(BigInt(4_000_000))).toBe(4613732n);
    });
});
