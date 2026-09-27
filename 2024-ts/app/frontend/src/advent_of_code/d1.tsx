import Uploader from "./uploader";
import {useState} from "react";

export default function D1() {
    const [res1, setResult1] = useState('')
    const [res2, setResult2] = useState('')

    return (
        <>
            <Uploader callback={(v) => {
                    setResult1(solveP1(v)?.toString() ?? '')
                    setResult2(solveP2(v)?.toString() ?? '')
                }
            }/>
            <h1>{'p1: ' + res1}</h1>
            <h1>{'p2: ' + res2}</h1>
        </>
    )
}

function solveP1(v: string): number | null {
    const lines = v.split('\n')
    const splitedLines = lines.map(v => v.split(/\s+/)).map(pair => [Number(pair[0]), Number(pair[1])])
    const left = splitedLines.map(v=> v[0])
    const right = splitedLines.map(v=> v[1])

    left.sort()
    right.sort()
    if (left.length !== right.length) return null;

    let res = 0;
    for(let i = 0; i < left.length; i++) {
        res += Math.abs(left[i] - right[i]);
    }
    
    return res;
}

function solveP2(_: string): number | null {
    return null;
}