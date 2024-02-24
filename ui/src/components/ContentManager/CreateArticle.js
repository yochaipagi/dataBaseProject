import "../../styles/Forms.css"
import {useForm} from "react-hook-form";
import consts from "../../consts";
import post from "../../api/post";
import {useState} from "react";

function CreateArticle() {
    const [isLoading, setIsLoading] = useState(false)
    const {register, handleSubmit, formState: {errors}, reset} = useForm();

    function onSubmit(data) {
        setIsLoading(true)
        post(`${consts.API_ADDRESS}/articles`, data, reset, setIsLoading)
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="content-form">
            <h2 className="content-form__header">Create article</h2>
            <label>Title</label>
            <input {...register("title", {required: true})} />
            {errors.title && <span className="content-form__error">This field is required</span>}
            <label>Author</label>
            <input {...register("author", {required: true})} />
            {errors.author && <span className="content-form__error">This field is required</span>}
            <label>Published at <span className="content-form__note">(Has to be according to RFC1123 - for example: Mon, 02 Jan 2006 15:04:05 MST)</span></label>
            <input {...register("published_at", {required: true})} />
            {errors.published_at && <span className="content-form__error">This field is required</span>}
            <label>Source</label>
            <input {...register("source", {required: true})} />
            {errors.source && <span className="content-form__error">This field is required</span>}
            <label>Content</label>
            <textarea rows="10" {...register("content", {required: true})} />
            {errors.content && <span className="content-form__error">This field is required</span>}
            {isLoading ? <label style={{textAlign:"center"}}>Sending request...</label> : <input type="submit"/>}
        </form>
    );
}

export default CreateArticle;