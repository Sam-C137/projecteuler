import { describe, expect, test } from "bun:test";

function run() {
    let sundays = 0;
    let start = new Date(1901, 0, 1);
    let end = new Date(2000, 11, 31);

    let curr = start;
    while (curr <= end) {
        if (curr.getDay() === 0 && curr.getDate() === 1) {
            sundays++;
        }

        curr = new Date(curr.getFullYear(), curr.getMonth() + 1, 1);
    }

    return sundays;
}

describe("counting sundays", () => {
    test("default", () => {
        expect(run()).toBe(171);
    });
});
