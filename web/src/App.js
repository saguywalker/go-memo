import React, { useState } from "react";
import "./App.css";

function Note({ note, index }) {
  return (
    <div className="Note">
      <h4>{note.title}</h4><br />
      {note.noteDetail}
    </div>
  );
}

function NoteForm({ addNote }) {
  const [value, setValue] = useState( { title: "", noteDetail: ""} );

  const handleSubmit = e => {
    console.log(value)
    e.preventDefault();
    if (!value.title || !value.noteDetail) return;
    addNote(value);
    setValue("", "");
  };

  return (
    <form>
      <input
        type="text"
        className="inputTitle"
        value={value.title}
        onChange={e => {
          console.log("onChange(title): ", e.target.value);
          e.preventDefault();
          setValue( {title: e.target.value.title} )
        }}
      />
      <input 
        type="text"
        className="inputDetail"
        value={value.noteDetail}
        onChange={e => {
          console.log("onChange(detail): ", e.target.value);
          e.preventDefault();
          setValue( {noteDetail: e.target.value.noteDetail} )
        }}
      />
      <button onClick={handleSubmit}>Add Note</button>
    </form>
  );
}

function App() {
  const [notes, setNotes] = useState([
    {
      title: "Learn about React",
      noteDetail: "this is a detail of the first note"
    },
    {
      title: "Meet friend for lunch",
      noteDetail: "At Tora Yakiniku on 29 Dec, 14:30"
    },
    {
      title: "Build really cool todo app",
      noteDetail: "It's a good time to learn React"
    }
  ]);

  const addNote = (title, noteDetail) => {
    const newNotes = [...notes, { title, noteDetail }];
    setNotes(newNotes);
  };

  const updateNote = (index, title, noteDetail) => {
    const newNotes = [...notes];
    newNotes.splice(index, 1, { title, noteDetail })
    setNotes(newNotes);
  };

  return (
    <div className="App">
      <div className="Note-list">
        {notes.map((note, index) => (
          <Note
            key={index}
            index={index}
            note={note}
          />
        ))}
        <NoteForm 
          addNote={addNote}
          updateNote={updateNote}
        />
      </div>
    </div>
  );
}

export default App;