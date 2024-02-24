import "../../styles/Forms.css"
import {useForm} from "react-hook-form";
import consts from "../../consts";
import post from "../../api/post";
import {useState} from "react";

function CreateWordGroup({setWordGroupsUpdated}) {
    const [isLoading, setIsLoading] = useState(false)
    const {register, handleSubmit, formState: {errors}, reset} = useForm();

    function onSubmit(data) {
        setIsLoading(true)
        const wordsList = data["words"].replaceAll(", ", ",").split(",")
        data["words"] = []
        for (const word of wordsList) {
            data["words"].push({word})
        }
        post(`${consts.API_ADDRESS}/word_groups`, data, reset, setIsLoading)
            .then(() => setWordGroupsUpdated((prev) => prev + 1))
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="content-form">
            <h2 className="content-form__header">Create word group</h2>
            <label>Word group name</label>
            <input {...register("name", {required: true})} />
            {errors.name && <span className="content-form__error">This field is required</span>}
            <label>Comma separated list of words</label>
            <input {...register("words", {required: true})} />
            {errors.words && <span className="content-form__error">This field is required</span>}
            {isLoading ? <label style={{textAlign: "center"}}>Sending request...</label> : <input type="submit"/>}
        </form>
    );
}

export default CreateWordGroup;