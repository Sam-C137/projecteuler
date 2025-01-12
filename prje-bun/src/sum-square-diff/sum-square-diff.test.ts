import { describe, expect, test } from "bun:test";

function run(upperbound = 10): number {
    return [
        new Array(upperbound)
            .fill(null)
            .map((_, idx) => idx + 1)
            .reduce(
                ([x, y], curr) => {
                    return [x + curr, y + curr ** 2];
                },
                [0, 0],
            ),
    ].reduce((_, [x, y]) => {
        return x * x - y;
    }, 0);
}

describe("sum square difference", () => {
    test("10", () => {
        expect(run()).toBe(2640);
    });

    test("100", () => {
        expect(run(100)).toBe(25164150);
    });
});
