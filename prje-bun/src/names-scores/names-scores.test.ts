import { describe, expect, test } from "bun:test";

async function run() {
    const f = await Bun.file("0022_names.txt").text();
    return f
        .split(",")
        .map((s) => s.replaceAll('"', ""))
        .sort((a, b) => a.localeCompare(b))
        .reduce((prev, curr, idx) => {
            return (
                prev +
                curr.split("").reduce((p, c) => p + (c.charCodeAt(0) - 64), 0) *
                    (idx + 1)
            );
        }, 0);
}

describe("name scores", () => {
    test("default", async () => {
        expect(await run()).toBe(871198282);
    });
});
