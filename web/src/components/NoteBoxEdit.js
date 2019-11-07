import React, { useState } from "react";

function NoteBox({ note, setIsEditCallback, setEditValueCallback, closeNoteBoxDetail }) {
  const [editValue, setEditValue] = useState({...note});

  function submitEdit() {
    setEditValueCallback(editValue);
    setIsEditCallback(false);
    closeNoteBoxDetail();
    setEditValue({ title: "", detail: "" });
  }

  return (
    <div>
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
        value={editValue.detail}
        onChange={e => setEditValue({
            ...editValue,
            detail: e.target.value
        })}
      ></textarea>
      <br />
      <br />
      <button className="btn btn-primary" onClick={submitEdit}>Submit</button>
      <button className="btn btn-danger" onClick={() => setIsEditCallback(false)}>Cancel</button>
    </div>
  );
}

export default NoteBox;