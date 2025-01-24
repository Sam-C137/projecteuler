import { describe, expect, test } from "bun:test";

function run(maxpower: number) {
    const exp = 1n << BigInt(maxpower);
    let acc = 0;

    for (const n of exp.toString(10)) {
        acc += Number(n);
    }

    return acc;
}

describe("power digit sum", () => {
    test("15", () => {
        expect(run(15)).toBe(26);
    });

    test("1000", () => {
        expect(run(1000)).toBe(1366);
    });
});
