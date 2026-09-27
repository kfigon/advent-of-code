import { describe, expect, it } from 'vitest'
import {parse, solveP1, solveP2} from './d1'

describe("d1", ()=> {
    const data = `3   4
4   3
2   5
1   3
3   9
3   3`

    it('p1', () => {
        const got = solveP1(parse(data)!);
        expect(got).toBe(11)        
    })

    it('p2', () => {
        const got = solveP2(parse(data)!);
        expect(got).toBe(31)        
    })
})