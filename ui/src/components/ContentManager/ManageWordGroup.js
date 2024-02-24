import {useState} from "react";
import CreateWordGroup from "./CreateWordGroup";
import UpdateWordGroup from "./UpdateWordGroup";

export default function ManageWordGroup() {
    const [wordGroupsUpdated, setWordGroupsUpdated] = useState(0)

    return <>
        <CreateWordGroup setWordGroupsUpdated={setWordGroupsUpdated}/>
        <UpdateWordGroup wordGroupsUpdated={wordGroupsUpdated}/>
    </>
}