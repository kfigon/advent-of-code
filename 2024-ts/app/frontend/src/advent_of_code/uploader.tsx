type UploaderPropos = {
    callback: (s: string)=>void
}
export default function Uploader(u: UploaderPropos) {
    return (
        <textarea 
            onChange={(e) => u.callback(e.target.value)}
            placeholder="Paste or type your text here..."
            rows={15}
         />
    )
}