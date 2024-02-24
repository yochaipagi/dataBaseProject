import {useState} from "react";
import "../../styles/Forms.css"
import Button from "../Button";
import CreateArticle from "./CreateArticle";
import CreateLingExpr from "./CreateLingExpr";
import ManageWordGroup from "./ManageWordGroup";

function ContentManager() {
    const [resource, setResource] = useState("article")

    let manageResource
    switch (resource) {
        case "wordGroup":
            manageResource = <ManageWordGroup />
            break
        case "lingExpr":
            manageResource = <CreateLingExpr/>
            break
        default:
            manageResource = <CreateArticle/>
    }

    return (
        <section className="content-manager">
            <div className="content-manager__nav">
                <Button onClick={() => setResource("article")} size="large">Create Article</Button>
                <Button onClick={() => setResource("wordGroup")} size="large">Create/Update Work Group</Button>
                <Button onClick={() => setResource("lingExpr")} size="large">Create Linguistic Expr</Button>
            </div>
            {manageResource}
        </section>
    );
}

export default ContentManager;
