import GenericSolver from "./genericSolver"

export default function D1() {
    return GenericSolver(parse, solveP1, solveP2)
}

type ParseRes = {
    left: number[]
    right: number[]
}
function parse(v :string): ParseRes | null {
    const lines = v.split('\n')
    const splitedLines = lines.map(v => v.split(/\s+/)).map(pair => [Number(pair[0]), Number(pair[1])])
    const left = splitedLines.map(v=> v[0])
    const right = splitedLines.map(v=> v[1])
    return {left,right}
}

function solveP1({left: leftGiven, right: rightGiven}: ParseRes): number | null {
    const left = leftGiven.slice();
    const right = rightGiven.slice();

    left.sort()
    right.sort()
    if (left.length !== right.length) return null;

    let res = 0;
    for(let i = 0; i < left.length; i++) {
        res += Math.abs(left[i] - right[i]);
    }
    
    return res;
}

function solveP2({left, right}: ParseRes): number | null {
    const occurences = new Map<number,number>();
    right.forEach(v => occurences.set(v, (occurences.get(v) ?? 0) + 1))
    return left.map(v => v * (occurences.get(v) ?? 0)).reduce((a,b) => a+b)
}