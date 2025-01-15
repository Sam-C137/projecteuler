import { describe, expect, test } from "bun:test";

function run(target: number): number {
    for (let c = 0; c < target; c++) {
        for (let b = 0; b < target; b++) {
            for (let a = 0; a < target; a++) {
                if (a + b + c === target && a * a + b * b === c * c) {
                    return a * b * c;
                }
            }
        }
    }

    return 0;
}

describe("special pythagorean triplet", () => {
    test("default", () => {
        expect(run(1000)).toBe(31875000);
    });
});
