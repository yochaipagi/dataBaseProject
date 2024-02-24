import consts from "../consts";
import '../styles/ContextWindow.css';
import useFetch from "../api/useFetch";
import Button from "./Button";
import {useEffect, useState} from "react";
import Highlighted from "./Highlighted";

function ContextWindow({pos, expr}) {
    const [currPos, setCurrPos] = useState(pos)
    const [res, isLoading, error] = useFetch(getUrlToFetch(currPos), null, {}, currPos)

    let lines
    if (res && res["lines"]) {
        lines = res["lines"].map((line) => <p key={line["line_number"]}>
            <b style={{marginRight: "5px"}}>
                {line["line_number"]}
            </b>
            <Highlighted text={line["content"]} highlight={expr}/>
        </p>)
    }

    useEffect(() => {
        setCurrPos(pos)
    }, [pos]);

    function handleArrowClick(delta) {
        setCurrPos((prevPos) => {
            const newLineNum = parseInt(prevPos["line"]) + delta
            if ((newLineNum === 0 && delta === -1) || (lines.length === 1 && delta === 1) ) {
                return prevPos
            }
            return {
                ...prevPos,
                page: Math.floor(newLineNum / 10) + 1,
                line: newLineNum,
                word: 1
            }
        })
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return <div className="context-window">
            <div className="context-window__arrows">
                <Button onClick={() => handleArrowClick(-1)} size="small">&#8593;</Button>
                <Button onClick={() => handleArrowClick(1)} size="small">&#8595;</Button>
            </div>
            <div className="context-window__lines">
                {lines}
            </div>
        </div>
    }
}

function getUrlToFetch(pos) {
    return `${consts.API_ADDRESS}/article_words?article_id=${pos["articleId"]}&line_num=${pos["line"]}&word_num=${pos["word"]}`
}

export default ContextWindow;