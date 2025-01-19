import { describe, expect, test } from "bun:test";

function run(max: number): number {
    let found = 0;
    let n = 1;

    do {
        let n1 = n;
        let n2 = n + 1;

        if (n1 % 2 === 0) {
            n1 /= 2;
        } else {
            n2 /= 2;
        }

        let factors = get_factors(n1).length * get_factors(n2).length;

        if (factors >= max) {
            found = (n * (n + 1)) / 2;
            break;
        }
        n++;
    } while (true);

    return found;
}

function get_factors(n: number): number[] {
    let out: number[] = [];

    for (let i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            out.push(i);

            if (i * i !== n) {
                out.push(n / i);
            }
        }
    }

    return out;
}

describe("highly divisible triangular number", () => {
    test("5", () => {
        expect(run(5)).toBe(28);
    });

    test("500", () => {
        expect(run(500)).toBe(76576500);
    });
});
