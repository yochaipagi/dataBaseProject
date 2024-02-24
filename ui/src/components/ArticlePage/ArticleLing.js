import '../../styles/Articles.css';
import consts from "../../consts";
import useFetch from "../../api/useFetch";
import {useState} from "react";
import {useForm} from "react-hook-form";
import ContextWindow from "../ContextWindow";

function ArticleLing({articleId}) {
    const [isLoading, setIsLoading] = useState(false)
    const [occurrences, setOccurrences] = useState([])
    const {register, handleSubmit, watch} = useForm();
    const [exprs] = useFetch(`${consts.API_ADDRESS}/ling_exprs`, null, [])

    const options = exprs.map((expr) => <option value={expr.expression}
                                                key={expr.ID}>{expr.expression}</option>)

    function onSubmit(data) {
        setIsLoading(true)
        fetch(`${consts.API_ADDRESS}/articles/${articleId}/ling_expr_pos?expr=${data.expr}`, {...options})
            .then(res => res.json())
            .then(
                (res) => {
                    setOccurrences(res.response)
                    setIsLoading(false);
                }, (error) => {
                    setIsLoading(false);
                }
            )
    }

    let exprCtxWindows = occurrences.map((occurrence) =>
        <div>
            <div>{`An occurrence found in line ${occurrence["line_number"]}:`}</div>
            <ContextWindow
                pos={{articleId, line: occurrence["line_number"], word: occurrence["word_number"]}}
                expr={watch("expr")}
            />
        </div>
    )

    return (
        <section>
            <form onSubmit={handleSubmit(onSubmit)} className="content-form">
                <label>Select a linguistic expression</label>
                {(exprs.length === 0) ?
                    <div>There are no linguistic expressions in the system</div>
                    :
                    <select {...register("expr", {required: true})} >
                        {options}
                    </select>
                }
                {isLoading ? <label style={{textAlign: "center"}}>Sending request...</label> :
                    <input value="Search" type="submit"/>}
            </form>
            {(exprCtxWindows.length) ?
                <h2 style={{margin: "20px 0px"}}>Occurrences</h2> :
                <div style={{margin: "20px 0px"}}>No occurrences found</div>
            }
            {exprCtxWindows}
        </section>
    )

}

export default ArticleLing;