import React, { useState } from "react";

function NoteInput({ addNote }) {
  const [value, setValue] = useState( { title: "", noteDetail: ""} );

  const handleSubmit = (e) => {
    console.log(value)
    e.preventDefault();
    if (!value.title || !value.noteDetail) return;
    addNote(value);
    setValue({ title: "", noteDetail: ""});
  };

  return (
    <div className="Note-input">
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="inputTitle">Title</label>
          <input
            id="inputTitle"
            type="text"
            className="form-control"
            value={value.title}
            placeholder="Enter tile"
            onChange={e => {
              setValue({
                ...value,
                title: e.target.value
              })
            }}
          />
        </div>
        <div className="form-group">
          <label htmlFor="detailInput">Detail</label>
          <textarea 
            id="detailInput"
            className="form-control"
            rows="5"
            value={value.noteDetail}
            placeholder="Enter detail"
            onChange={e => {
              setValue({
                ...value,
                noteDetail: e.target.value
              })
            }}
          ></textarea>
        </div>
        <button type="submit" className="btn btn-primary">Add Note</button>
      </form>
    </div>
  );
}

export default NoteInput;