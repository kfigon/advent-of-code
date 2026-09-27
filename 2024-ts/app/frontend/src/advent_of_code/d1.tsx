import Uploader from "./uploader";

export default function D1() {
    return (
        <>
            <Uploader callback={(v) => console.log(v)}/>
            <h1>foobar</h1>
        </>
    )
}