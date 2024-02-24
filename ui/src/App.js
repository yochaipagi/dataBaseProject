import './App.css';
import Navbar from "./components/Navbar";
import Articles from "./components/Articles";
import {useState} from "react";
import ContentManager from "./components/ContentManager/ContentManager";
import Index from "./components/Index/Index";

function App() {
    const [page, setPage] = useState("articles")

    let pageToRender
    switch (page) {
        case "index":
            pageToRender = <Index />
            break
        case "create":
            pageToRender = <ContentManager />
            break
        default:
            pageToRender = <Articles />
    }

    return (
        <div className="app">
            <Navbar setPage={setPage} page={page}/>
            <div className="page-container">
                {pageToRender}
            </div>
        </div>
    );
}

export default App;
