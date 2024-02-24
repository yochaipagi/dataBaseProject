import '../styles/Articles.css';

function ArticleCard({article, onClick}) {

    return (
        <div className="article_card" onClick={onClick}>
            <h4 className="article__title">{article["title"]}</h4>
            <div className="article__section">
                <span>By {article["author"]}</span>
                <span>Published at {new Date(article["published_at"]).toDateString()}</span>
                <span>Source: {article["source"]}</span>
                <span>Pages: {article["pages_count"]}</span>
            </div>
        </div>
    );
}

export default ArticleCard;