import '../styles/ContextWindow.css';

const Highlighted = ({text = "", highlight = ""}) => {
    const regex = new RegExp(`([^a-zA-Z0-9]${highlight}[^a-zA-Z0-9])`, "gi");
    const parts = text.split(regex);

    return (
        <span>
      {parts.filter(String).map((part, i) => {
          return regex.test(part) ? (
              <mark key={i}>{part}</mark>
          ) : (
              <span key={i}>{part}</span>
          );
      })}
    </span>
    );
};


export default Highlighted;