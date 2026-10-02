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

export function solve(v: string, min: number, max: number): Result {
    const lines: number[][] = v.split('\n').map(line => line.split(/\s+/).map(Number));

    let p1 = 0;
    let p2 = 0;
    for(let line of lines) {
        if (line.length < 2){
            continue
        }

        const upward = line[0] < line[1];
        if (safeLine(line, upward, min, max, 0)){
            p1++;
        }
        
        if (safeLine(line, upward, min, max, 1)){
            p2++;
        }
        
    }
    return {p1, p2}
}

const safeLine = (line: number[], upward: boolean, min: number, max: number, numberOfSkips: number): boolean => {
    const compareWith = (thisId: number, nextId: number): boolean | null => {
        if(nextId >= line.length) return null;

       const a = line[thisId];
       const b = line[nextId];

       const diff = upward 
        ? b-a 
        : a-b;

        return diff >= min && diff <= max
    }

    for(let i =0; i < line.length; i++){
        let res = compareWith(i, i+1)
        if (res === null) {
            break
        } else if (res === true) {
            continue;
        } else if (numberOfSkips <= 0) {
            return false;
        }
        console.log(`lien ${line} mismatch on ${line[i]} ${line[i+1]}, idx ${i}`)

        res = compareWith(i-1, i+2)
        if (res === true) {
            numberOfSkips--;
            console.log(`line ${line} skipping from ${line[i+1]} idx  ${i+1}`)
            i = i+1;            
        } else if (res === null) {
            break
        } else if (res === false) {
            return false
        }

    }
    return true
}
