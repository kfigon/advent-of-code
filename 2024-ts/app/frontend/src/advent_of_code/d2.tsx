import { useState } from "react";
import Uploader from "./uploader";

export default function D2(){
    const [p1, setP1] = useState('')
    const [p2, setP2] = useState('')

    const [min, setMin] = useState(1)
    const [max, setMax] = useState(3)
    
    const solveAndSet = (v: string) => {
        const got = solve(v, min, max)
        got.p1 && setP1(got.p1.toString())
        got.p2 && setP2(got.p2.toString())
    }

    return (
        <>
            <input type="number" value={min} onChange={v => setMin(Number(v.target.value ?? min))}></input>
            <input type="number" value={max} onChange={v => setMax(Number(v.target.value ?? max))}></input>
            <Uploader callback={solveAndSet}/>
            <h1>{p1 && `p1: ${p1}`}</h1>
            <h1>{p2 && `p2: ${p2}`}</h1>
        </>
    )
}

type Result = {
    p1?: number
    p2?: number
}

function solve(v: string, min: number, max: number): Result {
    const lines: number[][] = v.split('\n').map(line => line.split(/\s+/).map(Number));

    let p1 = 0;
    let p2 = 0;
    for(let line of lines) {
        if (line.length < 2){
            continue
        }

        const upward = line[0] < line[1];
        if (safeLine(line, upward, min, max)){
            p1++;
        }
    }
    return {p1, p2}
}

const safeLine = (line: number[], upward: boolean, min: number, max: number): boolean => {
    for(let i =0; i < line.length-1; i++){
       const a = line[i];
       const b = line[i+1];

       const diff = upward 
        ? b-a 
        : a-b;
        
        if (diff < min || diff > max){
            return false
        }
    }
    return true
}
