import '../../styles/Index.css';
import Button from "../Button";
import {useState} from "react";

function IndexWord({wordObj, selectWord}) {
    const [expand, setExpand] = useState(false)
    let wordsIndex = []

    if (expand) {
        const occurrences = wordObj["index"].split('\n')
        const indexTable = occurrences.map((occurrence) => {
            const pos = occurrence.split(',')
            const [articleId, pageNum, lineNum, wordNum] = pos
            return <div onClick={() => onWordClick(articleId, pageNum, lineNum, wordNum, wordObj["word"])}
                        key={occurrence} className="index__word__row index__word__row--clickable">
                <span>{articleId}</span>
                <span>{pageNum}</span>
                <span>{lineNum}</span>
                <span>{wordNum}</span>
            </ div>
        })
        wordsIndex =
            <div className="index__word__table">
                <div className="index__word__row">
                    <span>Article ID</span>
                    <span>Page</span>
                    <span>Line</span>
                    <span>Word</span>
                </div>
                {indexTable}
            </div>
    }

    function onWordClick(articleId, pageNum, lineNum, wordNum, wordTxt) {
        window.scrollTo({top: 0, left: 0, behavior: 'smooth'});
        selectWord({pos: {articleId, page: pageNum, line: lineNum, word: wordNum}, word: wordTxt})
        setExpand(false)
    }

    function expandOrCollapse() {
        setExpand((prev) => !prev)
    }


    return (
        <div className="index__word__container">
            <div className="index__word">
                <b>Word</b>
                <b>Occurrences</b>
                <div style={{gridColumnStart: "3", gridRow: "1 / 3"}}>
                    <Button onClick={expandOrCollapse} size="small">{expand ? `Hide Index` : `Show Index`}</Button>
                </div>
                <div>{wordObj["word"]} </div>
                <div>{wordObj["count"]}</div>
            </div>
                {expand ? wordsIndex : null }
        </div>
    )
}

export default IndexWord;