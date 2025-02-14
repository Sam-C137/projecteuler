import { describe, expect, test } from "bun:test";

function run(target: number): number {
    const lookup = new Map<number, number>;
    const set = new Set<number>;

    for (let i = 0; i < target; i++) {
        const div = d(i);
        lookup.set(i, div);
        if (lookup.get(div) === i) {
            if (div !== i) {
                set.add(div);
                set.add(i);
            }
        }
    }

    return Array.from(set).reduce((a, b) => a + b);
}

function d(num: number): number {
    let sum = 0;

    for (let i = 1; i < num; i++) {
        if (num % i === 0) sum += i;
    }

    return sum;
}

describe("amicable numbers", () => {
    test("under 10k", () => {
        expect(run(10_000)).toBe(31626);
    });
});
