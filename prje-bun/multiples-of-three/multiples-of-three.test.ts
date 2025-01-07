import {describe, test, expect} from "bun:test"

function run(max?: number): number {
    return new Array(max ?? 10).fill(null).reduce((prev, _, idx)=>{
        return prev + ((idx % 3 === 0 || idx % 5 === 0) ? idx : 0)
    }, 0)
}

describe('multiples of three', () => {
    test("default", () => {
        expect(run()).toBe(23)
    })

    test("with arg 1000", () => {
        expect(run(1000)).toBe(233168)
    })
});