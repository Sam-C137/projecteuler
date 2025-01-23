import { describe, expect, test } from "bun:test";

function run(size: number): number {
    const grid = get_grid(size);

    for (let r = 0; r < grid.length; r++) {
        grid[0][r] = 1;
        grid[r][0] = 1;
    }

    for (let r = 0; r < grid.length; r++) {
        for (let c = 0; c < grid[r].length; c++) {
            if (r > 0 && c > 0) {
                grid[r][c] = grid[r - 1][c] + grid[r][c - 1];
            }
        }
    }

    return grid[size][size];
}

function get_grid(size: number): number[][] {
    return new Array(size + 1)
        .fill(null)
        .map(() => new Array(size + 1).fill(null).fill(0));
}

describe("lattice paths", () => {
    test("2", () => {
        expect(run(2)).toBe(6);
    });

    test("20", () => {
        expect(run(20)).toBe(137846528820);
    });
});
