import "../../styles/Forms.css"
import {useForm} from "react-hook-form";
import consts from "../../consts";
import post from "../../api/post";
import {useState} from "react";
import useFetch from "../../api/useFetch";

function UpdateWordGroup({wordGroupsUpdated}) {
    const [isLoading, setIsLoading] = useState(false)
    const {register, handleSubmit, formState: {errors}, reset} = useForm();
    const [wordGroups] = useFetch(`${consts.API_ADDRESS}/word_groups`, null, [], wordGroupsUpdated)

    const options = wordGroups.map((wordGroup) => <option value={wordGroup.ID}
                                                          key={wordGroup.ID}>{wordGroup.name}</option>)

    function onSubmit(data) {
        setIsLoading(true)
        const id = data.id
        const body = {word: data.word}
        post(`${consts.API_ADDRESS}/word_groups/${id}`, body, reset, setIsLoading)
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="content-form">
            <h2 className="content-form__header">Add word to a word group</h2>
            <label>Select word group</label>
            <select {...register("id", {required: true})} >
                {options}
            </select>
            {errors.id && <span className="content-form__error">This field is required</span>}
            <label>The word to add</label>
            <input {...register("word", {required: true})} />
            {errors.word && <span className="content-form__error">This field is required</span>}
            {isLoading ? <label style={{textAlign: "center"}}>Sending request...</label> : <input type="submit"/>}
        </form>
    );
}

export default UpdateWordGroup;