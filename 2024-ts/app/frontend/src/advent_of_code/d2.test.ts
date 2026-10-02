import { describe, expect, it } from 'vitest'
import {solve} from './d2'

describe("d2", ()=> {
    const data = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

    it('solve', () => {
        const got = solve(data, 1, 3);
        expect(got).toBe({p1: 2, p2: 4})        
    })
})