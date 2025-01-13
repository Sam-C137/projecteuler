import { describe, expect, test } from "bun:test";

function run(target = 6): number {
    let count = 0;
    let num = 2;

    while (count < target) {
        if (is_prime(num)) {
            count++;
            if (count === target) return num;
        }

        num = num === 2 ? 3 : num + 2;
    }

    return num;
}

function is_prime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;

    for (let i = 3; i * i <= num; i += 2) {
        if (num % i === 0) return false;
    }

    return true;
}

describe("10_001st prime number", () => {
    test("default", () => {
        expect(run()).toBe(13);
    });

    test("10_001st", () => {
        expect(run(10_001)).toBe(104743);
    });
});
