import React, { useState } from "react";
import NoteBoxDetail from "./NoteBoxDetail.js";

function NoteBox({ note, setEditValueCallbackToParent }) {
  const [isSeeMore, setIsSeeMore] = useState(false);

  if (note == null) return;

  return (
    <>
      {
        !isSeeMore ? 
        <div className="Note-box">
          <div className="box-header">
            <h4>{note.title}</h4>
            <p>Last edit: {note.lastEdit}</p>
          </div>
          <p className="text">{note.detail}</p>
          <button type="button" className="btn btn-link" onClick={() => setIsSeeMore(true)}>More</button>
        </div> :
        <NoteBoxDetail
          note={note}
          onCloseNoteBoxDetail={(isSeeMoreData) => setIsSeeMore(isSeeMoreData)}
          setEditValueCallback={(editValue) => setEditValueCallbackToParent(editValue)}
        />
      }
    </>
    
  );
}

export default NoteBox;