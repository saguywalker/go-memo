import axios from "axios";

const ENDPOINT = "http://localhost:3000/api";

const FetchData = async () => {
    const result = await axios({
        method: 'get',
        url: `${ENDPOINT}/notes`,
    });
    return result;
};

const FetchByID = async (noteId) => {
    const result = await axios({
        method: 'get',
        url: `${ENDPOINT}/notes/${noteId}`,
    });
    return result.data;
};

const Store = async (note) => {
    const result = await axios({
        method: 'post',
        url: `${ENDPOINT}/notes`,
        data: note,
    });
    return result;
}

const Update = async (note) => {
    const result = await axios({
        method: 'put',
        url: `${ENDPOINT}/notes/${note.id}`,
        data: note,
    });
    return result;
}
export {
    FetchData,
    FetchByID,
    Store,
    Update
};