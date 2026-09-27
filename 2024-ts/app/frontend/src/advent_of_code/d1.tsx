import Uploader from "./uploader";
import {useState} from "react";

export default function D1() {
    const [res1, setResult1] = useState('')
    const [res2, setResult2] = useState('')

    return (
        <>
            <Uploader callback={(v) => {
                    const [left, right] = parse(v);
                    setResult1(solveP1(left, right)?.toString() ?? '')
                    setResult2(solveP2(left, right)?.toString() ?? '')
                }
            }/>
            <h1>{'p1: ' + res1}</h1>
            <h1>{'p2: ' + res2}</h1>
        </>
    )
}

function parse(v :string): [number[], number[]] {
    const lines = v.split('\n')
    const splitedLines = lines.map(v => v.split(/\s+/)).map(pair => [Number(pair[0]), Number(pair[1])])
    const left = splitedLines.map(v=> v[0])
    const right = splitedLines.map(v=> v[1])
    return [left, right];
}

function solveP1(leftGiven: number[], rightGiven: number[]): number | null {
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

function solveP2(left: number[], right: number[]): number | null {
    const occurences = new Map<number,number>();
    right.forEach(v => occurences.set(v, (occurences.get(v) ?? 0) + 1))
    return left.map(v => v * (occurences.get(v) ?? 0)).reduce((a,b) => a+b)
}