import "../../styles/Forms.css"
import {useForm} from "react-hook-form";
import consts from "../../consts";
import post from "../../api/post";
import {useState} from "react";

function CreateLingExpr() {
    const [isLoading, setIsLoading] = useState(false)
    const {register, handleSubmit, formState: {errors}, reset} = useForm();

    function onSubmit(data) {
        setIsLoading(true)
        post(`${consts.API_ADDRESS}/ling_exprs`, data, reset, setIsLoading)
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="content-form">
            <h2 className="content-form__header">Create linguistic expression</h2>
            <label>Expression</label>
            <input {...register("expression", {required: true})} />
            {errors.expression && <span className="content-form__error">This field is required</span>}
            {isLoading ? <label style={{textAlign:"center"}}>Sending request...</label> : <input type="submit"/>}
        </form>
    );
}

export default CreateLingExpr;