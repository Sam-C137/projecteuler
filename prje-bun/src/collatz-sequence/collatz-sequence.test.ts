import { describe, expect, test } from "bun:test";

function run(upperbound: number): number {
    let maxSeq = 0;
    let maxOcc = 0;

    for (let i = 0; i < upperbound; i++) {
        const newSeq = seq(i);
        if (newSeq > maxSeq) {
            maxSeq = newSeq;
            maxOcc = i;
        }
    }

    return maxOcc;
}

function seq(start: number): number {
    let count = 1;

    while (start > 1) {
        if ((start & 1) === 0) {
            start = start / 2;
        } else {
            start = 3 * start + 1;
        }
        count++;
    }

    return count;
}

describe("max collatz sequence", () => {
    test("under 1m", () => {
        expect(run(1_000_000)).toBe(837799);
    });
});
