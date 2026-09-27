import Uploader from "./uploader";
import {useState} from "react";

export default function D1() {
    const [res, setResult] = useState('')
    return (
        <>
            <Uploader callback={(v) => solveP1(v, setResult)}/>
            <h1>{res}</h1>
        </>
    )
}

function solveP1(v: string, setResult: (s:string) => void) {
    setResult(v)
}