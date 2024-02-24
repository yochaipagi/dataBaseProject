import '../styles/Button.css';

function Button({onClick, size, children}) {
    let className = "btn"
    if (size === "small") {
        className += " btn--small"
    } else if (size === "large") {
        className += " btn--large"
    }

    return (
        <button className={className} onClick={onClick}>{children}</button>
    );
}

export default Button;