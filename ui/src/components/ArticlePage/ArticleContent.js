import '../../styles/Articles.css';
import consts from "../../consts";
import useFetch from "../../api/useFetch";

function ArticleContent({articleId}) {
    const [article, isLoading, error] = useFetch(`${consts.API_ADDRESS}/articles/${articleId}`)

    let lineElems
    if (article["content"]) {
        const articleLines = article["content"].split("\n")
        lineElems = articleLines.map((line, i) => <p key={i}><b style={{marginRight: "2px"}}>{i + 1} </b>{line}</p>)
        for (let i = 0; i < lineElems.length; i += 10) {
            lineElems.splice(i, 0, <h4 key={lineElems.length+1} style={{margin: "10px 0px"}}>{`Page ${i / 10 + 1}`}</h4>);
        }
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return <section className="article__content">
            {lineElems}
        </section>
    }
}

export default ArticleContent;