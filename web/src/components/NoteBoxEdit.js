import React, { useState } from "react";

function NoteBox({ note, setIsEditCallback, setEditValueCallback }) {
  const [editValue, setEditValue] = useState({...note});

  function submitEdit() {
    // Edit request to api
    setEditValueCallback(editValue);
    setIsEditCallback(false);
    setEditValue({ title: "", noteDetail: "" });
  }

  return (
    <div className="Note-box-edit">
        <input
            type="text"
            className="form-control"
            value={editValue.title}
            placeholder="Enter tile"
            onChange={e => setEditValue({
                ...editValue,
                title: e.target.value
            })}
        />
        <textarea
            className="form-control"
            rows="5"
            value={editValue.noteDetail}
            onChange={e => setEditValue({
                ...editValue,
                noteDetail: e.target.value
            })}
        ></textarea>
      <button className="btn btn-primary" onClick={submitEdit}>Submit</button>
      <button className="btn btn-danger" onClick={() => setIsEditCallback(false)}>Cancel</button>
    </div>
  );
}

export default NoteBox;