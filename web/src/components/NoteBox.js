import React, { useState } from "react";
import NoteBoxEdit from "./NoteBoxEdit.js";

function NoteBox({ note, setEditValueCallbackParent }) {
  const [isEdit, setIsEdit] = useState(false);

  if (note == null) return;
  console.log(note);

  return (
    <>
      {
        !isEdit ? 
        <div className="Note-box">
          <h4>{note.title}</h4>
          <p className="text">{note.noteDetail}</p>
          <button type="button" className="btn btn-link" onClick={() => setIsEdit(true)}>Edit</button>
        </div> :
        <NoteBoxEdit
          note={note}
          setIsEditCallback={(isEditData) => setIsEdit(isEditData)}
          setEditValueCallback={(editValue) => setEditValueCallbackParent(editValue)}
        />
      }
    </>
    
  );
}

export default NoteBox;