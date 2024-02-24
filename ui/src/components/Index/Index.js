import '../../styles/Index.css';
import consts from "../../consts";
import useFetch from "../../api/useFetch";
import IndexWord from "./IndexWord";
import {useState} from "react";
import ContextWindow from "../ContextWindow";
import {useForm} from "react-hook-form";
import Button from "../Button";

function Index({articleId = ""}) {
    const wordsPerPage = 500
    const [selectedWord, setSelectedWord] = useState(null)
    const {register, handleSubmit} = useForm();
    const [wordGroupID, setWordGroupID] = useState("")
    const [pageNum, setPageNum] = useState(1)
    const [index, isLoading, error] = useFetch(`${consts.API_ADDRESS}/article_words/index?article_id=${articleId}&word_group_id=${wordGroupID}`, null, [], pageNum)
    const words = index.slice((pageNum-1)*wordsPerPage,pageNum*wordsPerPage).map((wordObj) => <IndexWord key={wordObj["word"]} wordObj={wordObj} selectWord={setSelectedWord}/>)
    const [wordGroups] = useFetch(`${consts.API_ADDRESS}/word_groups`, null, [])
    const options = wordGroups.map((expr) => <option value={expr.ID}
                                                key={expr.ID}>{expr.name}</option>)

    function onSubmit(data) {
        setWordGroupID(data["wordGroupID"])
        setPageNum(1)
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return (
            <section>
                {(selectedWord !== null) ?
                    <ContextWindow pos={selectedWord["pos"]} expr={selectedWord["word"]} /> :
                    ""
                }
                <form onSubmit={handleSubmit(onSubmit)} className="content-form">
                    <label>Select a word group to filter by</label>
                    {(wordGroups.length === 0) ?
                        <div>There are no word groups in the system</div>
                        :
                        <select {...register("wordGroupID", {required: true})} >
                            {options}
                        </select>
                    }
                    {isLoading ? <label style={{textAlign: "center"}}>Sending request...</label> :
                        <input type="submit"/>}
                </form>
                <div className="index">
                    {words}
                </div>
                <div className="index__pagination">
                    {(pageNum > 1) ? <Button onClick={() => setPageNum(prev => prev - 1)}>Previous page</Button> : null}
                    {(pageNum < Math.ceil(index.length / wordsPerPage)) ? <Button onClick={() => setPageNum(prev => prev + 1)}>Next page</Button> : null}
                </div>
            </section>
        )
    }
}

export default Index;