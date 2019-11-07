import React, { useState, useEffect } from "react";
import NoteBox from "./components/NoteBox";
import NoteInput from "./components/NoteInput";
import {FetchData, Store, Update} from "./services";

function App() {
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    FetchData().then((res) => {
      setNotes(res.data)
    });
  }, []);

  const addNote = (note) => {
    Store(note).then((res) => {
      setNotes([...notes, res.data])
    });
  };

  const setEditNoteMain = (editValue) => {
    Update(editValue).then((res) => {
      const editValueIndex = notes.indexOf(notes.find(note => note.id === editValue.id));
      notes[editValueIndex] = res.data;
      setNotes(notes);
    });
  };

  return (
    <div className="App">
      <NoteInput 
        addNote={addNote}
      />
      <h2 style={{color: 'white'}}>Go-Memo</h2>
      <div className="Note-list">
        {
          notes.length !== 0 ? 
            notes.map((note) => (
              <NoteBox
                key={note.id}
                note={note}
                setEditValueCallbackToParent={(edit) => setEditNoteMain(edit)}
              />
            )) :
            <h4>There is no note.</h4>  
        }
      </div>
    </div>
  );
}

export default App;