import React, { useState } from "react";
import NoteBox from "./components/NoteBox";
import NoteInput from "./components/NoteInput";

function App() {
  const [notes, setNotes] = useState([]);

  function addNote({ title, noteDetail }) {
    const newNotes = [...notes, { title, noteDetail }];
    
    setNotes(newNotes);
  };

  // function updateNote(index, {title, noteDetail}) {
  //   const newNotes = [...notes];
  //   newNotes.splice(index, 1, { title, noteDetail })
  //   setNotes(newNotes);
  // };

  function setEditNote(editValue, index) {
    const newNotes = [...notes];
    newNotes[index] = editValue;
    setNotes(newNotes);
  }

  return (
    <div className="App">
      <NoteInput 
        addNote={addNote}
      />
      <h2 style={{color: 'white'}}>My Notes</h2>
      <div className="Note-list">
        {
          notes.length !== 0 ? 
            notes.map((note, index) => (
              <NoteBox
                key={index}
                index={index}
                note={note}
                setEditValueCallbackParent={(edit) => setEditNote(edit, index)}
              />
            )) :
            <h4>There is no note.</h4>  
        }
      </div>
    </div>
  );
}

export default App;