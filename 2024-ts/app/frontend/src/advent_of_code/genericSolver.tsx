import { useState } from "react";
import Uploader from "./uploader";

export default function GenericSolver<T, K, V>(parser: (v: string) => T | null, p1: (v: T)=>K, p2: (v:T) => V) {
    const [res1, setResult1] = useState('')
    const [res2, setResult2] = useState('')

    return (
        <>
            <Uploader callback={(v) => {
                    const res = parser(v);
                    if (!res)return;
                    setResult1(p1(res)?.toString() ?? '')
                    setResult2(p2(res)?.toString() ?? '')
                }
            }/>
            <h1>{'p1: ' + res1}</h1>
            <h1>{'p2: ' + res2}</h1>
        </>
    )
}