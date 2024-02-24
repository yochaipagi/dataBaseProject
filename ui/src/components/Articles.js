import '../styles/Articles.css';
import ArticleCard from "./ArticleCard";
import {useState} from "react";
import consts from '../consts.js'
import ArticlePage from "./ArticlePage/ArticlePage";
import useFetch from "../api/useFetch";

function Articles() {
    const [selectedArticleID, setSelectedArticleID] = useState(0)
    const [articles, isLoading, error] = useFetch(`${consts.API_ADDRESS}/articles`, null, [])

    function handleCardClick(articleID) {
        setSelectedArticleID(articleID)
    }

    let articleCards = []
    if (articles) {
        articleCards = articles.map(article => <ArticleCard key={article.ID} article={article}
                                                            onClick={() => handleCardClick(article.ID)}/>)
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else if (selectedArticleID !== 0) {
        return <ArticlePage articleMeta={articles.filter((article) => (article.ID === selectedArticleID))[0]}/>
    } else {
        return (articles !== null && articles.length) ?
            <div>
                <h1 style={{textAlign:"center", fontSize:"xxx-large"}}>Articles</h1>
                {articleCards}
            </div> :
            <div>No articles found</div>
    }
}

export default Articles;