import '../styles/Navbar.css';
import logo from '../assets/Concordance.png';

function Navbar({page, setPage}) {

    return (
        <div className="navbar">
            <img src={logo} alt="Logo" />
            <div onClick={() => setPage("articles")} className={(page === "articles") ? "navbar__item navbar__item--black" : "navbar__item"}>Articles</div>
            <div onClick={() => setPage("index")} className={(page === "index") ? "navbar__item navbar__item--black" : "navbar__item"}>Index</div>
            <div onClick={() => setPage("create")} className={(page === "create") ? "navbar__item navbar__item--black" : "navbar__item"}>Create</div>
        </div>
    );
}

export default Navbar;