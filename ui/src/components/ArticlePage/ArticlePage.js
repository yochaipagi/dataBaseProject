import '../../styles/Articles.css';
import Button from "../Button";
import ArticleContent from "./ArticleContent";
import {useState} from "react";
import Index from "../Index/Index";
import ArticleLing from "./ArticleLing";

function ArticlePage({articleMeta}) {
    const [articleView, setArticleView] = useState("content")

    function switchView(view) {
        setArticleView(view)
    }

    let articleBody
    switch (articleView) {
        case "index":
            articleBody = <Index articleId={articleMeta["ID"]}/>
            break
        case "lingExpr":
            articleBody = <ArticleLing articleId={articleMeta["ID"]}/>
            break
        default:
            articleBody = <ArticleContent articleId={articleMeta["ID"]}/>
    }

    return (
        <div className="article_page">
            <h2 className="article__title">{articleMeta["title"]}</h2>
            <div className="article__section">
                <span>By {articleMeta["author"]}</span>
                <span>Published at {new Date(articleMeta["published_at"]).toDateString()}</span>
                <span>Source: {articleMeta["source"]}</span>
                <span>Pages: {articleMeta["pages_count"]}</span>
            </div>
            <section className="article__buttons">
                <Button onClick={() => switchView("content")}>Content</Button>
                <Button onClick={() => switchView("index")}>Index</Button>
                <Button onClick={() => switchView("lingExpr")}>Linguistic Expression</Button>
            </section>
            {articleBody}
        </div>
    )

}

export default ArticlePage;